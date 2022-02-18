package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// 发布项目
type DeployProject struct {
	Id             int    `orm:"column(id);" json:"id"`
	Name           string `orm:"column(name);size(32);" json:"name"`
	JenkinsJob	  string `orm:"column(jenkins_job);size(32);" json:"jenkins_job"`
	Env	  string `orm:"column(env);size(128);" json:"env"`
	Business       string `orm:"column(business);size(64);" json:"business"`
	Type           string `orm:"column(type);size(1024);" json:"type"`
	Describe       string `orm:"column(describe);size(1024);" json:"describe"`
	GitUrl         string `orm:"column(git_url);size(1024);" json:"git_url"`
	BuildCmd       string `orm:"column(build_cmd);size(4096);" json:"build_cmd"`
	Port           string `orm:"column(port);size(100);" json:"port"`
	SvcType        string `orm:"column(svc_type);size(100);" json:"svc_type"`
	LimitMem       string `orm:"column(limit_mem);size(100);" json:"limit_mem"`
	LimitCpu       string `orm:"column(limit_cpu);size(100);" json:"limit_cpu"`
	ReqCpu         string `orm:"column(req_cpu);size(100);" json:"req_cpu"`
	ReqMem         string `orm:"column(req_mem);size(100);" json:"req_mem"`
	LivenessProbe  string `orm:"column(liveness_probe);size(2048);" json:"liveness_probe"`
	ReadinessProbe string `orm:"column(readiness_probe);size(2048);" json:"readiness_probe"`
	CreateUser     string `orm:"column(create_user);size(1024);" json:"create_user"`
}

type DeployProjectManager struct{}

func NewDeployProjectManager() *DeployProjectManager {
	return &DeployProjectManager{}
}

func (m *DeployProjectManager) GetById(id int) *DeployProject {
	deploy := &DeployProject{}
	ormer := orm.NewOrm()
	err := ormer.QueryTable(deploy).Filter("Id__exact", id).One(deploy)
	if err == nil {
		return deploy
	}
	return nil
}

func (m *DeployProjectManager) GetByName(name string) *DeployProject {
	deployProject := &DeployProject{}
	err := orm.NewOrm().QueryTable(deployProject).Filter("Name__exact", name).One(deployProject)
	if err == nil {
		return deployProject
	}
	return nil
}

func (m *DeployProjectManager) Query(q string, start int64, length int, userName string) ([]*DeployProject, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&DeployProject{})
	condition := orm.NewCondition()
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	var deploy []*DeployProject

	if q == "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("id__icontains", q)
		query = query.Or("create_user__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	} else {
		q = userName
		query := orm.NewCondition()
		query = query.Or("create_user__icontains", q)

		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}

	queryset.SetCond(condition).Limit(length).Offset(start).All(&deploy)
	return deploy, total, qtotal
}

func (m *DeployProjectManager) Create(deploy *DeployProject) (*DeployProject, error) {
	ormer := orm.NewOrm()
	deploys := &DeployProject{
		Name:           deploy.Name,
		Type:           deploy.Type,
		Describe:       deploy.Describe,
		CreateUser:     deploy.CreateUser,
		Business:       deploy.Business,
		GitUrl:         deploy.GitUrl,
		BuildCmd:       deploy.BuildCmd,
		Port:           deploy.Port,
		SvcType:        deploy.SvcType,
		LimitMem:       deploy.LimitMem,
		LimitCpu:       deploy.LimitCpu,
		ReqCpu:         deploy.ReqCpu,
		ReqMem:         deploy.ReqMem,
		LivenessProbe:  deploy.LivenessProbe,
		ReadinessProbe: deploy.ReadinessProbe,
	}

	if _, err := ormer.Insert(deploys); err != nil {
		return nil, err
	}
	return deploys, nil
}

func (m *DeployProjectManager) Modify(id int, status, name, env, disposeUser, detail, remark string, doneTime *time.Time) (*DeployProject, error) {
	ormer := orm.NewOrm()
	if deploy := m.GetById(id); deploy != nil {
		deploy.Name = name

		if _, err := ormer.Update(deploy); err != nil {
			return nil, err
		}
		return deploy, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}

func (m *DeployProjectManager) Start(id int) error {
	ormer := orm.NewOrm()
	if deploy := m.GetById(id); deploy != nil {
		if _, err := ormer.Update(deploy); err != nil {
			return nil
		}
		return nil
	}
	return nil

}

func (m *UserManager) DeployManager(id int, name string, gender int, birthday *time.Time, tel, email, addr, remark string) (*User, error) {
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

// 发布服务
type DeployService struct {
	Id         int        `orm:"column(id);" json:"id"`
	Env        string     `orm:"column(env);size(32);" json:"env"`
	Name       string     `orm:"column(name);size(32);" json:"name"`
	Branch     string     `orm:"column(branch);size(32);" json:"branch"`
	Version    string     `orm:"column(version);size(32);" json:"version"`
	UpdateTime *time.Time `orm:"column(update_time);" json:"update_time"`
}
type DeployServiceManager struct{}

func NewDeployServiceManager() *DeployServiceManager {
	return &DeployServiceManager{}
}

func (m *DeployServiceManager) Query(q string, start int64, length int) ([]*DeployService, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&DeployService{})
	condition := orm.NewCondition()
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	var deploy []*DeployService

	if q == "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("id__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	} else {
		//q = userName
		query := orm.NewCondition()
		query = query.Or("create_user__icontains", q)

		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}

	queryset.SetCond(condition).Limit(length).Offset(start).All(&deploy)
	return deploy, total, qtotal
}

func (m *DeployServiceManager) GetById(id int) *DeployService {
	deploy := &DeployService{}
	ormer := orm.NewOrm()
	err := ormer.QueryTable(deploy).Filter("Id__exact", id).One(deploy)
	if err == nil {
		return deploy
	}
	return nil
}

func (m *DeployServiceManager) Modify(id int, branch, version string, deploytime *time.Time) (*DeployService, error) {
	ormer := orm.NewOrm()
	if deploy := m.GetById(id); deploy != nil {
		deploy.Branch = branch
		deploy.Version = version
		deploy.UpdateTime = deploytime
		if _, err := ormer.Update(deploy); err != nil {
			return nil, err
		}
		return deploy, nil
	}
	return nil, fmt.Errorf("操作对象不存在")
}



var DefaultDeployProjectManager = NewDeployProjectManager()
var DefaultDeployServiceManager = NewDeployServiceManager()

func init() {
	orm.RegisterModel(new(DeployProject)) //注册数据库，用户models执行使用
	orm.RegisterModel(new(DeployService)) //注册数据库，用户models执行使用
}
