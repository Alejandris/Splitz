package money

import "math"

const scale = 100

// Round converts a monetary value to two decimal places using commercial rounding.
func Round(value float64) float64 {
	return math.Round(value*scale) / scale
}
