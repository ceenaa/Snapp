package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Rule struct {
	ID          uint
	Routes      []Route
	Airlines    pq.StringArray `gorm:"type:varchar(50)[]"`
	Agencies    pq.StringArray `gorm:"type:varchar(50)[]"`
	Suppliers   pq.StringArray `gorm:"type:varchar(50)[]"`
	AmountType  string
	AmountValue int
}

type Route struct {
	ID          uint
	RuleID      uint
	Origin      string
	Destination string
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
