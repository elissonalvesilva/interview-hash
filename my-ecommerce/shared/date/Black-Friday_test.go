package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var blackFridayDate = time.Date(2021, 11, 29, 00, 00, 00, 00, time.UTC)

func TestIsBlackFriday(t *testing.T) {
	t.Run("Should return false if date is not equals to black friday date", func(t *testing.T) {
		t.Parallel()
		expectedResponse := false
		fakeDate := time.Date(2021, 10, 29, 00, 00, 00, 00, time.UTC)
		isBlackFridayDate := IsBlackFriday(fakeDate, blackFridayDate)
		assert.Equal(t, expectedResponse, isBlackFridayDate)
	})

	t.Run("Should return true if date is not equals to black friday date", func(t *testing.T) {
		t.Parallel()
		expectedResponse := true
		fakeDate := time.Date(2021, 11, 29, 00, 00, 00, 00, time.UTC)

		isBlackFridayDate := IsBlackFriday(fakeDate, blackFridayDate)
		assert.Equal(t, expectedResponse, isBlackFridayDate)
	})
}
