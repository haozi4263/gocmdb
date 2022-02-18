package v1

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"gocmdb/server/base/response"
	"gocmdb/server/controllers/auth"
	"gocmdb/server/forms"
	"gocmdb/server/models"
	"gocmdb/server/services"
)

type PrometheusController struct {
	auth.APIController
}


func (c *PrometheusController) Register() {
	node := &forms.NodeRegisterForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &node); err == nil {
		services.DefaultPrometheusNode.Register(node)
		c.Data["json"] = response.Ok
	}else {
		fmt.Println("err:",err)
		c.Data["json"] = response.BadRequest
	}
	c.ServeJSON()
}

func (c *PrometheusController) Config()  {
	uuid := c.GetString("uuid")
	// job target
	jobs := models.DefaultJobManager.GetAllJobByUUID(uuid)
	c.Data["json"]= response.NewJSONResponse(200,"ok",jobs)
	c.ServeJSON()
}

func (c *PrometheusController) Alert() {
	fmt.Println("prometheus alert")
	fmt.Println(string(c.Ctx.Input.RequestBody))
	gjson.GetBytes(c.Ctx.Input.RequestBody, "alerts").ForEach(func(key, alert gjson.Result) bool {
		var form forms.AlertForm
		if err := json.Unmarshal([]byte(alert.Raw), &form); err == nil {
			fmt.Println("form",&form)
		} else {
			fmt.Println(err)
		}
		return true
	})
	c.Data["json"] = response.NewJSONResponse(200,"ok","alert")
	c.ServeJSON()

}