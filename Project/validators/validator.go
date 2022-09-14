package validators

import (
	"Project/initializers"
	"Project/models"
	"errors"
	"github.com/lib/pq"
)

func CheckCity(city string) error {
	if city == "" {
		return nil
	}
	if initializers.RDB.SIsMember(initializers.Ctx, "cities", city).Val() {
		return nil
	}
	return errors.New("invalid city " + city)
}

func CheckSupplier(supplier string) error {
	if supplier == "" {
		return nil
	}
	if initializers.RDB.SIsMember(initializers.Ctx, "suppliers", supplier).Val() {
		return nil
	}
	return errors.New("invalid city " + supplier)
}

func CheckAirline(airline string) error {
	if airline == "" {
		return nil
	}
	if initializers.RDB.SIsMember(initializers.Ctx, "airlines", airline).Val() {
		return nil
	}
	return errors.New("invalid airline " + airline)
}

func CheckAgency(agency string) error {
	if agency == "" {
		return nil
	}
	if initializers.RDB.SIsMember(initializers.Ctx, "agencies", agency).Val() {
		return nil
	}
	return errors.New("invalid agency" + agency)
}

func CheckRoutes(routes []models.Route) error {
	for _, i := range routes {
		err1 := CheckCity(i.Origin)
		err2 := CheckCity(i.Destination)
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func CheckDuplicateRule(rule models.Rule) error {
	var hRule models.HashRule
	var hRoutes []models.HashRoute
	for _, i2 := range rule.Routes {
		hRoutes = append(hRoutes, models.HashRoute{Origin: i2.Origin, Destination: i2.Destination})
	}
	hRule.Airlines = rule.Airlines
	hRule.Agencies = rule.Agencies
	hRule.Suppliers = rule.Suppliers

	t := initializers.RDB.SIsMember(initializers.Ctx, "HashRules", initializers.Hash(hRule)).Val()
	if t == true {
		return errors.New("duplicate rule")
	}
	return nil
}

func CheckAirlines(airlines pq.StringArray) error {
	for _, i2 := range airlines {
		err := CheckAirline(i2)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckAgencies(agencies pq.StringArray) error {
	for _, i2 := range agencies {
		err := CheckAgency(i2)
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckSuppliers(supplier pq.StringArray) error {
	for _, i2 := range supplier {
		err := CheckSupplier(i2)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateRule(rule models.Rule) error {
	supplierErr := CheckSuppliers(rule.Suppliers)
	if supplierErr != nil {
		return supplierErr
	}
	agencyErr := CheckAgencies(rule.Agencies)
	if agencyErr != nil {
		return agencyErr
	}
	routesErr := CheckRoutes(rule.Routes)
	if routesErr != nil {
		return routesErr
	}
	airlinesErr := CheckAirlines(rule.Airlines)
	if airlinesErr != nil {
		return airlinesErr
	}
	return nil
}

func ValidateChangePrice(cp models.ChangePrice) error {
	supplierErr := CheckSupplier(cp.Supplier)
	if supplierErr != nil {
		return supplierErr
	}
	agencyErr := CheckAgency(cp.Agency)
	if agencyErr != nil {
		return agencyErr
	}
	originErr := CheckCity(cp.Origin)
	if originErr != nil {
		return originErr
	}

	destErr := CheckCity(cp.Destination)
	if destErr != nil {
		return destErr
	}

	airlineErr := CheckAirline(cp.Airline)
	if airlineErr != nil {
		return airlineErr
	}

	return nil

}
