package main

import (
	"zumar-school/internal/config"
	"zumar-school/internal/models"
)

func init() {
	config.LoadENV()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.Class{})
	config.DB.AutoMigrate(&models.Student{})
}
