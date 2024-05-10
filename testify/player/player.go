package player

import "errors"

type Stats struct {
	Name string
	Runs int
	Wkts int
}

func IsMOTM(s Stats) (bool, error) {
	if s.Name == "" {
		return false, errors.New("please enter valid name")
	}
	if s.Runs < 0 || s.Wkts < 0 {
		return false, errors.New("runs or wkts can be negetive")
	}
	if s.Runs > 50 || s.Wkts > 3 {
		return true, nil
	} else {
		return false, nil
	}
}
