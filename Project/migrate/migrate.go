package main

import (
	"Project/initializers"
	"Project/models"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Rule{}, &models.Route{}, &models.Airline{}, &models.Agency{}, &models.Supplier{}, &models.City{})
}
