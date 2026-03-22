package api

import (
	"MottoGo/global"
	"MottoGo/middleware"
	"MottoGo/models"
	"github.com/gin-gonic/gin"
	"math/rand"
)

var Ratelimit = &middleware.Ratelimit{}

func Get(r *gin.Engine) {
	r.GET("/hitokoto", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too many requests!"})
			return
		}
		// 获取信息
		category := c.Query("c")
		all := c.Query("all")
		// 身份验证
		if !middleware.SecurityVerification(c, global.KeyAll) {
			c.JSON(401, gin.H{"error": "Who are you?"})
			return
		}
		// 返回所有句子
		if all == "true" {
			if category != "" {
				// 指定分类
				if hits, ok := global.Hit[category]; ok {
					c.JSON(200, hits)
				} else {
					c.JSON(200, []models.Hitokoto{})
				}
			} else {
				// 所有分类
				c.JSON(200, global.Hit)
			}
			return
		}

		// 随机返回句子
		if category != "" {
			if hits, ok := global.Hit[category]; ok && len(hits) > 0 {
				hitokoto := hits[rand.Intn(len(hits))]
				c.JSON(200, hitokoto)
			} else {
				c.JSON(200, gin.H{"hitokoto": "No data", "from": "", "from_who": "", "type": category})
			}
		} else {
			allHits := []models.Hitokoto{}
			for _, hits := range global.Hit {
				allHits = append(allHits, hits...)
			}
			if len(allHits) > 0 {
				hitokoto := allHits[rand.Intn(len(allHits))]
				c.JSON(200, hitokoto)
			} else {
				c.JSON(200, gin.H{"hitokoto": "No data", "from": "", "from_who": "", "type": ""})
			}
		}
	})
}
