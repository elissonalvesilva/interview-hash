package date

import (
	"time"
)

func IsBlackFriday(actualDate time.Time, blackFridayDate time.Time) bool {
	isBlackFriday := false

	if 	actualDate.Month() == blackFridayDate.Month() &&
		actualDate.Day() == blackFridayDate.Day(){
		isBlackFriday = true
		return isBlackFriday
	}

	return isBlackFriday
}
