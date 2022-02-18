package main

import "fmt"

type authPlugin interface {
	Name() string
	Index(string)
}

type Manager struct {
	plugins map[string]authPlugin
}

func NewManager() *Manager {
	return &Manager{
		plugins: map[string]authPlugin{},
	}
}

func (m *Manager)Register(a authPlugin)  {
	m.plugins[a.Name()] = a
}

func (m *Manager)GetPlugin() authPlugin  {
	for _,plugin := range m.plugins{
		return plugin
	}
	return nil
}

func (m *Manager)GoIndex(name string)  {
	if plugins := m.GetPlugin(); plugins != nil {
		plugins.Index(name)
	}
}

var DefaultManager = NewManager()

// session实现
type Session struct {}
func (s *Session)Name() string  {
	return "session"
}

func (s *Session)Index(name string)  {
	for pluginName,_ := range DefaultManager.plugins{
		if pluginName == name {
			fmt.Println("run index in",name)
		}
	}
}

// token 实现

type Token struct {}
func (s *Token) Name() string  {
	return "token"
}

func (s *Token)Index(name string)  {
	for pluginName,_ := range DefaultManager.plugins{
		if pluginName == name {
			fmt.Println("run index in",name)
		}
	}
}




func init()  {
	//注册插件
	DefaultManager.Register(new(Session))
	DefaultManager.Register(new(Token))
}

func main()  {
	DefaultManager.GoIndex("session")
	DefaultManager.GoIndex("token")
	DefaultManager.GoIndex("api") //未注册不会运行
}