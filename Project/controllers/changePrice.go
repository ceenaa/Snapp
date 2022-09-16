package controllers

import (
	"Project/coding"
	"Project/initializers"
	"Project/models"
	"Project/validators"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func IdsToCheck(origin string, destination string) []string {

	route1 := origin + "-" + destination
	route2 := origin + "-"
	route3 := "-" + destination
	route4 := "-"

	var ids []string

	ids1, err := initializers.RDB.LRange(initializers.Ctx, route1, 0, -1).Result()
	ids2, err := initializers.RDB.LRange(initializers.Ctx, route2, 0, -1).Result()
	ids3, err := initializers.RDB.LRange(initializers.Ctx, route3, 0, -1).Result()
	ids4, err := initializers.RDB.LRange(initializers.Ctx, route4, 0, -1).Result()

	if err != nil {
		log.Println("failed to get all rule ids")
		return nil
	}

	initializers.RDB.SAdd(initializers.Ctx, "ids", ids1)
	initializers.RDB.SAdd(initializers.Ctx, "ids", ids2)
	initializers.RDB.SAdd(initializers.Ctx, "ids", ids3)
	initializers.RDB.SAdd(initializers.Ctx, "ids", ids4)

	ids, err = initializers.RDB.SMembers(initializers.Ctx, "ids").Result()
	if err != nil {
		log.Println("failed to get all rule ids")
		return nil
	}

	initializers.RDB.Del(initializers.Ctx, "ids")

	return ids
}

func containAgency(checkID string, agency string) bool {
	containAgency := initializers.RDB.SIsMember(initializers.Ctx, "agencies"+checkID, agency).Val()
	if !containAgency {
		if initializers.RDB.SCard(initializers.Ctx, "agencies"+checkID).Val() == 0 {
			containAgency = true
		}
	}
	return containAgency
}

func containSupplier(checkID string, supplier string) bool {
	containSupplier := initializers.RDB.SIsMember(initializers.Ctx, "suppliers"+checkID, supplier).Val()
	if !containSupplier {
		if initializers.RDB.SCard(initializers.Ctx, "suppliers"+checkID).Val() == 0 {
			containSupplier = true
		}
	}
	return containSupplier
}

func containAirline(checkID string, airline string) bool {
	containAirline := initializers.RDB.SIsMember(initializers.Ctx, "airlines"+checkID, airline).Val()
	if !containAirline {
		if initializers.RDB.SCard(initializers.Ctx, "airlines"+checkID).Val() == 0 {
			containAirline = true
		}
	}
	return containAirline
}

func ChangePrices(c *gin.Context) {
	// Get data off req body

	var changePrices []models.ChangePrice
	err := c.Bind(&changePrices)
	if err != nil {
		log.Println("failed to bind data")
		return
	}

	for _, changePrice := range changePrices {
		err = validators.ValidateChangePrice(changePrice)
		if err != nil {
			createRuleError(c, err)
			continue
		}
		var price = 0.0
		var temp = 0.0
		var bs = changePrice.BasePrice
		checkIDs := IdsToCheck(changePrice.Origin, changePrice.Destination)

		for _, id := range checkIDs {
			ruleJson := initializers.RDB.HGet(initializers.Ctx, "rules", id).Val()
			var rule = coding.UnHash(ruleJson)

			id := strconv.Itoa(int(rule.ID))
			if containAgency(id, changePrice.Agency) && containSupplier(id, changePrice.Supplier) && containAirline(id, changePrice.Airline) {
				if rule.AmountType == "PERCENTAGE" {
					temp = bs + (bs * (float64(rule.AmountValue) / 100))
					if temp > price {
						price = temp
						changePrice.RuleId = rule.ID
					}
				} else {
					temp = bs + float64(rule.AmountValue)
					if temp > price {
						price = temp
						changePrice.RuleId = rule.ID
					}
				}
			}
		}

		if price != 0 {
			changePrice.Markup = price - changePrice.BasePrice
			changePrice.PayablePrice = price
		} else {
			changePrice.PayablePrice = changePrice.BasePrice
		}
		c.JSON(200, changePrice)

	}

}
