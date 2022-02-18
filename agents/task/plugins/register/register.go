package register

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"gocmdb/agents/config"
	"os"
	"strings"
	"time"
)

type Register struct {
	config *config.AgentConfig
}

func (r *Register) Init(config *config.AgentConfig)  {
	r.config = config
}

func (r *Register) Run()  {
	ticker := time.NewTicker(r.config.TaskConfig.RegisterConfig.Interval)
	defer ticker.Stop()
	// UUID ADDR Hostname
	hostname,_ := os.Hostname()
	api := fmt.Sprintf("%s/v1/prometheus/register", strings.TrimRight(r.config.ServerConfig.Addr, "/"))
	fmt.Println("api::", api)
	params := req.Param{
		"uuid": r.config.UUID,
		"addr": r.config.Addr,
		"hostname": hostname,
	}
	fmt.Println("params:", params)
	fmt.Println("token::",r.config.ServerConfig.Token)
	headers := req.Header{
		"Authorization": fmt.Sprintf("Token %s",r.config.ServerConfig.Token),
		//"Accept": "application/json",
	}


	for{
		//ticker 程序启动立马发送注册信息，不先等待5s
		if resp, err := req.Post(api, req.BodyJSON(&params), headers); err != nil{
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("error request regiser")
		}else {
			body,_ := resp.ToString()
			logrus.WithFields(logrus.Fields{
				"response": body,// 返回json字符串
			}).Debug("success register")
		}
		fmt.Println("发送注册信息")
		<- ticker.C
	}
}

