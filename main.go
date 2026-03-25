package main

import (
	"MottoGo/api"
	"MottoGo/database"
	"MottoGo/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.SetPrefix("[MottoGo] ")
	// global.Hit = database.LoadHitokoto()
	global.Configs = database.LoadConfig()
	global.Port = global.Configs.Server.Port
	global.KeyAdmin = global.Configs.Security.Key.Admin
	global.KeyUser = global.Configs.Security.Key.User
	global.KeyAll = append(global.KeyAdmin, global.KeyUser...)
	database.DBConnect()
	log.Println("初始化完成")
}
func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api.Get(r)
	api.AddHit(r)
	api.DelHit(r)
	err := r.Run(":" + global.Port)
	if err != nil {
		return
	}
}
