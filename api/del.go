package api

import (
	"MottoGo/global"
	"MottoGo/middleware"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func DelHit(r *gin.Engine) {
	r.GET("/hitokoto/DelHit", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too many requests!"})
			return
		}
		// 获取信息
		uuid := c.Query("uuid")
		key := c.GetHeader("X-API-key")

		// 权限验证
		if !middleware.AuthKey(global.KeyAll, key) {
			c.JSON(401, gin.H{"error": "Who are you?"})
			return
		}
		if !middleware.AuthKey(global.KeyAdmin, key) {
			c.JSON(403, gin.H{"error": "You cannot do it"})
			return
		}
		if uuid == "" {
			c.JSON(400, gin.H{"error": "UUID required"})
			return
		}

		// 全局搜索 UUID
		var targetCategory string
		var delIndex = -1
		var delHit string

		// 遍历所有分类
		for cat, hits := range global.Hit {
			for i, hit := range hits {
				if hit.Uuid == uuid {
					targetCategory = cat
					delIndex = i
					delHit = hit.Hitokoto
					break
				}
			}
			if delIndex != -1 {
				break
			}
		}

		if delIndex == -1 {
			c.JSON(404, gin.H{"error": "UUID not found in any category"})
			return
		}

		// 从全局变量中删除该元素
		global.Hit[targetCategory] = append(global.Hit[targetCategory][:delIndex], global.Hit[targetCategory][delIndex+1:]...)

		// 重写文件
		f, err := os.OpenFile("./cartoon.jsonl", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to open file: " + err.Error()})
			return
		}
		defer f.Close()

		// 写入文件
		isFirstLine := true
		for _, hits := range global.Hit {
			for _, item := range hits {
				byteHit, err := json.Marshal(item)
				if err != nil {
					continue
				}
				if !isFirstLine {
					f.WriteString("\n")
				}
				f.Write(byteHit)
				isFirstLine = false
			}
		}
		c.JSON(200, gin.H{
			"ok":       "Delete successful",
			"hitokoto": delHit,
		})
	})
}
