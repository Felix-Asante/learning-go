package conversions

import "strconv"

func StringToFloat64(value string) float64 {
	amount, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return amount
}
