package forms

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)



type DeployControllerDeployModifyForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name"`
	Env        string `form:"env"`
	Branch        string `form:"branch"`
	Version        string `form:"version"`
	DTime *time.Time
}

func (f *DeployControllerDeployModifyForm) Valid(v *validation.Validation) {

	f.Name = strings.TrimSpace(f.Name)
	f.Branch = strings.TrimSpace(f.Branch)
	f.Version = strings.TrimSpace(f.Version)
	f.Env = strings.TrimSpace(f.Env)

	str := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05 CST"))
	dtime, _ := time.Parse("2006-01-02 15:04:05 CST", str)

	f.DTime = &dtime


}
