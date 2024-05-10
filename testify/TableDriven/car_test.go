package tabledriven

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestData struct {
	car   Car
	tobuy bool
	msg   error
}

func TestCar(t *testing.T) {
	testcase := []TestData{
		{Car{Name: "Jaguar", Year: 16, Power: 450, Price: 4999}, true, nil},
		{Car{Name: "Nano", Year: 13, Power: 150, Price: 1000}, false, nil},
		{Car{Name: "BMW", Year: -16, Power: 450, Price: 5999}, false, errors.New("negetive value found")},
	}

	for _, tc := range testcase {
		buy, err := toBuy(tc.car)
		if tc.msg != nil {
			require.NotNil(t, err)
			require.Equal(t, tc.msg, err)
		} else {
			require.Nil(t, err)
			require.Equal(t, tc.tobuy, buy)
		}
	}
}
