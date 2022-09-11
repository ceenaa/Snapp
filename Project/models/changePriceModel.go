package models

type ChangePrice struct {
	RuleId       uint
	Origin       string
	Destination  string
	Airline      string
	Agency       string
	Supplier     string
	BasePrice    int
	Markup       int
	PayablePrice int
}
