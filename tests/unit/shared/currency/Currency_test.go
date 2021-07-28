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

func TestParseToCents(t *testing.T) {
	t.Run("Should return 0 if value param is equals to 0", func(t *testing.T) {
		t.Parallel()

		valueParam := 0.0

		response := currency.ParseToCents(valueParam)
		assert.Equal(t, int64(0), response)
	})

	t.Run("Should return a response in cents if has number after dot", func(t *testing.T) {
		t.Parallel()

		valueParam := 548.15

		response := currency.ParseToCents(valueParam)
		assert.Equal(t, int64(54815), response)
	})

	t.Run("Should return a response in cents if no has number after dot", func(t *testing.T) {
		t.Parallel()

		valueParam := 548.00

		response := currency.ParseToCents(valueParam)
		assert.Equal(t, int64(54800), response)
	})
}
