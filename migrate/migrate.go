package main

import (
	"github.com/karsteneugene/absen-backend/models"
	"github.com/karsteneugene/absen-backend/settings"
)

func init() {
	settings.LoadEnvVariables()
	settings.ConnectToDB()
}

func main() {
	settings.DB.AutoMigrate(&models.User{})
	settings.DB.AutoMigrate(&models.Image{})
}
