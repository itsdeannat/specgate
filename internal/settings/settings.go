package settings

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type SpecGateConfig struct {
	ServerBlockList []string `yaml:"server_block_list"`
}

func CreateConfig() {

	defaultConfig := SpecGateConfig{
		ServerBlockList: []string{"http://www.example.com"},
	}

	data, err := yaml.Marshal(&defaultConfig)

	err = os.WriteFile(".specgate.yaml", []byte(data), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
