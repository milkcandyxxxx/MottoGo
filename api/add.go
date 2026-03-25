package api

import (
	"MottoGo/database"
	"MottoGo/global"
	"MottoGo/middleware"
	"MottoGo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// AddHit 添加句子
func AddHit(r *gin.Engine) {
	r.POST("/hitokoto/AddHit", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"code": 429, "msg": "Too many requests!", "data": nil})
			c.Abort()
			return
		}

		// 身份验证
		authorization := c.GetHeader("Authorization")
		if !middleware.AuthKey(global.KeyAll, authorization) {
			c.JSON(401, gin.H{"code": 401, "msg": "Who are you?", "data": nil})
			c.Abort()
			return
		}
		if !middleware.AuthKey(global.KeyAdmin, authorization) {
			c.JSON(403, gin.H{"code": 403, "msg": "You cannot do it", "data": nil})
			c.Abort()
			return
		}
		// 获取信息
		var hitokoto models.Hitokoto
		err := c.ShouldBindJSON(&hitokoto)
		if err != nil {
			c.JSON(400, gin.H{"code": 400, "msg": "Invalid JSON", "data": nil})
			c.Abort()
			return
		}
		// 生成 UUID
		hitokoto.Uuid = uuid.New().String()
		// 写入句子
		result := database.DB.Create(&hitokoto)
		if result.Error != nil {
			log.Printf("插入失败: %v", result.Error)
			c.Abort()
			return
		}
		c.JSON(201, gin.H{"code": 201, "msg": "添加成功", "data": hitokoto})
	})
}
