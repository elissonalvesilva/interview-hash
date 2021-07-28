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

func ParseToCents(value float64) int64 {
	if value == 0.0 {
		return int64(0)
	}
	cents := 100

	return int64(value * float64(cents))
}
