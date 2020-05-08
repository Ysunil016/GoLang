package main

import (
	"fmt"
	"math"
)

func main() {
	ARR := [][]int{{0, 0}, {1, 1}}
	fmt.Println(checkStraightLine(ARR))
}

func checkStraightLine(coordinates [][]int) bool {
	if coordinates == nil || len(coordinates) == 0 || len(coordinates) == 1 {
		return false
	}

	X1 := coordinates[0][0]
	Y1 := coordinates[0][1]
	X2 := coordinates[1][0]
	Y2 := coordinates[1][1]

	primarySlope := math.Inf(1)
	if X2-X1 != 0 {
		primarySlope = float64(Y2-Y1) / float64(X2-X1)
	}

	for i := 0; i < len(coordinates)-1; i++ {
		X1 = coordinates[i][0]
		Y1 = coordinates[i][1]
		X2 = coordinates[i+1][0]
		Y2 = coordinates[i+1][1]

		XDiff := X2 - X1
		YDiff := Y2 - Y1

		if XDiff == 0 {
			if primarySlope != math.Inf(1) {
				return false
			}
		} else {
			slope := float64(YDiff) / float64(XDiff)
			if slope != primarySlope {
				return false
			}
		}

	}

	return true
}
