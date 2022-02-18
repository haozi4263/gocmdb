package controllers

import (
	"fmt"
	"gocmdb/server/controllers/auth"
	"strings"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare() {
	c.LoginRequiredController.Prepare()
	c.Layout = "layouts/base.html"

	c.LayoutSections = map[string]string{
		"LayoutStyle":  "",
		"LayoutScript": "",
	}

	c.Data["menu"] = ""
	c.Data["expand"] = ""
}

func (c *LayoutController)GetMenu() string {
	controllerName, _ := c.GetControllerAndAction()
	controllerName = strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
	return fmt.Sprintf("%s_management", controllerName)

}
