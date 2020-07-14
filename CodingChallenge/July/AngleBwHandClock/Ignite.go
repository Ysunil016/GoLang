package main

import "fmt"

func main() {
	fmt.Println(angleClock(12, 30))
}

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	var angleDifference float64 = abs(float64(hour)*30 + float64(minutes)*(0.5-6))
	return min(angleDifference, float64(360)-angleDifference)
}
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}
