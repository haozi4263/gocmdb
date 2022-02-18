package models

const (
	StatusUnlock = 0
	StatusLock   = 1
)


// 工单类型

func GetTicketType() []string {
	ticketType := []string{"数据库","主机","网络","权限"}
	return ticketType
}
// 工单所属环境
func GetEnv() []string {
	envs := []string{"dev","test","uat","pro"}
	return envs
}

// 工单处理状态
func GetStatus() []string {
	status := []string{"指派中","驳回","处理中","转交","完成","重新打开"}
	return status
}


// 发布
func GetSvcType() []string {
	svcType := []string{"ClusterIP","Headless","NoePort"}
	return svcType
}

func GetType() []string {
	Type := []string{"java","golang","nodejs","static"}
	return Type
}