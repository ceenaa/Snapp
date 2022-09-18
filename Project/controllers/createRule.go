package controllers

import (
	"Project/initializers"
	"Project/models"
	"Project/validators"
	"github.com/gin-gonic/gin"
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

	var rules []models.Rule
	err := c.Bind(&rules)

	if err != nil {
		createRuleError(c, err)
		return
	}

	for _, rule := range rules {

		err = validators.CheckDuplicateRule(rule)
		if err != nil {
			createRuleError(c, err)
			continue
		}

		err = validators.ValidateRule(rule)
		if err != nil {
			createRuleError(c, err)
			continue
		}
		if rule.Routes == nil {
			newRoute := models.Route{
				Origin:      "",
				Destination: "",
			}
			rule.Routes = append(rule.Routes, newRoute)
		}

		result := initializers.DB.Create(&rule)

		if result.Error != nil {
			createRuleError(c, result.Error)
			return
		}
		initializers.LoadRule(rule)
		initializers.LoadRoute(rule)

		createRuleResponse(c)

	}

}
