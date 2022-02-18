package controllers

import (
	"fmt"
	"gocmdb/server/models"
	"strings"
)
var allns = []map[string]string{}
var ns string

type K8sDeploymentPageController struct {
	LayoutController
}

func (c *K8sDeploymentPageController) Prepare()  {
	c.LayoutController.Prepare()
	c.Data["expand"] = "kubernetes_management"
	c.Data["menu"] = c.GetMenu()
}

func (c *K8sDeploymentPageController) Index() {
	c.TplName = "kubernetes/deployment/index.html"
	c.LayoutSections["LayoutScript"] = "kubernetes/deployment/index.script.html"
}

func (c *K8sDeploymentPageController) List() {
	draw, _ := c.GetInt64("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt64("length")
	nsIndex,_ := c.GetInt64("ns")
	if len(allns) == 0{
		allns = make([]map[string]string, 2, 20)
		allns[0] = map[string]string{
			"id": "0",
			"text":"default",
		}
	}
	var ns string
	q := strings.TrimSpace(c.GetString("q"))
		fmt.Println(draw, start, length,q)
	for _,v := range allns[nsIndex]{
		ns = v
	}
	deployment := models.Deployment.List(draw, start, length, ns)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "deploymnet获取成功",
		"result":          deployment,
	}
	c.ServeJSON()
}

func (c *K8sDeploymentPageController)GetNs()  {
	//allns = make([]map[string]string, 1, 20)
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
type DeploymentController struct {
	K8sDeploymentPageController
}










