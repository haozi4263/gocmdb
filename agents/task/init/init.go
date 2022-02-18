package init

import (
	"gocmdb/agents/task"
	"gocmdb/agents/task/plugins/profile"
	"gocmdb/agents/task/plugins/register"
)

func init()  {
	task.Register("register", &register.Register{})
	task.Register("profile", &profile.Profile{})
}
