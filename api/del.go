package api

import (
	"MottoGo/database"
	"MottoGo/global"
	"MottoGo/middleware"
	"MottoGo/models"
	"github.com/gin-gonic/gin"
)

func DelHit(r *gin.Engine) {
	r.DELETE("/hitokoto/DelHit/:uuid", func(c *gin.Context) {
		// 限流
		if !Ratelimit.Limit(c.ClientIP()) {
			c.JSON(429, gin.H{"code": 429, "msg": "Too many requests!", "data": nil})
			c.Abort()
			return
		}
		// 权限验证
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
		uuid := c.Param("uuid")
		if uuid == "" {
			c.JSON(422, gin.H{"code": 422, "msg": "Please enter uuid", "data": nil})
			c.Abort()
			return
		}
		// 依据uuid删除
		delhit := database.DB.Where("uuid = ?", uuid).Delete(&models.Hitokoto{})
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
			"data": gin.H{
				"delete_count": delhit.RowsAffected,
				"delete_uuid":  uuid,
			},
		})
		return
	})
}
