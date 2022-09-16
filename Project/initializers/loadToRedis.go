package initializers

import (
	"Project/coding"
	"Project/convert"
	"Project/models"
	"strconv"
)

func LoadRule(rule models.Rule) {
	ruleHash := coding.Hash(rule)
	RDB.HSet(Ctx, "rules", rule.ID, ruleHash)

	rawRule := convert.RuleConvert(rule)
	RDB.SAdd(Ctx, "HashRules", coding.HashRaw(rawRule))

	id := strconv.Itoa(int(rule.ID))
	var airlines []string
	airlines = append(airlines, rule.Airlines...)
	var agencies []string
	agencies = append(agencies, rule.Agencies...)
	var suppliers []string
	suppliers = append(suppliers, rule.Suppliers...)

	RDB.SAdd(Ctx, "agencies"+id, agencies)
	RDB.SAdd(Ctx, "airlines"+id, airlines)
	RDB.SAdd(Ctx, "suppliers"+id, suppliers)
}

func LoadRoute(rule models.Rule) {
	for _, i2 := range rule.Routes {
		t := i2.Origin + "-" + i2.Destination
		RDB.LPush(Ctx, t, i2.RuleID)
	}
}

func LoadRules() {
	var rules []models.Rule
	DB.Find(&rules)
	for _, rule := range rules {
		LoadRule(rule)
	}
}

func LoadRoutes() {
	var routes []models.Route
	DB.Find(&routes)
	for _, route := range routes {
		t := route.Origin + "-" + route.Destination
		RDB.LPush(Ctx, t, route.RuleID)
	}
}

func LoadCities() {
	var cities []models.City
	DB.Find(&cities)
	for _, i2 := range cities {
		RDB.SAdd(Ctx, "cities", i2.Code)
	}
}

func LoadSuppliers() {
	var suppliers []models.Supplier
	DB.Find(&suppliers)
	for _, i2 := range suppliers {
		RDB.SAdd(Ctx, "suppliers", i2.Name)
	}
}

func LoadAgencies() {
	var agencies []models.Agency
	DB.Find(&agencies)
	for _, i2 := range agencies {
		RDB.SAdd(Ctx, "agencies", i2.Name)
	}
}

func LoadAirlines() {
	var airlines []models.Airline
	DB.Find(&airlines)
	for _, i2 := range airlines {
		RDB.SAdd(Ctx, "airlines", i2.Name)
	}
}
