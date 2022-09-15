package convert

import (
	"Project/models"
)

func RuleConvert(rule models.Rule) models.RawRule {
	var rawRule models.RawRule
	var rawRoutes []models.RawRoute
	for _, i2 := range rule.Routes {
		rawRoutes = append(rawRoutes, models.RawRoute{Origin: i2.Origin, Destination: i2.Destination})
	}
	rawRule.Airlines = rule.Airlines
	rawRule.Agencies = rule.Agencies
	rawRule.Suppliers = rule.Suppliers
	rawRule.AmountType = rule.AmountType
	rawRule.AmountValue = rule.AmountValue
	return rawRule
}
