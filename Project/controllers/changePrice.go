package controllers

import (
	"Project/initializers"
	"Project/models"
	"encoding/json"
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

func ChangePrice(c *gin.Context) {
	// Get data off req body

	var cps []models.ChangePrice
	err := c.Bind(&cps)
	if err != nil {
		log.Fatal("failed to bind data")
	}

	for _, cp := range cps {

		var price = 0
		var temp = 0
		var bs = cp.BasePrice

		route1 := cp.Origin + "-" + cp.Destination
		route2 := cp.Origin + "-"
		route3 := "-" + cp.Destination
		route4 := "-"

		ids1, er := initializers.RDB.LRange(initializers.Ctx, route1, 0, -1).Result()
		ids2, er := initializers.RDB.LRange(initializers.Ctx, route2, 0, -1).Result()
		ids := append(ids1, ids2...)
		ids3, er := initializers.RDB.LRange(initializers.Ctx, route3, 0, -1).Result()
		ids = append(ids, ids3...)
		ids4, er := initializers.RDB.LRange(initializers.Ctx, route4, 0, -1).Result()
		ids = append(ids, ids4...)

		// appending must be better
		
		if er != nil {
			log.Fatal("failed to get all rules")
			return
		}

		for _, i2 := range ids {
			ruleJson := initializers.RDB.HGet(initializers.Ctx, "rules", i2).Val()

			var rule models.Rule
			err := json.Unmarshal([]byte(ruleJson), &rule)
			if err != nil {
				log.Fatal("failed to unmarshal rule")

			}

			if contains(rule.Airlines, cp.Airline) && contains(rule.Agencies, cp.Agency) && contains(rule.Suppliers, cp.Supplier) {
				if rule.AmountType == "percentage" {
					temp = bs + (bs * (rule.AmountValue / 100))
					if temp > price {
						price = temp
						cp.RuleId = rule.ID
					}
				} else {
					temp = bs + rule.AmountValue
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
