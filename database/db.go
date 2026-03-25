package database

import (
	"MottoGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// DB DB驱动
var DB *gorm.DB

// DBConnect 数据库连接
func DBConnect() {
	var err error
	// 连接数据库文件
	dbPath := "./cartoon.db"
	// 若没有文件则自动创建
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	// 如果格式不符则自动格式化
	err = DB.AutoMigrate(&models.Hitokoto{})
	if err != nil {
		log.Fatalf("数据初始化: %v", err)
	}
	log.Println("数据库初始化成功")
}

// GetRandomHitokoto 按照需求随机获得句子
func GetRandomHitokoto(c *gin.Context) (models.Hitokoto, error) {
	category := c.Query("c")
	author := c.Query("a")
	source := c.Query("s")
	query := DB.Model(&models.Hitokoto{})
	var item models.Hitokoto
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}
	if source != "" {
		query = query.Where("source = ?", source)
	}
	err := query.Order("RANDOM()").First(&item).Error
	return item, err
}
