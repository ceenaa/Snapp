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

func GetAll(db *gorm.DB) ([]Rule, error) {
	var rules []Rule
	err := db.Model(&Rule{}).Preload("Routes").Find(&rules).Error
	return rules, err
}

type Route struct {
	ID          uint
	RuleID      uint
	Origin      string
	Destination string
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
