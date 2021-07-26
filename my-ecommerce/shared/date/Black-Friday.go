package date

import "time"

func IsBlackFriday(actualDate time.Time, blackFridayDate time.Time) bool {
	isBlackFriday := false

	if actualDate == blackFridayDate {
		isBlackFriday = true
		return isBlackFriday
	}

	return isBlackFriday
}
