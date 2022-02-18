package profile

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"gocmdb/agents/config"
	"strings"
	"time"
)

type Profile struct {
	config *config.AgentConfig
}

func (r *Profile) Init(config *config.AgentConfig)  {
	r.config = config
}

func (r *Profile) Run()  {
	ticker := time.NewTicker(r.config.TaskConfig.Profile.Interval)
	defer ticker.Stop()
	// UUID
	api := fmt.Sprintf("%s/v1/prometheus/config", strings.TrimRight(r.config.ServerConfig.Addr, "/"))
	params := req.Param{
		"uuid": r.config.UUID,
	}
	fmt.Println("params:", params)
	fmt.Println("token::",r.config.ServerConfig.Token)
	headers := req.Header{
		"Authorization": fmt.Sprintf("Token %s",r.config.ServerConfig.Token),
	}


	for{
		//ticker 程序启动立马发送注册信息，不先等待5s
		if resp, err := req.Get(api, params, headers); err != nil{
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("error request config")
		}else {
			var req Response
			if err := resp.ToJSON(&req); err == nil{
				logrus.WithFields(logrus.Fields{
					"jobs": req.Result ,// 返回json字符串
				}).Debug(" jobs josn succfull")
				writePrometheus(r.config.TaskConfig.Profile.Tpl, r.config.TaskConfig.Profile.Output, req.Result)

			}else {
				body,_ := resp.ToString()
				logrus.WithFields(logrus.Fields{
					"response": body,// 返回json字符串
				}).Error("json unmarshal josn profile")
			}
		}
		fmt.Println("获取配置信息")
		<- ticker.C
	}
}

