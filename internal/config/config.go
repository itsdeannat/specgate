package config

import (
	"os"
	"gopkg.in/yaml.v3"
	"log"
)


type RulesConfig struct {
	Rules struct {
        ServerBlockList []string `yaml:"server_block_list"`
    } `yaml:"rules"`
}

func CreateConfig () {

	defaultConfig := RulesConfig{
			Rules: struct {
				ServerBlockList []string `yaml:"server_block_list"`
			}{
				ServerBlockList: []string{"http://www.example.com"},
			},
		}

		data, err := yaml.Marshal(&defaultConfig)
		
		err = os.WriteFile(".specgate.yaml", []byte(data), 0644) 

		if err != nil {
			log.Fatal(err)
		}
}