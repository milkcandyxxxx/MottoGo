package database

import (
	"MottoGo/models"
	"crypto/md5"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

// ynToBool 将yn转换为bool
func ynToBool(input string) bool {
	return input == "y" || input == "yes"
}

// LoadConfig 加载配置文件
func LoadConfig() models.Config {
	var config models.Config
	if _, err := os.Stat("config.yaml"); errors.Is(err, os.ErrNotExist) {
		var port string      // 端口
		var allowCors string // 跨域
		var allowCorsBool bool
		var requireUserkey string
		var requireUserkeyBool bool
		var keyAdmin string   // admin的密码
		var keyUser string    // user的密码
		var limit string      // 限流
		var limitRate string  // 总量
		var limitBurst string // 每秒恢复
		log.Println("配置文件不存在，自动创建中ing")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("###请输入有效信息###")
		fmt.Print("程序允许端口:")
		fmt.Scanln(&port)
		fmt.Print("是否开启用户认证(y or n):")
		fmt.Scanln(&requireUserkey)
		if requireUserkey == "y" || requireUserkey == "n" {
			if ynToBool(requireUserkey) {
				requireUserkeyBool = true
			} else {
				requireUserkeyBool = false
			}
		}
		fmt.Print("限流器设置(y:默认 or n:自定义):")
		fmt.Scanln(&limit)
		if limit == "y" || limit == "n" {
			if !ynToBool(limit) {
				fmt.Print("每秒生成的令牌数:")
				fmt.Scanln(&limitRate)
				fmt.Print("最大令牌容量:")
				fmt.Scanln(&limitBurst)
			} else {
				limitRate = "10"
				limitBurst = "10"
			}
		}
		fmt.Print("是否允许跨域请求(y or n):")
		fmt.Scanln(&allowCors)
		if allowCors == "y" || allowCors == "n" {
			if ynToBool(allowCors) {
				allowCorsBool = true
			} else {
				allowCorsBool = false
			}
		}
		fmt.Print("管理员密码:")
		fmt.Scanln(&keyAdmin)
		// 使用反引号包裹，保持原样格式
		basicFormat := fmt.Sprintf(`server:
  port: %s # 运行端口
  allow_cors: %t    # 是否开启跨域访问 (CORS)
security:
  require_userkey: %t # 是否开启普通用户密钥，管理员为必须
  key:
    admin:
      - %s  # 管理员密钥
    user:
      - %s  # 用户密钥
limit:
  rate: %s   # 每秒允许生成的令牌数（QPS）
  burst: %s  # 桶的最大容量（允许瞬间爆发的请求数）
`, port, allowCorsBool, requireUserkeyBool, fmt.Sprintf("%x", md5.Sum([]byte(keyAdmin))), fmt.Sprintf("%x", md5.Sum([]byte(keyUser))), limitRate, limitBurst)
		fmt.Println(basicFormat)
		err := os.WriteFile("config.yaml", []byte(basicFormat), 0644)
		if err != nil {
			log.Fatal("写入失败")
		}
	}
	// 读取配置文件
	res, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = yaml.Unmarshal(res, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	log.Printf("%+v\n", config)
	return config
}
