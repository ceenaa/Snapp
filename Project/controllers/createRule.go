package controllers

import (
	"Project/initializers"
	"Project/models"
	"Project/validators"
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
	c.JSON(200, gin.H{
		"status":  "Failed",
		"message": err.Error(),
	})
}

func RulesCreate(c *gin.Context) {

	var rls []models.Rule
	err := c.Bind(&rls)

	if err != nil {
		log.Fatal("failed to bind data")
		return
	}

	for _, rl := range rls {

		err = validators.CheckDuplicateRule(rl)
		if err != nil {
			createRuleError(c, err)
			continue
		}

		err = validators.ValidateRule(rl)
		if err != nil {
			createRuleError(c, err)
			continue
		}
		if rl.Routes == nil {
			newRoute := models.Route{
				Origin:      "",
				Destination: "",
			}
			rl.Routes = append(rl.Routes, newRoute)
		}

		result := initializers.DB.Create(&rl)

		if result.Error != nil {
			createRuleError(c, result.Error)
			return
		}
		initializers.LoadRule(rl)
		initializers.LoadRoute(rl)

		createRuleResponse(c)

	}

}
