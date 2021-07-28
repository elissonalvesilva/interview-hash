package currency

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/currency"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncateNaive(t *testing.T) {
	t.Run("Should truncate a number with 2 places after dot", func(t *testing.T) {
		number := currency.TruncateNaive(152.00451354613)
		assert.Equal(t, 152.00, number)
	})

	t.Run("Should truncate a number with 2 places without dot", func(t *testing.T) {
		number := currency.TruncateNaive(152)
		assert.Equal(t, 152.00, number)
	})
}
