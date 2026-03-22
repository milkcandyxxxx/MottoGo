package main

import (
	"MottoGo/api"
	"MottoGo/database"
	"MottoGo/global"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.SetPrefix("[MottoGo] ")
	global.Hit = database.LoadHitokoto()
	global.Configs = database.LoadConfig()
	global.Port = global.Configs.Server.Port
	global.KeyAdmin = global.Configs.Security.Key.Admin
	global.KeyUser = global.Configs.Security.Key.User
	global.KeyAll = append(global.KeyAdmin, global.KeyUser...)
	log.Println("初始化完成")

}
func main() {
	r := gin.Default()
	api.Get(r)
	api.AddHit(r)
	api.DelHit(r)
	err := r.Run(":" + global.Port)
	if err != nil {
		return
	}
}
