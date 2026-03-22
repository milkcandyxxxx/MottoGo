package api

import (
	"MottoGo/global"
	"MottoGo/middleware"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

var Ratelimit = &middleware.Ratelimit{}

func Get(r *gin.Engine) {
	r.GET("/hitokoto", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too many requests!"})
			return
		}
		// 身份验证
		id := c.Query("id")
		if !middleware.SecurityVerification(c, global.KeyAll) {
			c.JSON(401, gin.H{"error": "who are you?"})
			return
		}
		// 根据 id 返回
		switch id {
		case "":
			hitokoto := global.Hit[rand.Intn(len(global.Hit))]
			c.JSON(200, hitokoto)
			return
		case "0":
			hitokoto := global.Hit
			c.JSON(http.StatusOK, hitokoto)
			return
		default:
			idInt, err := strconv.Atoi(id)
			if err != nil || idInt < 0 || idInt > len(global.Hit) {
				c.JSON(400, gin.H{"error": "Invalid ID"})
				return
			}
			c.JSON(200, global.Hit[idInt-1])
		}
	})
}
