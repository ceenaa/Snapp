package controllers

import (
	"Project/initializers"
	"Project/models"
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
		"status":  "FAiled",
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
		
		result := initializers.DB.Create(&rl)
		AddRuleToRedis(rl)
		AddRouteToRedis(rl)

		if result.Error != nil {
			createRuleError(c, result.Error)
			return
		}
	}

	createRuleResponse(c)
}
