package minDeltaUtil

import (
	"math"
	"strconv"
)

func Finder(compareIndexOne, compareIndexTwo, returnIndex int) func([]string) string {
	var smallestDelta float64
	smallestDelta = 1000
	returnField := ""
	return func(input []string) string {
		val1, err := strconv.ParseFloat(input[compareIndexOne], 64)
		if err != nil {
			panic(err)
		}
		val2, err := strconv.ParseFloat(input[compareIndexOne], 64)
		if err != nil {
			panic(err)
		}
		delta := math.Abs((float64)(val1 - val2))

		if delta < smallestDelta {
			smallestDelta = delta
			returnField = input[returnIndex]
		}
		return returnField
	}
}
