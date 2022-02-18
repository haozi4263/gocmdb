package main

import (
	//guuid "github.com/satori/go.uuid"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/imroc/req"

	"gocmdb/agents/config"
	_ "gocmdb/agents/task/init"
	"gocmdb/agents/task"

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func initConfig(path string) *config.AgentConfig  {
	var config config.AgentConfig
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil{
		logrus.Fatal(err)
	}

	 if err := viper.Unmarshal(&config); err != nil{
		 logrus.Fatal(err)
	 }
	uuidPath := "promagent.uuid"
	if ctx, err := ioutil.ReadFile(uuidPath); err == nil{
		config.UUID = string(ctx)
	}else if os.IsNotExist(err){
		//fmt.Println(guuid.New())
		config.UUID = uuid.NewV4().String() //多个机器能保持随机
		ioutil.WriteFile(uuidPath, []byte(config.UUID), os.ModePerm)

	}else {
		logrus.Fatal(err)
	}



	return &config
}

func initLog(verbose bool,config *config.AgentConfig)  {
	logger := &lumberjack.Logger{
		Filename: config.LogConfig.Filename,
		MaxBackups: config.LogConfig.Maxbackups,
		MaxSize: config.LogConfig.Maxsize,
		Compress: config.LogConfig.Compress,
	}
	if verbose{
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true) // 设置打印代码行数和文件名字
		logrus.SetFormatter(&logrus.TextFormatter{}) // 日志格式为text
	}else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(logger)
	}
}


func main()  {
	var (
		verbose bool
		help bool
		path string
	)
	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&path, "path", "./etc/promagent.yaml","config path")

	flag.Usage = func() {
		fmt.Println("usage: promagent [ --verbose] [--path]")
		flag.PrintDefaults()
	}
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// 初始化配置
	config := initConfig(path)
	//fmt.Println("config:", config)
	// 初始化日志
	initLog(verbose, config)
	if verbose{
		req.Debug = true
	}
	// 日志打印测试
	//logrus.WithFields(logrus.Fields{
	//	"key":"val",
	//}).Debug("程序启动debug")
	//
	//logrus.WithFields(logrus.Fields{
	//	"agents":"info",
	//}).Info("程序启动info")

	// 启动程序
	stop := make(chan os.Signal,1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT) //监听kill ctrl + c

	errChan := make(chan error, 1)

	go func() {
		task.Run(config, errChan)
	}()
	logrus.WithFields(logrus.Fields{
		"pid":os.Getpid(),
	}).Info("promagent is running")
	select {
	case <- stop:
		logrus.Info("promagent stopped")
	case err := <- errChan:
		logrus.Error(err)
	}
	//reload := make(chan os.Signal, 1)
	//signal.Notify(reload, syscall.SIGHUP)
	//go func() {
	//	for  {
	//		<- reload
	//		fmt.Println("promagent reload")
	//	}
	//}()
}
