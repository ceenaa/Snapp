package controllers

import (
	"Project/initializers"
	"Project/models"
	"Project/validators"
	"encoding/json"
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
		"status":  "Failed",
		"message": err.Error(),
	})
}

func AddRuleToRedis(rule models.Rule) {
	rl, _ := json.Marshal(rule)
	initializers.RDB.HSet(initializers.Ctx, "rules", rule.ID, string(rl))
}

func AddRouteToRedis(rule models.Rule) {
	for _, i2 := range rule.Routes {
		t := i2.Origin + "-" + i2.Destination
		initializers.RDB.LPush(initializers.Ctx, t, i2.RuleID)
	}
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
		AddRuleToRedis(rl)
		AddRouteToRedis(rl)

	}

	createRuleResponse(c)
}
