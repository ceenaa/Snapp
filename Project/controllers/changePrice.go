package controllers

import (
	"Project/coding"
	"Project/initializers"
	"Project/models"
	"Project/validators"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
)

func contains(s pq.StringArray, e string) bool {
	if s == nil || len(s) == 0 || e == "" {
		return true
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IdsToCheck(route1 string, route2 string, route3 string, route4 string) []string {
	// appending must be better
	var ids []string

	m := make(map[string]bool)

	ids1, err := initializers.RDB.LRange(initializers.Ctx, route1, 0, -1).Result()
	for _, i := range ids1 {
		if !m[i] {
			ids = append(ids, i)
			m[i] = true
		}
	}
	ids2, err := initializers.RDB.LRange(initializers.Ctx, route2, 0, -1).Result()
	for _, i := range ids2 {
		if !m[i] {
			ids = append(ids, i)
			m[i] = true
		}
	}
	ids3, err := initializers.RDB.LRange(initializers.Ctx, route3, 0, -1).Result()
	for _, i := range ids3 {
		if !m[i] {
			ids = append(ids, i)
			m[i] = true
		}
	}
	ids4, err := initializers.RDB.LRange(initializers.Ctx, route4, 0, -1).Result()
	for _, i := range ids4 {
		if !m[i] {
			ids = append(ids, i)
			m[i] = true
		}
	}

	if err != nil {
		log.Fatal("failed to get all rule ids")
		return nil
	}

	return ids
}

func ChangePrice(c *gin.Context) {
	// Get data off req body

	var cps []models.ChangePrice
	err := c.Bind(&cps)
	if err != nil {
		log.Fatal("failed to bind data")
		return
	}

	for _, cp := range cps {
		err = validators.ValidateChangePrice(cp)
		if err != nil {
			createRuleError(c, err)
			continue
		}

		var price = 0.0
		var temp = 0.0
		var bs = cp.BasePrice

		route1 := cp.Origin + "-" + cp.Destination
		route2 := cp.Origin + "-"
		route3 := "-" + cp.Destination
		route4 := "-"

		// appending must be better
		ids := IdsToCheck(route1, route2, route3, route4)

		for _, i2 := range ids {
			ruleJson := initializers.RDB.HGet(initializers.Ctx, "rules", i2).Val()
			var rule = coding.UnHash(ruleJson)
			if err != nil {
				log.Fatal("failed to unmarshal rule")
				return
			}

			if contains(rule.Airlines, cp.Airline) && contains(rule.Agencies, cp.Agency) && contains(rule.Suppliers, cp.Supplier) {
				if rule.AmountType == "PERCENTAGE" {
					temp = bs + (bs * (float64(rule.AmountValue) / 100))
					if temp > price {
						price = temp
						cp.RuleId = rule.ID
					}
				} else {
					temp = bs + float64(rule.AmountValue)
					if temp > price {
						price = temp
						cp.RuleId = rule.ID
					}
				}
			}
		}
		if price != 0 {
			cp.Markup = price - cp.BasePrice
			cp.PayablePrice = price
		} else {
			cp.PayablePrice = cp.BasePrice
		}
		c.JSON(200, cp)

	}

}
