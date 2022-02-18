package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Node struct {
	ID       int        `orm:"column(id);" json:"id"`
	UUID     string     `orm:"column(uuid);varcher(64)" json:"uuid"`
	Hostname string     `orm:"varchar(64)" json:"hostname"`
	Addr     string     `orm:"varchar(512)" json:"addr"`
	CreateAt *time.Time `orm:"auto_now_add" json:"create_at"`
	UpdateAt *time.Time `orm:"auto_now" json:"update_at"`
	DeleteAt *time.Time `orm:"null" json:"delete_at"`
	// 一对多
	Jobs []*Job `orm:"reverse(many)" json:"jobs"`
}

type NodeManager struct{}

func (m *NodeManager) Create(uuid, hostname, addr string) (*Node, error) {
	ormer := orm.NewOrm()
	result := &Node{
		Hostname: hostname,
		UUID:     uuid,
		Addr:     addr,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *NodeManager) List(q string, start int64, length int) ([]*Node, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Node{})

	condition := orm.NewCondition()
	condition = condition.And("delete_at__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("uuid__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Node

	queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

func (m *NodeManager) GetById(id int) *Node {
	ormer := orm.NewOrm()
	result := &Node{}
	err := ormer.QueryTable(result).Filter("Id__exact", id).Filter("Deleteat__isnull", true).One(result)
	if err == nil {
		return result
	}
	return nil
}

func (m *NodeManager) GetAllUUID() []*Node {
	ormer := orm.NewOrm()
	nodes := []*Node{}
	_, err := ormer.QueryTable(&Node{}).Filter("Deleteat__isnull", true).All(&nodes)
	fmt.Println("orm err:", err)
	if err == nil {
		return nodes
	}
	return nil
}

func (m *NodeManager) Modify(id int, hostname, addr string) (*Node, error) {
	ormer := orm.NewOrm()
	if node := m.GetById(id); node != nil {
		node.Hostname = hostname
		node.Addr = addr
		if _, err := ormer.Update(node); err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}

func (m *NodeManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&Node{}).Filter("id__exact", id).Update(orm.Params{"deleteat": time.Now()})
	return nil
}

func NewNodeManager() *NodeManager {
	return &NodeManager{}
}

var DefaultNodeManager = NewNodeManager()

// job
type Job struct {
	ID       int        `orm:"column(id)" json:"id"`
	Key      string     `orm:"varchar(64)" json:"key"`
	Remark   string     `orm:"varchar(512)" json:"remark"`
	CreateAt *time.Time `orm:"auto_now_add" json:"create_at"`
	UpdateAt *time.Time `orm:"auto_now" json:"update_at"`
	DeleteAt *time.Time `orm:"null" json:"delete_at"`

	Node *Node `orm:"rel(fk)" json:"node"`
	// 一对多
	Targets []*Target `orm:"reverse(many)" json:"targets"`
}

type JobManager struct{}

func NewJobManager() *JobManager {
	return &JobManager{}
}

var DefaultJobManager = NewJobManager()

func (m *JobManager) List(q string, start int64, length int) ([]*Job, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Job{})

	condition := orm.NewCondition()
	condition = condition.And("delete_at__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("key__icontains", q)
		query = query.Or("remark__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var jobs []*Job

	queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&jobs)
	//for _, job := range jobs{  //反向查targets
	//	ormer.LoadRelated(job, "Targets")
	//}
	return jobs, total, qtotal
}

func (m *JobManager) GetByUUId(uuid string) *Node {
	ormer := orm.NewOrm()
	result := &Node{}
	err := ormer.QueryTable(result).Filter("UUID__exact", uuid).Filter("Deleteat__isnull", true).One(result)
	if err == nil {
		return result
	}
	return nil
}




func (m *JobManager) Create(uuid string, key, remark string) (*Job, error) {
	node := m.GetByUUId(uuid)
	fmt.Println("node", node)
	ormer := orm.NewOrm()
	result := &Job{
		Key:    key,
		Remark: remark,
		Node:   node,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *JobManager) GetById(id int) *Job {
	ormer := orm.NewOrm()
	result := &Job{}
	err := ormer.QueryTable(result).Filter("Id__exact", id).Filter("Deleteat__isnull", true).One(result)
	if err == nil {
		return result
	}
	return nil
}

func (m *JobManager) Modify(id int, key, remark string) (*Job, error) {
	ormer := orm.NewOrm()
	if job := m.GetById(id); job != nil {
		job.Key = key
		job.Remark = remark
		if _, err := ormer.Update(job); err != nil {
			return nil, err
		}
		return job, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}

func (m *JobManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&Job{}).Filter("id__exact", id).Update(orm.Params{"deleteat": time.Now()})
	return nil
}

func (m *JobManager) GetAllJobByUUID(uuid string) []*Job{
	ormer := orm.NewOrm()
	jobs := []*Job{}
	queryset := ormer.QueryTable(&Job{})
	queryset.Filter("Deleteat__isnull", true).Filter("node__uuid",uuid).All(&jobs)
	for _, job := range jobs{
		ormer.LoadRelated(job, "Targets") //反向关联查询 RelatedSel只能正向关联fk
		ormer.LoadRelated(job, "Node") //反向关联查询
	}
	return jobs
}

// Target
type Target struct {
	ID       int        `orm:"column(id)" json:"id"`
	Name     string     `orm:"varchar(64)" json:"name"`
	Addr     string     `orm:"varchar(128)" json:"addr"`
	CreateAt *time.Time `orm:"auto_now_add" json:"create_at" json:"create_at"`
	UpdateAt *time.Time `orm:"auto_now" json:"update_at" json:"update_at"`
	DeleteAt *time.Time `orm:"null" json:"delete_at" json:"delete_at"`

	Job *Job `orm:"rel(fk)" json:"job"`
}
type TargetManager struct{}

func NewTargetManager() *TargetManager {
	return &TargetManager{}
}

func (m *TargetManager) List(q string, start int64, length int) ([]*Target, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Target{})

	condition := orm.NewCondition()
	condition = condition.And("delete_at__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("Name__icontains", q)
		query = query.Or("Addr__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Target
	fmt.Println("start::", start)
	fmt.Println("length::", length)
	queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}


func (m *TargetManager) GetByKey(key string) *Job {
	ormer := orm.NewOrm()
	result := &Job{}
	err := ormer.QueryTable(result).Filter("Key__exact", key).Filter("Deleteat__isnull", true).One(result)
	if err == nil {
		return result
	}
	return nil
}

func (m *TargetManager) Create(jobKey string, name, addr string) (*Target, error) {
	job := m.GetByKey(jobKey)
	fmt.Println("jobs::", job)
	ormer := orm.NewOrm()
	result := &Target{
		Name: name,
		Addr: addr,
		Job: job,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *TargetManager) GetById(id int) *Target {
	ormer := orm.NewOrm()
	result := &Target{}
	err := ormer.QueryTable(result).Filter("Id__exact", id).Filter("Deleteat__isnull", true).One(result)
	if err == nil {
		return result
	}
	return nil
}

func (m *NodeManager) GetAllJob() []*Job {
	ormer := orm.NewOrm()
	jobs := []*Job{}
	_, err := ormer.QueryTable(&Job{}).Filter("Deleteat__isnull", true).All(&jobs)
	if err == nil {
		return jobs
	}
	return nil
}

func (m *TargetManager) Modify(id int, name, addr string) (*Target, error) {
	ormer := orm.NewOrm()
	if target := m.GetById(id); target != nil {
		target.Name = name
		target.Addr = addr
		if _, err := ormer.Update(target); err != nil {
			return nil, err
		}
		return target, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}

func (m *TargetManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&Target{}).Filter("id__exact", id).Update(orm.Params{"deleteat": time.Now()})
	return nil
}

var DefaultTargetManager = NewTargetManager()

func init() {
	orm.RegisterModel(new(Node), new(Job), new(Target))
}
