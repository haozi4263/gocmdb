package services

import (
	"github.com/astaxie/beego/orm"
	"gocmdb/server/forms"
	"gocmdb/server/models"
)

type PrometheusNode struct {
}

func NewPrometheusNode() *PrometheusNode {
	return &PrometheusNode{}
}
func (n *PrometheusNode) Register(nodeform *forms.NodeRegisterForm) *models.Node  {
	ormer := orm.NewOrm()
	node := &models.Node{UUID: nodeform.UUID}
	if err := ormer.Read(node, "UUID"); err == nil {
		//有数据更新
		node.Hostname = nodeform.Hostname
		node.Addr = nodeform.Addr
		node.DeleteAt = nil
		ormer.Update(node)
	} else if err == orm.ErrNoRows{
		//无数据创建
		node.Hostname = nodeform.Hostname
		node.Addr = nodeform.Addr
		ormer.Insert(node)
	}else {
		return nil
	}
	return node
}


var DefaultPrometheusNode = NewPrometheusNode()