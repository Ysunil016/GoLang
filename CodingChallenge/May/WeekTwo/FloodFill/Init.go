package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(image)
	image = floodFill(image, 1, 1, 2)
	fmt.Println(image)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	traceK(image, sr, sc, newColor, len(image), len(image[0]), image[sr][sc])
	return image
}

func traceK(image [][]int, sr int, sc int, newColor int, r int, c int, key int) {

	if sr < 0 || sc < 0 || sr >= r || sc >= c {
		return
	}

	if image[sr][sc] == newColor {
		return
	}
	if image[sr][sc] == key {
		image[sr][sc] = newColor
	} else {
		return
	}
	fmt.Println(image)

	traceK(image, sr, sc+1, newColor, r, c, key)
	traceK(image, sr+1, sc, newColor, r, c, key)
	traceK(image, sr-1, sc, newColor, r, c, key)
	traceK(image, sr, sc-1, newColor, r, c, key)
}
