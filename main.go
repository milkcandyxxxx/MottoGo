package main

import (
	"MottoGo/api"
	"MottoGo/database"
	"MottoGo/models"
	"log"

	"github.com/gin-gonic/gin"
)

var Hit []models.Hitokoto
var configs models.Config
var port string
var key_admin []string
var key_user []string
var key_all []string

func init() {
	Hit = database.LoadHitokoto()
	configs = database.LoadConfig()
	port = configs.Server.Port
	key_admin = configs.Security.Key.Admin
	key_user = configs.Security.Key.User
	key_all = append(key_admin, key_user...)

	log.Println(configs.Security.Key)
}
func main() {
	r := gin.Default()
	api.Get(r, &Hit, key_all)
	api.AddHit(r, &Hit, key_admin)
	api.DelHit(r, Hit, key_admin)
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
