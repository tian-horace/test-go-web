package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/gin-gorm-sqlite/controllers"
	"github.com/user/gin-gorm-sqlite/database"
	"log"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 创建Gin引擎
	r := gin.Default()

	// 用户路由组
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/", controllers.CreateUser)      // 创建用户
		userGroup.GET("/", controllers.GetUsers)         // 获取所有用户
		userGroup.GET("/:id", controllers.GetUserByID)   // 获取单个用户
		userGroup.PUT("/:id", controllers.UpdateUser)    // 更新用户
		userGroup.DELETE("/:id", controllers.DeleteUser) // 删除用户
	}

	// 启动服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
