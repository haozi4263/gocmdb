package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Ticket struct {
	Id          int        `orm:"column(id);" json:"id"`
	Name        string     `orm:"column(name);size(32);" json:"name"`
	Type        string     `orm:"column(type);size(1024);" json:"type"`
	Env         string     `orm:"column(env);size(1024);" json:"env"`
	CreateUser  string     `orm:"column(create_user);size(1024);" json:"create_user"`
	DisposeUser string     `orm:"column(dispose_user);size(1024);" json:"dispose_user"`
	Detail      string     `orm:"column(detail);size(4096);" json:"detail"`
	Remark      string     `orm:"column(remark);size(4096);" json:"remark"`
	Tel         string     `orm:"column(tel);size(1024);" json:"tel"`
	Status      string      `orm:"column(status);" json:"status"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	DoneTime    *time.Time `orm:"column(done_time);" json:"done_time"`
}

type TicketManager struct{}

func NewTicketManager() *TicketManager {
	return &TicketManager{}
}

func (m *TicketManager) GetById(id int) *Ticket {
	ticket := &Ticket{}
	ormer := orm.NewOrm()
	err := ormer.QueryTable(ticket).Filter("Id__exact", id).One(ticket)
	if err == nil {
		return ticket
	}
	return nil
}

func (m *TicketManager) GetByName(name string) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Name__exact", name).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		return user
	}

	return nil
}

func (m *TicketManager) Query(q, menu string, start int64, length int, userName string) ([]*Ticket, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Ticket{})
	condition := orm.NewCondition()
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	var ticket []*Ticket
	if menu == "my_ticket_management"{
		if q != "" {
			query := orm.NewCondition()
			query = query.Or("name__icontains", q)
			query = query.Or("id__icontains", q)
			query = query.Or("remark__icontains", q)
			query = query.Or("env__icontains", q)
			query = query.Or("create_user__icontains", q)
			query = query.Or("dispose_user__icontains", q)
			query = query.Or("status__icontains", q)
			condition = condition.AndNot("status__exact", "完成") // 未完成的工单
			condition = condition.AndCond(query)
			qtotal, _ = queryset.SetCond(condition).Count()
		}else {
			q = userName
			query := orm.NewCondition()
			query = query.Or("create_user__icontains", q)
			query = query.Or("dispose_user__icontains", q)

			condition = condition.AndNot("status__exact", "完成") // 未完成的工单
			condition = condition.AndCond(query)
			qtotal, _ = queryset.SetCond(condition).Count()
		}
	}
	if menu == "all_ticket_management" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("id__icontains", q)
		query = query.Or("remark__icontains", q)
		query = query.Or("env__icontains", q)
		query = query.Or("create_user__icontains", q)
		query = query.Or("dispose_user__icontains", q)
		query = query.Or("status__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}


	queryset.SetCond(condition).Limit(length).Offset(start).All(&ticket)
	return ticket, total, qtotal
}

func (m *TicketManager) Create(ticket *Ticket) (*Ticket, error) {
	ormer := orm.NewOrm()
	tickets := &Ticket{
		Name:     ticket.Name,
		Type: ticket.Type,
		Env: ticket.Env,
		CreateUser: ticket.CreateUser,
		DisposeUser: ticket.DisposeUser,
		Detail: ticket.Detail,
		Remark: ticket.Remark,
		Tel: ticket.Tel,
		Status: ticket.Status,
		CreatedTime: ticket.CreatedTime,
		DoneTime: ticket.DoneTime,
	}

	if _, err := ormer.Insert(tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (m *TicketManager) Modify(id int, status, name, env,disposeUser,detail,remark string, doneTime *time.Time) (*Ticket, error) {
	ormer := orm.NewOrm()
	if ticket := m.GetById(id); ticket != nil {
		ticket.Status = status
		ticket.Name = name
		ticket.Env = env
		ticket.DisposeUser = disposeUser
		ticket.Detail = detail
		ticket.Remark = remark
		ticket.DoneTime = doneTime
		if _, err := ormer.Update(ticket); err != nil {
			return nil, err
		}
		return ticket, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}


func (m *TicketManager) Done(id int) error {
	ormer := orm.NewOrm()
	if ticket := m.GetById(id); ticket != nil {
		ticket.Status = "完成"
		str := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05 MST"))
		if done, err := time.Parse("2006-01-02 15:04:05 MST", str); err == nil {
			ticket.DoneTime = &done
		}

		if _, err := ormer.Update(ticket); err != nil {
			return nil
		}
		return nil
	}
	return nil


}


func (m *TicketManager) Start(id int) error {
	ormer := orm.NewOrm()
	if ticket := m.GetById(id); ticket != nil {
		ticket.Status = "处理中"
		if _, err := ormer.Update(ticket); err != nil {
			return nil
		}
		return nil
	}
	return nil


}


func (m *UserManager) TicketManager(id int, name string, gender int, birthday *time.Time, tel, email, addr, remark string) (*User, error) {
	ormer := orm.NewOrm()
	if user := m.GetById(id); user != nil {
		user.Name = name
		user.Gender = gender
		user.Birthday = birthday
		user.Tel = tel
		user.Email = email
		user.Addr = addr
		user.Remark = remark
		if _, err := ormer.Update(user); err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, fmt.Errorf("操作对象不存在")
}

var DefaultTicketManager = NewTicketManager()

func init() {
	orm.RegisterModel(new(Ticket)) //注册数据库，用户models执行使用
}
