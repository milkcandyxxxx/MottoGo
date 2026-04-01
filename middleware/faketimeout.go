package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func FakeTimeout() gin.HandlerFunc {
	whiteList := map[string]bool{
		"/hitokoto":              true,
		"/hitokoto/all":          true,
		"/hitokoto/AddHit":       true,
		"/hitokoto/DelHit/:uuid": true,
	}

	return func(c *gin.Context) {
		// 获取注册时的路由模板,因为要识别的是:uuid
		path := c.FullPath()

		// 有没有在上面的表中
		if whiteList[path] {
			c.Next()
			return
		}

		log.Printf("非法访问拦截: %s 来自 %s", c.Request.URL.Path, c.ClientIP())

		select {
		// 对方不退出则延迟一分钟
		case <-time.After(60 * time.Second):
			c.AbortWithStatus(504)
		// 退出则恢复
		case <-c.Request.Context().Done():
			return
		}
	}
}
