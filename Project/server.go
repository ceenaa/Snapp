package main

import (
	"Project/controllers"
	"Project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.RedisConnect()
	initializers.LoadRules()
	initializers.LoadRoutes()
}
func main() {
	r := gin.Default()
	r.POST("/createRule", controllers.RulesCreate)
	r.POST("/changePrice", controllers.ChangePrice)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
