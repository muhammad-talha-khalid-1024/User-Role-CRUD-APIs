package main

import (
	"mustaqel/config"
	"mustaqel/models"
)

func main() {
	DB := config.ConnectDB()
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Role{})
}
