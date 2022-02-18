package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"gocmdb/server/controllers/auth"
	"gocmdb/server/forms"
	"gocmdb/server/models"
	"strings"
)

type PrometheusNodePageController struct {
	LayoutController
}

func (c *PrometheusNodePageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "prometheus_management"
	c.Data["menu"] = c.GetMenu()
}

func (c *PrometheusNodePageController) Index() {
	c.TplName = "prometheus/node/index.html"
	c.LayoutSections["LayoutScript"] = "prometheus/node/index.script.html"
}

func (c *PrometheusNodePageController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultNodeManager.List(q, start, length)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

type PrometheusNodeController struct {
	auth.LoginRequiredController
}

func (c *PrometheusNodeController) Create()  {
	if c.Ctx.Input.IsPost() {
		form := &forms.NodeForm{}
		valid := &validation.Validation{}
		json := map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}

		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultNodeManager.Create(
					form.UUID,
					form.Hostname,
					form.Addr,
				)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "创建失败, 请重试",
						"result": nil,
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "prometheus/node/create.html"
	}
}

func (c *PrometheusNodeController) Modify()  {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.NodeForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				node, err := models.DefaultNodeManager.Modify(form.Id, form.Hostname, form.Addr)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": node,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()

	}else {
		pk, _ := c.GetInt("pk")
		c.TplName = "prometheus/node/modify.html"
		c.Data["info"] = models.DefaultNodeManager.GetById(pk)
	}
}

func (c *PrometheusNodeController) Delete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultNodeManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}

// Job
type PrometheusJobPageController struct {
	LayoutController
}

func (c *PrometheusJobPageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "prometheus_management"
	c.Data["menu"] = c.GetMenu()
}

func (c *PrometheusJobPageController) Index() {
	c.TplName = "prometheus/job/index.html"
	c.LayoutSections["LayoutScript"] = "prometheus/job/index.script.html"
}

func (c *PrometheusJobPageController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultJobManager.List(q, start, length)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

type PrometheusJobController struct {
	auth.LoginRequiredController
}

func (c *PrometheusJobController) Create()  {
	if c.Ctx.Input.IsPost() {
		form := &forms.JobForm{}
		valid := &validation.Validation{}
		json := map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}

		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			fmt.Println("formxx:", form.UUID)
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultJobManager.Create(
					form.UUID,
					form.Key,
					form.Remark,
				)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "创建失败, 请重试",
						"result": nil,
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "prometheus/job/create.html"
		c.Data["node"] = models.DefaultNodeManager.GetAllUUID()
	}
}

func (c *PrometheusJobController) Modify()  {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.JobForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				node, err := models.DefaultJobManager.Modify(form.Id, form.Key, form.Remark)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": node,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()

	}else {
		pk, _ := c.GetInt("pk")
		c.TplName = "prometheus/job/modify.html"
		c.Data["info"] = models.DefaultJobManager.GetById(pk)
	}
}

func (c *PrometheusJobController) Delete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultJobManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}



// Target

type PrometheusTargetPageController struct {
	LayoutController
}

func (c *PrometheusTargetPageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "prometheus_management"
	c.Data["menu"] = c.GetMenu()
}

func (c *PrometheusTargetPageController) Index() {
	c.TplName = "prometheus/target/index.html"
	c.LayoutSections["LayoutScript"] = "prometheus/target/index.script.html"
}

func (c *PrometheusTargetPageController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	fmt.Println(draw, start, length,q)

	result, total, queryTotal := models.DefaultTargetManager.List(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

type PrometheusTargetController struct {
	auth.LoginRequiredController
}

func (c *PrometheusTargetController) Create()  {
	if c.Ctx.Input.IsPost() {
		form := &forms.TargetForm{}
		valid := &validation.Validation{}
		json := map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}

		if err := c.ParseForm(form); err != nil {

			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			fmt.Println("forms::", form.JobKey)
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultTargetManager.Create(
					form.JobKey,
					form.Name,
					form.Addr,
				)
				fmt.Println("err", err)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "创建失败, 请重试",
						"result": nil,
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "prometheus/target/create.html"
		c.Data["job"] = models.DefaultNodeManager.GetAllJob()
	}
}

func (c *PrometheusTargetController) Modify()  {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.TargetForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				node, err := models.DefaultTargetManager.Modify(form.Id, form.Name, form.Addr)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": node,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()

	}else {
		pk, _ := c.GetInt("pk")
		c.TplName = "prometheus/target/modify.html"
		c.Data["info"] = models.DefaultTargetManager.GetById(pk)
	}
}

func (c *PrometheusTargetController) Delete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultTargetManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}
