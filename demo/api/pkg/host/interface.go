package host

import "context"

type Service interface {
	SaveHost(ctx context.Context, host *Host)(*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)

}


type QueryHostRequest struct {
	PageSize uin
}

// 主机列表
type HostSet struct {
	Items []*Host	`json:"items"`
	Total int `json:"total"`
}