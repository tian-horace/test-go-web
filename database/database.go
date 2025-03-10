package database

import (
	"github.com/user/gin-gorm-sqlite/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gin-gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移模型
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("自动迁移模型失败: %v", err)
	}

	log.Println("数据库初始化成功")
}
