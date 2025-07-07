package main

import (
	"crud/initializers"
	"crud/models"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
