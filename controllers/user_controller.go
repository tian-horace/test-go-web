package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/user/gin-gorm-sqlite/database"
	"github.com/user/gin-gorm-sqlite/models"
	"net/http"
	"strconv"
)

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var user models.User

	// 绑定JSON请求体到user结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建用户记录
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {
	var users []models.User

	// 查询所有用户
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserByID 根据ID获取用户
func GetUserByID(c *gin.Context) {
	var user models.User

	// 获取URL参数中的id
	id := c.Param("id")

	// 查询用户
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var user models.User

	// 获取URL参数中的id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 检查用户是否存在
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 绑定JSON请求体到更新数据
	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新用户信息
	updateData.ID = uint(id)
	database.DB.Model(&user).Updates(updateData)

	// 获取更新后的用户信息
	database.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var user models.User

	// 获取URL参数中的id
	id := c.Param("id")

	// 检查用户是否存在
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 删除用户
	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": "用户删除成功"})
}
