package task

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gocmdb/agents/config"
)

type Task interface {
	Init (config * config.AgentConfig)
	Run()
}

type manager map[string]Task

var mgr = make(manager)

func Register(name string, task Task)   {
	if _, ok := mgr[name]; ok {
		logrus.WithFields(logrus.Fields{
			"task": name,
		}).Fatal("task is exists")
	}
	mgr[name] = task
	logrus.WithFields(logrus.Fields{
		"task": name,
	}).Info("task is register")
}

func Run(config *config.AgentConfig, errChan chan <- error)  {
	for name, task := range mgr{
		task.Init(config)
		go func(name string, task Task) {
			logrus.WithFields(logrus.Fields{
				"task": name,
			}).Info("task is running ..")
			task.Run()

			logrus.WithFields(logrus.Fields{
				"task": name,
			}).Error("task is stopped")
			errChan <- fmt.Errorf("task %s is stopped", name)
		}(name, task)
	}
}