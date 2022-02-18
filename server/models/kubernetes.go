package models

import (
	"context"
	"github.com/astaxie/beego"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type deployments struct {
}

func (c *deployments)List(draw, start, length int64, ns string) ([]appsV1.Deployment) {
	kubeconf := beego.AppConfig.DefaultString("k8s::path","conf/kube.conf")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconf)
	if err != nil{
		return []appsV1.Deployment{}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil{
		return []appsV1.Deployment{}
	}
	deploymentList, err := clientset.AppsV1().Deployments(ns).List(context.TODO(), metaV1.ListOptions{})
	if err != nil{
		return []appsV1.Deployment{}
	}
	if len(deploymentList.Items) < 10 {
		depNum := len(deploymentList.Items)
		return deploymentList.Items[0:depNum]
	}
	if length == 0{
		return deploymentList.Items[0:10]
	}
	return deploymentList.Items[start:length*draw]
}
var Deployment = new(deployments)
