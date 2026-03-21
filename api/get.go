package api

import (
	"MottoGo/middleware"
	"MottoGo/models"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Ratelimit = &middleware.Ratelimit{}

func Get(r *gin.Engine, hit *[]models.Hitokoto, k []string) {
	r.GET("/hitokoto", func(c *gin.Context) {
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too many requests!"})
			return
		}
		idOld := c.DefaultQuery("id", "0")
		if !middleware.SecurityVerification(c, k) {
			c.JSON(401, gin.H{"error": "who are you?"})
			return
		}
		if idOld == "0" {
			hitokoto := (*hit)[rand.Intn(len(*hit))]
			c.JSON(200, hitokoto)
			return
		}
		id, err := strconv.Atoi(idOld)
		if err != nil || id < 0 || id >= len(*hit) {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		c.JSON(200, (*hit)[id])
	})
}
