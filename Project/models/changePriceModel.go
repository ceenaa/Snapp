package models

type ChangePrice struct {
	RuleId       uint
	Origin       string
	Destination  string
	Airline      string
	Agency       string
	Supplier     string
	BasePrice    float64
	Markup       float64
	PayablePrice float64
}
