package forms

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

// Node
type NodeForm struct {
	Id       int    `form:"id"`
	UUID     string `form:"uuid"`
	Hostname string `form:"hostname"`
	Addr     string `form:"addr"`
}

func (f *NodeForm) Valid(v *validation.Validation) {
	f.Hostname = strings.TrimSpace(f.Hostname)
	f.Addr = strings.TrimSpace(f.Addr)

	v.AlphaDash(f.Hostname, "hostname.hostname").Message("主机名只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Hostname, 5, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Hostname, 32, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.Addr, 64, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)

}

type NodeRegisterForm struct {
	ID       int        `orm:"column(id);" json:"id"`
	UUID     string     `orm:"column(uuid);varcher(64)" json:"uuid"`
	Hostname string     `orm:"varchar(64)" json:"hostname"`
	Addr     string     `orm:"varchar(512)" json:"addr"`
	CreateAt *time.Time `orm:"auto_now_add" json:"create_at"`
	UpdateAt *time.Time `orm:"auto_now" json:"update_at"`
	DeleteAt *time.Time `orm:"null" json:"delete_at"`
}

func (f *NodeRegisterForm) Valid(v *validation.Validation) {
	f.Hostname = strings.TrimSpace(f.Hostname)
	f.Addr = strings.TrimSpace(f.Addr)

	v.AlphaDash(f.Hostname, "hostname.hostname").Message("主机名只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Hostname, 5, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Hostname, 32, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.Addr, 64, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)

}

//Job
type JobForm struct {
	Id     int    `form:"id"`
	UUID   string `form:"uuid"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
}

func (f *JobForm) Valid(v *validation.Validation) {
	f.Key = strings.TrimSpace(f.Key)
	f.Remark = strings.TrimSpace(f.Remark)

	v.AlphaDash(f.Key, "key.key").Message("key只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Key, 5, "key.key").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Key, 32, "key.key").Message("名字长度必须在%d-%d之内", 5, 32)

	v.MinSize(f.Remark, 1, "remark.remark").Message("备注不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.Remark, 64, "remark.remark").Message("备注不能为空且长度必须在%d之内", 64)

}

//Target
type TargetForm struct {
	JobKey string    `form:"job_key"`
	Id    int    `form:"id"`
	Name  string `form:"name"`
	Addr  string `form:"addr"`
}

func (f *TargetForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Addr = strings.TrimSpace(f.Addr)

	v.AlphaDash(f.Name, "name.name").Message("名字只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Name, 2, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.Addr, 64, "addr.addr").Message("地址不能为空且长度必须在%d之内", 64)

}

type AlertForm struct {
	Fingerprint  string            `json:"fingerprint"`
	Status       string            `json:"status"`
	StartsAt     *time.Time        `json:"startsAt"`
	EndsAt       *time.Time        `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
}