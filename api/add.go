package api

import (
	"MottoGo/database"
	"MottoGo/middleware"
	"MottoGo/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func AddHit(r *gin.Engine, hit *[]models.Hitokoto, k []string) {
	r.POST("/hitokoto/AddHit", func(c *gin.Context) {
		var hitokoto models.Hitokoto
		err := c.ShouldBindJSON(&hitokoto)
		Authorization := c.GetHeader("Authorization")
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		// AuthKey 进行验证
		middleware.AuthKey(k, Authorization)
		// 获取Id
		hitokoto.Id = (*hit)[len(*hit)-1].Id + 1
		byteHit, err := json.Marshal(hitokoto)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to format data"})
			return
		}
		f, err := os.OpenFile("./cartoon.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.JSON(500, gin.H{"error": "打开文件失败"})
			return
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
			}
		}(f)
		if hitokoto.Id != 0 {
			_, err = f.WriteString("\n")
			if err != nil {
				c.JSON(500, gin.H{"error": "写入失败"})
				return
			}
		}
		_, err = f.Write(byteHit)
		if err != nil {
			c.JSON(500, gin.H{"error": "写入失败"})
			return
		}
		c.JSON(200, gin.H{"message": "写入成功", "data": hitokoto})
		*hit = database.LoadHitokoto()
		fmt.Println(*hit)
	})
}
