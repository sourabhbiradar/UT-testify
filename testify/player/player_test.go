package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsMOTM(t *testing.T) {
	t.Run("error testing", func(t *testing.T) {
		s := Stats{
			Name: "ABD",
			Runs: -107,
			Wkts: 0,
		}
		_, err := IsMOTM(s)
		require.NotNil(t, err) // there SHOULD be error here
	})

	t.Run("isMotm", func(t *testing.T) {
		s := Stats{
			Name: "ABD",
			Runs: 107,
			Wkts: 0,
		}
		isMotm, err := IsMOTM(s)
		require.Nil(t, err)
		assert.True(t, isMotm)

	})
}
