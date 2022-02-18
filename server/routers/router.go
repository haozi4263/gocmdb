package routers

import (
	"github.com/astaxie/beego"
	v1 "gocmdb/server/controllers/api/v1"

	"gocmdb/server/controllers"
	"gocmdb/server/controllers/auth"
)

func init() {
	// 认证
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	// 认证
	beego.AutoRouter(&auth.AuthController{})

	// 用户页面
	beego.AutoRouter(&controllers.UserPageController{})

	// 用户
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})

	// 云平台页面
	beego.AutoRouter(&controllers.CloudPlatformPageController{})

	// 云平台
	beego.AutoRouter(&controllers.CloudPlatformController{})

	// 云主机页面
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	// 云主机
	beego.AutoRouter(&controllers.VirtualMachineController{})


	// 工单
	beego.AutoRouter(&controllers.TicketPageController{})
	beego.AutoRouter(&controllers.TicketController{})

	// 发布管理
	beego.AutoRouter(&controllers.DeployPageController{})
	beego.AutoRouter(&controllers.DeployController{})

	// Prometheus Node管理
	beego.AutoRouter(&controllers.PrometheusNodePageController{})
	beego.AutoRouter(&controllers.PrometheusNodeController{})
	// Prometheus Job管理
	beego.AutoRouter(&controllers.PrometheusJobPageController{})
	beego.AutoRouter(&controllers.PrometheusJobController{})
	// Prometheus Target管理
	beego.AutoRouter(&controllers.PrometheusTargetPageController{})
	beego.AutoRouter(&controllers.PrometheusTargetController{})


	// K8s
	beego.AutoRouter(&controllers.K8sDeploymentPageController{})
	beego.AutoRouter(&controllers.DeploymentController{})

	v1 := beego.NewNamespace("/v1", beego.NSAutoRouter(&v1.PrometheusController{}))
	beego.AddNamespace(v1)
}
