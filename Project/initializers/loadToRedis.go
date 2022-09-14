package initializers

import (
	"Project/models"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func LoadRules() {
	var rules []models.Rule
	DB.Find(&rules)
	for _, rule := range rules {
		var hRule models.HashRule
		var hRoutes []models.HashRoute
		for _, i2 := range rule.Routes {
			hRoutes = append(hRoutes, models.HashRoute{Origin: i2.Origin, Destination: i2.Destination})
		}
		hRule.Airlines = rule.Airlines
		hRule.Agencies = rule.Agencies
		hRule.Suppliers = rule.Suppliers
		RDB.SAdd(Ctx, "HashRules", Hash(hRule))
		rl, _ := json.Marshal(rule)
		RDB.HSet(Ctx, "rules", rule.ID, rl)
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

func Hash(rule models.HashRule) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(rule)
	return b.Bytes()
}
