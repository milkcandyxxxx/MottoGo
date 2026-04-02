package api

import (
	"MottoGo/database"
	"MottoGo/global"
	"MottoGo/middleware"
	"MottoGo/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(r *gin.Engine) {
	r.GET("/hitokoto/all", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"code": 429, "msg": "Too many requests!", "data": nil})
			c.Abort()
			return
		}
		if global.Configs.Server.RequireUserkey {
			authorization := c.GetHeader("authorization")
			// 身份验证
			if !middleware.AuthKey(global.KeyAll, authorization) {
				c.JSON(401, gin.H{"msg": "Who are you?", "data": nil})
				c.Abort()
				return
			}
		}

		// 获取句子
		var hitokoto []models.Hitokoto
		err := database.DB.Find(&hitokoto).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{"code": 404, "msg": "No hitokoto found", "data": nil})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{"code": 200, "msg": "查询成功", "data": hitokoto})
	})
}
