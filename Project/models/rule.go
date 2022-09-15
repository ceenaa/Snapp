package models

import (
	"github.com/lib/pq"
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

type RawRule struct {
	Routes      []RawRoute
	Airlines    pq.StringArray `json:"airlines"`
	Agencies    pq.StringArray `json:"agencies"`
	Suppliers   pq.StringArray `json:"suppliers"`
	AmountType  string         `json:"amountType"`
	AmountValue int            `json:"amountValue"`
}

type RawRoute struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Airline struct {
	ID   uint
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
