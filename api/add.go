package api

import (
	"MottoGo/global"
	"MottoGo/middleware"
	"MottoGo/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
)

// AddHit 添加句子
func AddHit(r *gin.Engine) {
	r.POST("/hitokoto/AddHit", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"msg": "Too many requests!"})
			return
		}
		// 获取信息
		var hitokoto models.Hitokoto
		err := c.ShouldBindJSON(&hitokoto)
		XAPIKey := c.GetHeader("X-API-Key")
		if err != nil {
			c.JSON(400, gin.H{"msg": "Invalid JSON"})
			return
		}
		// AuthKey 进行验证
		if !middleware.AuthKey(global.KeyAdmin, XAPIKey) {
			c.JSON(401, gin.H{"msg": "Who are you?"})
			return
		}
		// 生成 UUID
		hitokoto.Uuid = uuid.New().String()
		byteHit, err := json.Marshal(hitokoto)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Format failed"})
			return
		}
		// 写入到 cartoon.jsonl 文件
		f, err := os.OpenFile("./cartoon.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Failed to open file"})
			return
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
			}
		}(f)
		// 如果文件不为空，先写入换行符
		fileInfo, err := f.Stat()
		if err == nil && fileInfo.Size() > 0 {
			_, err = f.WriteString("\n")
			if err != nil {
				c.JSON(500, gin.H{"msg": "Write failure"})
				return
			}
		}
		_, err = f.Write(byteHit)
		if err != nil {
			c.JSON(500, gin.H{"msg": "Write failure"})
			return
		}
		// 更新全局变量
		global.Hit[hitokoto.Type] = append(global.Hit[hitokoto.Type], hitokoto)
		c.JSON(200, gin.H{"msg": "Write success", "data": hitokoto})
	})
}
