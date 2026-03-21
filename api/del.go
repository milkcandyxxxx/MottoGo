package api

import (
	"MottoGo/middleware"
	"MottoGo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelHit(r *gin.Engine, hit []models.Hitokoto, k []string) {
	r.GET("/hitokoto/DelHit", func(c *gin.Context) {
		// 获取信息
		Id := c.Query("id")
		Authorization := c.GetHeader("Authorization")
		Id_int, err := strconv.Atoi(Id)
		// 验证信息合法
		if !middleware.AuthKey(k, Authorization) {
			c.JSON(400, gin.H{"error": "无权限"})
		}
		if err != nil {
			c.JSON(400, gin.H{"error": "id错误"})
			return
		}
		if !(hit[len(hit)-1].Id > Id_int) {
			c.JSON(400, gin.H{"error": "不存在id"})
			return
		}

		del_hit := hit[Id_int].Hitokoto
		hit[Id_int] = hit[len(hit)-1]
		hit = hit[:len(hit)-1]
		c.JSON(200, gin.H{"ok": "删除成功", "Hitokoto": del_hit})

	})
}
