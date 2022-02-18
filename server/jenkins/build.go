package jenkins

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
)

type Jenkins struct {
	url string
	username string
	password string
}

func NewJenkins() *Jenkins {
	return &Jenkins{}
}

var JenkinsObj  *gojenkins.Jenkins

func init()  {
	info := Jenkins{
		url: "http://127.0.0.1:8080/",
		username: "admin",
		password: "111111",
	}
	JenkinsObj = gojenkins.CreateJenkins(nil, info.url, info.username, info.password)
	JenkinsObj.Init(context.TODO())
}






func (j *Jenkins)Build (parm map[string]string, jobName string)  {
	qid, err := JenkinsObj.BuildJob(context.TODO(),jobName,parm)
	if err != nil {
		fmt.Println("build_err", err)
	}
	build, err := JenkinsObj.GetBuildFromQueueID(context.TODO(), qid)
	fmt.Println(build.GetConsoleOutput(context.TODO())) //打印当前构建信息,优化建议使用协程

	//build, _:=JenkinsObj.GetJob(context.TODO(), jobName, "/test")
	//info, err := build.GetLastBuild(context.TODO())
	//fmt.Println("err:", err)
	//fmt.Println(info.GetConsoleOutput(context.TODO()))

}

func (j *Jenkins)GetJobName () []*gojenkins.Job  {
	job, _ := JenkinsObj.GetAllJobs(context.TODO())
	return job
}


var DefaultJenkins = NewJenkins()

