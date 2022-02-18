package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocmdb/server/base/response"
	"gocmdb/server/controllers/base"
	"strings"
)

type APIController struct {
	base.BaseController
}

func (c *APIController) Prepare()  {
	c.EnableXSRF = false
	_,action := c.GetControllerAndAction()
	if strings.ToLower(action) == "alert"{ //alert去掉token认证
		c.Data["json"] = response.Ok
	}else {
		token := fmt.Sprintf("Token %s", beego.AppConfig.String("token"))
		headerToken := c.Ctx.Input.Header("Authorization")
		if token != headerToken{
			c.Data["json"] = response.Unauthorization
			c.ServeJSON()
		}
	}

}





