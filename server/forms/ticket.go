package forms

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

type TicketControllerCreateForm struct {
	Name        string `form:"name"`
	DisposeUser string `form:"dispose_user"`
	Detail      string `form:"detail"`
	Tel         string `form:"tel"`
	Remark      string `form:"remark"`
	Done        string `form:"done_time"`

	DoneTimes *time.Time
}

func (f *TicketControllerCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.DisposeUser = strings.TrimSpace(f.DisposeUser)
	f.Detail = strings.TrimSpace(f.Detail)
	f.Tel = strings.TrimSpace(f.Tel)
	f.Remark = strings.TrimSpace(f.Remark)

	str := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05 MST"))
	if done, err := time.Parse("2006-01-02 15:04:05 MST", str); err != nil {
		v.SetError("birthday", "完成日期不正确")
	} else {
		f.DoneTimes = &done
	}

	v.MinSize(f.Name, 3, "name.name").Message("名字长度必须在%d-%d之内", 3, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("名字长度必须在%d-%d之内", 3, 32)

	v.MinSize(f.DisposeUser, 1, "dispose_user.dispose_user").Message("创建者为空且长度必须在%d之内", 64)
	v.MaxSize(f.DisposeUser, 64, "dispose_user.dispose_user").Message("创建者不能为空且长度必须在%d之内", 64)


	v.MinSize(f.Tel, 1, "tel.tel").Message("tel不能为空且长度必须在%d之内", 1024)
	v.MaxSize(f.Tel, 1024, "tel.tel").Message("tel不能为空不能为空且长度必须在%d之内", 1024)

	v.MinSize(f.Detail, 1, "detail.detail").Message("detail%d之内", 1024)
	v.MaxSize(f.Detail, 1024, "detail.detail").Message("detail%d之内", 1024)

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("备注长度必须在%d之内", 1024)

}

type TicketControllerModifyForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name"`
	DisposeUser string `form:"dispose_user"`
	Env         string `form:"env"`
	Detail      string `form:"detail"`
	Remark      string `form:"remark"`
	Done        string `form:"done_time"`

	DoneTime *time.Time
}

func (f *TicketControllerModifyForm) Valid(v *validation.Validation) {

	f.Name = strings.TrimSpace(f.Name)
	f.DisposeUser = strings.TrimSpace(f.DisposeUser)
	f.Env = strings.TrimSpace(f.Env)
	f.Detail = strings.TrimSpace(f.Detail)
	f.Remark = strings.TrimSpace(f.Remark)

	if done, err := time.Parse("2006-01-02 15:04:05 -0700 MST", f.Done); err != nil {
		v.SetError("birthday", "完成日期不正确")
	} else {
		f.DoneTime = &done
	}

	v.MinSize(f.Name, 3, "name.name").Message("名字长度必须在%d-%d之内", 3, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("名字长度必须在%d-%d之内", 3, 32)

	v.MinSize(f.DisposeUser, 1, "dispose_user.dispose_user").Message("区域不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.DisposeUser, 64, "dispose_user.dispose_user").Message("区域不能为空且长度必须在%d之内", 64)

	v.MinSize(f.Env, 1, "env.env").Message("env不能为空且长度必须在%d之内", 1024)
	v.MaxSize(f.Env, 1024, "env.env").Message("env不能为空不能为空且长度必须在%d之内", 1024)

	v.MinSize(f.Detail, 1, "detail.detail").Message("detail%d之内", 1024)
	v.MaxSize(f.Detail, 1024, "detail.detail").Message("detail%d之内", 1024)

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("备注长度必须在%d之内", 1024)

}
