package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"milksay/api"
	"milksay/database"
	"milksay/models"
)

var hit []models.Hitokoto
var configs models.Config
var port string
var key_admin []string
var key_user []string
var key_all []string

func init() {
	hit = database.Load_hitokoto()
	configs = database.Load_config()
	port = configs.Server.Port
	key_admin = configs.Security.Key.Admin
	key_user = configs.Security.Key.User
	key_all = append(key_admin, key_user...)

	log.Println(configs.Security.Key)
}
func main() {
	r := gin.Default()
	api.Get(r, hit, key_all)
	api.Add_hit(r, hit, key_admin)
	api.Del_hit(r, hit, key_admin)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
