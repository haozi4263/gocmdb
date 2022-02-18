package profile

import (
	"gocmdb/agents/utils"
	"os"

	"gopkg.in/yaml.v2"
)

type TargetConfig struct {
	Targets []string `yaml:"targets`
}

func NewTargetConfig(targets ...string) *TargetConfig {
	return &TargetConfig{targets}
}

func writeTarget(path string, targets []*Target) error {
	utils.MkPdir(path)
	f, err := os.Create(path)
	if err != nil{
		return err
	}
	defer f.Close()
	addrs := make([]string, len(targets))
	for i, target := range targets{
		addrs[i] = target.Addr
	}
	encoder := yaml.NewEncoder(f)
	config := []*TargetConfig{NewTargetConfig(addrs...)}
	return encoder.Encode(&config)
}
