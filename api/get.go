package api

import (
	"MottoGo/database"
	"MottoGo/global"
	"MottoGo/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Ratelimit = &middleware.Ratelimit{}

func Get(r *gin.Engine) {
	r.GET("/hitokoto", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"code": 429, "msg": "Too many requests!", "data": nil})
			c.Abort()
			return
		}
		authorization := c.GetHeader("authorization")
		// 身份验证
		if !middleware.AuthKey(global.KeyAll, authorization) {
			c.JSON(401, gin.H{"msg": "Who are you?", "data": nil})
			c.Abort()
			return
		}
		// 获取句子
		hitokoto, err := database.GetRandomHitokoto(c)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(500, gin.H{"code": 404, "msg": "No hitokoto found", "data": nil})
			c.Abort()
			return
		}

		c.JSON(200, gin.H{"code": 200, "msg": "查询成功", "data": hitokoto})
	})
}
