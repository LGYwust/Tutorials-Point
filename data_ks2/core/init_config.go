package core

import (
	"data_ks2/config"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalln("read yaml err", err.Error())
	}
	err = yaml.Unmarshal(byteData, &c)
	if err != nil {
		log.Fatalln("解析 yaml err", err.Error())
	}
	return c
}
