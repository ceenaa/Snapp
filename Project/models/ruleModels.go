package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Rule struct {
	ID          uint           `json:"id"`
	Routes      []Route        `json:"routes"`
	Airlines    pq.StringArray `gorm:"type:varchar(50)[]" json:"airlines"`
	Agencies    pq.StringArray `gorm:"type:varchar(50)[]" json:"agencies"`
	Suppliers   pq.StringArray `gorm:"type:varchar(50)[]" json:"suppliers"`
	AmountType  string         `json:"amountType"`
	AmountValue int            `json:"amountValue"`
}

type Route struct {
	ID          uint   `json:"id"`
	RuleID      uint   `json:"ruleId"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

func GetWithRoute(db *gorm.DB, origin string, destination string) ([]Rule, error) {
	var routes []Route
	err := db.Model(&Route{}).Where("(Origin = ? AND Destination = ?) OR (Origin = ? AND Destination = ?) OR (Origin = ? AND Destination = ?)", origin, destination, origin, "", "", destination).Find(&routes).Error
	var ids []uint
	for _, i2 := range routes {
		ids = append(ids, i2.RuleID)
	}
	var rules []Rule
	if len(ids) > 0 {
		err = db.Find(&rules, ids).Error
	}

	return rules, err
}

type Airline struct {
	ID   uint
	Code string
	Name string
}

type Agency struct {
	ID   uint
	Name string
}

type Supplier struct {
	ID   uint
	Name string
}
type City struct {
	ID   uint
	Code string
}
