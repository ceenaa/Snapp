package main

import (
	"Project/initializers"
	"Project/models"
	"log"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Rule{}, &models.Route{}, &models.Airline{}, &models.Agency{}, &models.Supplier{}, &models.City{})
	if err != nil {
		log.Println(err)
		return
	}
}
