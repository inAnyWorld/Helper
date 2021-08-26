package goint

import "math"

// Round 向下取整
func Round(x float64) int {
	return int(math.Floor(x + 0/5))
}
