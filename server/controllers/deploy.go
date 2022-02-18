package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"gocmdb/server/jenkins"
	"strings"

	"gocmdb/server/controllers/auth"
	"gocmdb/server/forms"
	"gocmdb/server/models"
)


type DeployPageController struct {
	LayoutController
}

type DeployController struct {
	auth.LoginRequiredController
}

func (c *DeployPageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "deploy_management"
}


// 发布项目列表
func (c *DeployPageController) ProjectIndex() {
	c.Data["menu"] = "project_list_management"
	c.TplName = "deploy/index.html"
	c.LayoutSections["LayoutScript"] = "deploy/index.script.html"
}


func (c *DeployController) ProjectList() {
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))
	deploy, total, queryTotal := models.DefaultDeployProjectManager.Query(q, start, length, c.User.Name)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          deploy,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *DeployController) ProjectCreate() {
	if c.Ctx.Input.IsPost() {
		envs := c.Input().Get("envs")
		types := c.Input().Get("types")
		form := &forms.TicketControllerCreateForm{}
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
				deploy := &models.Ticket{
					Name: form.Name,
					Type: types,
					Env: envs,
					CreateUser: c.User.Name,
					DisposeUser: form.DisposeUser,
					Detail: form.Detail,
					Remark: form.Remark,
					Tel: form.Tel,
					Status: "指派中",
					CreatedTime: nil,
					DoneTime: form.DoneTimes,
				}
				result, err := models.DefaultTicketManager.Create(deploy)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
					fmt.Println("resutl:", result)
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
		c.TplName = "deploy/create.html"
		c.Data["svc_type"] = models.GetSvcType()
		c.Data["type"] = models.GetType()
		c.Data["envs"] = models.GetEnv()
		c.Data["jenkins"] = jenkins.DefaultJenkins.GetJobName()

	}
}

func (c *DeployController) ProjectDelete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultCloudPlatformManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *DeployController) ProjectModify()  {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		status := c.Input().Get("status")
		form := &forms.TicketControllerModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				fmt.Println("time:", form.DoneTime)
				deploy, err := models.DefaultTicketManager.Modify(form.Id,status, form.Name, form.Env,form.DisposeUser,form.Detail,form.Remark, form.DoneTime)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": deploy,
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
		c.TplName = "deploy/modify.html"
		c.Data["deploy"] = models.DefaultTicketManager.GetById(pk)
		c.Data["envs"] = models.GetEnv()
		c.Data["status"] = models.GetStatus()
	}
}


// 发布服务列表
func (c *DeployPageController) DeployIndex() {
	c.Data["menu"] = "deploy_list_management"
	c.TplName = "deploy/deploy.html"
	c.LayoutSections["LayoutScript"] = "deploy/deploy.script.html"
}

func (c *DeployController) DeployList() {
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	deploy, total, queryTotal := models.DefaultDeployServiceManager.Query(q, start, length)


	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          deploy,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *DeployController) DeployCreate() {
	if c.Ctx.Input.IsPost() {
		envs := c.Input().Get("envs")
		types := c.Input().Get("types")
		form := &forms.TicketControllerCreateForm{}
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
				deploy := &models.Ticket{
					Name: form.Name,
					Type: types,
					Env: envs,
					CreateUser: c.User.Name,
					DisposeUser: form.DisposeUser,
					Detail: form.Detail,
					Remark: form.Remark,
					Tel: form.Tel,
					Status: "指派中",
					CreatedTime: nil,
					DoneTime: form.DoneTimes,
				}
				result, err := models.DefaultTicketManager.Create(deploy)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": result,
					}
					fmt.Println("resutl:", result)
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
		c.TplName = "deploy/create.html"
		c.Data["envs"] = models.GetEnv()
	}
}

func (c *DeployController) DeployDelete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultCloudPlatformManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *DeployController) DeployModify()  {
	c.LayoutSections = map[string]string{
		"LayoutStyle":  "",
		"LayoutScript": "",
	}

	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.DeployControllerDeployModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				// 从deployProject获取应用的具体信息
				projectDetail := models.DefaultDeployProjectManager.GetByName(form.Name)
				fmt.Println(projectDetail)
				// 调用jenkins api发布
				parm := map[string]string{
					"version":form.Version,
					"app":form.Name,
					"branch":form.Branch,
				}
				fmt.Println("parm:", parm)
				jenkins.DefaultJenkins.Build(parm, projectDetail.JenkinsJob)
				//更新发布信息
				deploy, err := models.DefaultDeployServiceManager.Modify(form.Id,form.Branch,form.Version,form.DTime)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "发布成功",
						"result": deploy,
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
		c.Data["menu"] = "deploy_list_management"
		c.TplName = "deploy/deploy.modify.html"
		c.Data["deploy"] = models.DefaultDeployServiceManager.GetById(pk)
		c.Data["envs"] = models.GetEnv()
	}
}




func (c *DeployController)GetNs()  {
	allns[0] = map[string]string{
		"id": "0",
		"text":"default",
	}
	allns[1] = map[string]string{
		"id": "1",
		"text":"ee-sdk",
	}

	fmt.Println("allNs:", allns[0])
	c.Data["json"] = allns
	c.ServeJSON()
}