package api

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"milksay/middleware"
	"milksay/models"
	"strconv"
)

var Ratelimit = &middleware.Ratelimit{}

func Get(r *gin.Engine, hit []models.Hitokoto, k []string) {
	r.GET("/hitokoto", func(c *gin.Context) {
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too many requests!"})
			return
		}
		id_old := c.DefaultQuery("id", "0")
		if !middleware.Security_verification(c, k) {
			c.JSON(401, gin.H{"error": "who are you?"})
			return
		}
		if id_old == "0" {
			hitokoto := hit[rand.Intn(len(hit))]
			c.JSON(200, hitokoto)
			return
		}
		id, err := strconv.Atoi(id_old)
		if err != nil || id < 0 || id >= len(hit) {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		c.JSON(200, hit[id])
	})
}
