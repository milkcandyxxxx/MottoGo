package database

import (
	"MottoGo/models"
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Load_hitokoto 加载句子
func Load_hitokoto() []models.Hitokoto {
	var hit []models.Hitokoto
	res, err := os.Open("./cartoon.jsonl")
	if err != nil {
		panic(err)
	}
	defer func(res *os.File) {
		err := res.Close()
		if err != nil {
		}
	}(res)
	// 使用 bufio.Scanner 逐行读取 JSONL
	scanner := json.NewDecoder(res)
	for {
		var item models.Hitokoto
		err := scanner.Decode(&item)
		if err != nil {
			break
		}
		hit = append(hit, item)
	}
	return hit
}

// Load_config 加载配置文件
func Load_config() models.Config {
	var config models.Config
	res, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = yaml.Unmarshal(res, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	log.Println("加载完成")
	return config
}
