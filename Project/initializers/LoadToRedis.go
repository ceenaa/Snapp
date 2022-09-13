package initializers

import (
	"Project/models"
	"encoding/json"
	"fmt"
)

func LoadRules() {
	var rules []models.Rule
	DB.Find(&rules)
	for _, rule := range rules {
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
		fmt.Println(t)
	}
}
