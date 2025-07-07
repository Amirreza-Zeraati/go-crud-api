package controllers

import (
	"crud/initializers"
	"crud/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var userBase struct {
		Name string
		Email string
		Age int
	}
	c.Bind(&userBase)
	newUser := models.User{Name: userBase.Name, Age: userBase.Age, Email: userBase.Email}
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"msg": "user created successfully",
		"user": newUser,
	})
}

func UsersShow(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg": "Show all users",
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	initializers.DB.First(&user, id)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg": "Show specific user",
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	var userBase struct {
		Name string
		Email string
		Age int
	}
	c.Bind(&userBase)
	
	var user models.User
	id := c.Param("id")
	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).Updates(models.User{Name: userBase.Name, Age: userBase.Age, Email: userBase.Email})
	c.JSON(200, gin.H{
		"status": "ok",
		"msg": "user updated successfully",
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.User{}, id)
	c.JSON(200, gin.H{
		"status": "ok",
		"msg": "user deleted successfully",
	})
}