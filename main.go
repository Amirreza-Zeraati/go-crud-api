package main

import (
	"crud/initializers"
	"github.com/gin-gonic/gin"
  "crud/controllers"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
}

func main() {
  r := gin.Default()
  r.POST("/user", controllers.UserCreate)
  r.PUT("/user/:id", controllers.UserUpdate)

  r.GET("/user", controllers.UsersShow)
  r.GET("/user/:id", controllers.UserShow)

  r.DELETE("/user/:id", controllers.UserDelete)
  
  r.Run()
}