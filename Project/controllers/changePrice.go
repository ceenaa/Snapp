package controllers

import (
	"Project/initializers"
	"Project/models"
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

		res, er := models.GetWithRoute(initializers.DB, cp.Origin, cp.Destination)
		if er != nil {
			log.Fatal("failed to get all rules")
			return
		}

		for _, i2 := range res {
			if contains(i2.Airlines, cp.Airline) && contains(i2.Agencies, cp.Agency) && contains(i2.Suppliers, cp.Supplier) {
				if i2.AmountType == "percentage" {
					temp = bs + (bs * (i2.AmountValue / 100))
					if temp > price {
						price = temp
						cp.RuleId = i2.ID
					}
				} else {
					temp = bs + i2.AmountValue
					if temp > price {
						price = temp
						cp.RuleId = i2.ID
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
