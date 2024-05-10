package tabledriven

import "errors"

type Car struct {
	Name  string
	Year  int
	Power int
	Price int
}

func toBuy(c Car) (bool, error) {
	if c.Name == "" {
		return false, errors.New("name invalid")
	}

	if c.Year < 15 || c.Power < 300 || c.Price > 5000 {
		return false, errors.New("not worth buying")
	} else {
		return true, nil
	}
}
