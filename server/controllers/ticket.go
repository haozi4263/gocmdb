package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"

	"gocmdb/server/controllers/auth"
	"gocmdb/server/forms"
	"gocmdb/server/models"
)

var menu  string

type TicketPageController struct {
	LayoutController
}

func (c *TicketPageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "ticket_management"

}

func (c *TicketPageController) Index() {
	c.Data["menu"] = "my_ticket_management"
	menu = "my_ticket_management"
	c.TplName = "ticket/index.html"
	c.LayoutSections["LayoutScript"] = "ticket/index.script.html"
}

func (c *TicketPageController) IndexAll() {
	c.Data["menu"] = "all_ticket_management"
	menu = "all_ticket_management"
	c.TplName = "ticket/index.html"
	c.LayoutSections["LayoutScript"] = "ticket/index.script.html"
}

type TicketController struct {
	auth.LoginRequiredController
}

func (c *TicketController) List() {
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))
	ticket, total, queryTotal := models.DefaultTicketManager.Query(q,menu, start, length, c.User.Name)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          ticket,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *TicketController) Create() {
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
				ticket := &models.Ticket{
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
				result, err := models.DefaultTicketManager.Create(ticket)
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
		c.TplName = "ticket/create.html"
		c.Data["types"] = models.GetTicketType()
		c.Data["envs"] = models.GetEnv()
	}
}

func (c *TicketController) Delete() {
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

func (c *TicketController) Modify()  {
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
				ticket, err := models.DefaultTicketManager.Modify(form.Id,status, form.Name, form.Env,form.DisposeUser,form.Detail,form.Remark, form.DoneTime)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": ticket,
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
		c.TplName = "ticket/modify.html"
		c.Data["ticket"] = models.DefaultTicketManager.GetById(pk)
		c.Data["envs"] = models.GetEnv()
		c.Data["status"] = models.GetStatus()
	}
}

func (c *TicketController) Disable()  {
	pk, _ := c.GetInt("pk")
	if err := models.DefaultCloudPlatformManager.Disable(pk); err == nil{
		c.Data["json"] = map[string]interface{} {
			"code":   200,
			"text":   "云账号禁用成功",
			"result": nil, //可以返回删除的用户
		}
	}else {
		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "云账号禁用失败",
			"result": nil, //可以返回删除的用户
		}
	}
	c.ServeJSON()
}

func (c *TicketController) Done()  {
	pk, _ := c.GetInt("pk")
	if err := models.DefaultTicketManager.Done(pk); err == nil{
		c.Data["json"] = map[string]interface{} {
			"code":   200,
			"text":   "工单处理完成",
			"result": nil, //可以返回删除的用户
		}
	}else {

		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "工单处理失败",
			"result": nil, //可以返回删除的用户
		}
	}
	c.ServeJSON()
}


func (c *TicketController) Start()  {
	pk, _ := c.GetInt("pk")
	if err := models.DefaultTicketManager.Start(pk); err == nil{
		c.Data["json"] = map[string]interface{} {
			"code":   200,
			"text":   "开始接单",
			"result": nil, //可以返回删除的用户
		}
	}else {

		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "接单失败",
			"result": nil, //可以返回删除的用户
		}
	}
	c.ServeJSON()
}

func (c *TicketController)GetNs()  {
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