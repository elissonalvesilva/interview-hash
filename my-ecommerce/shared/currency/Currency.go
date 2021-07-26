package currency

import (
	"fmt"
	"strconv"
)

func TruncateNaive(value float64) float64 {
	if formattedNumber, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64); err == nil {
		return formattedNumber
	}

	return value
}
