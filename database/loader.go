package database

import (
	"MottoGo/models"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

// LoadHitokoto 加载句子，按 Type 字段分类
func LoadHitokoto() map[string][]models.Hitokoto {
	hitMap := make(map[string][]models.Hitokoto)
	// 读取文件
	res, err := os.Open("./cartoon.jsonl")
	if err != nil {
		log.Printf("无法打开文件：%v", err)
		return hitMap
	}
	defer res.Close()
	// 逐行解析JSONL
	decoder := json.NewDecoder(res)
	for {
		var item models.Hitokoto
		err := decoder.Decode(&item)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("解析 JSON 行出错：%v", err)
			break
		}
		// 按 Type 分类
		if item.Type != "" {
			hitMap[item.Type] = append(hitMap[item.Type], item)
		}

	}

	return hitMap
}

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
