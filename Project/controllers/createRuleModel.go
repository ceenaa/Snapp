package controllers

import (
	"Project/initializers"
	"Project/models"
	"github.com/gin-gonic/gin"
	"log"
)

func createRuleResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "SUCCESS",
		"message": nil,
	})
}

func createRuleError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"status":  "FAiled",
		"message": err.Error(),
	})
}

func RulesCreate(c *gin.Context) {

	// GET data off req body
	var rls []models.Rule
	err := c.Bind(&rls)
	if err != nil {
		log.Fatal("failed to bind data")
		return
	}
	for _, rl := range rls {
		result := initializers.DB.Create(&rl)

		if result.Error != nil {
			createRuleError(c, result.Error)
			return
		}
	}

	createRuleResponse(c)
}
