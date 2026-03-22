package api

import (
	"MottoGo/database"
	"MottoGo/global"
	"MottoGo/middleware"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelHit(r *gin.Engine) {
	r.GET("/hitokoto/DelHit", func(c *gin.Context) {
		// 获取信息
		Id := c.Query("id")
		Key := c.GetHeader("X-API-Key")
		Id_int, err := strconv.Atoi(Id)
		// 验证信息合法
		if !middleware.AuthKey(global.KeyAll, Key) {
			c.JSON(401, gin.H{"error": "who are you?"})
			return
		} else {
			if !middleware.AuthKey(global.KeyAdmin, Key) {
				c.JSON(403, gin.H{"error": "you cannot do it"})
				return
			}
		}

		if err != nil {
			c.JSON(400, gin.H{"error": "id 错误"})
			return
		}
		if !(global.Hit[len(global.Hit)-1].Id >= Id_int) {
			c.JSON(400, gin.H{"error": "不存在 id"})
			return
		}

		del_hit := global.Hit[Id_int-1].Hitokoto
		global.Hit[len(global.Hit)-1].Id = global.Hit[Id_int-1].Id
		global.Hit[Id_int-1] = global.Hit[len(global.Hit)-1]
		global.Hit = global.Hit[:len(global.Hit)-1]
		f, err := os.OpenFile("./cartoon.jsonl", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			c.JSON(500, gin.H{"error": "打开文件失败：" + err.Error()})
			return
		}
		defer func(f *os.File) {
			_ = f.Close() // 忽略关闭错误，或按需处理
		}(f)

		// 逐行写入
		for i, item := range global.Hit {
			// 序列化单个Hitokoto为JSON字符串
			byteHit, err := json.Marshal(item)
			if err != nil {
				c.JSON(500, gin.H{"error": "格式化数据失败：" + err.Error()})
				return
			}

			// 如果不是第一个先换行
			if i > 0 {
				_, err = f.WriteString("\n")
				if err != nil {
					c.JSON(500, gin.H{"error": "写入换行符失败：" + err.Error()})
					return
				}
			}

			// 写入当前JSON行
			_, err = f.Write(byteHit)
			if err != nil {
				c.JSON(500, gin.H{"error": "写入数据失败：" + err.Error()})
				return
			}
		}
		database.LoadConfig()
		c.JSON(200, gin.H{"ok": "删除成功", "Hitokoto": del_hit})

	})
}
