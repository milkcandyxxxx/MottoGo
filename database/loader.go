package database

import (
	"MottoGo/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// LoadConfig 加载配置文件
func LoadConfig() models.Config {
	var config models.Config
	// 读取配置文件
	res, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	// 解析 YAML
	err = yaml.Unmarshal(res, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}
