package main

import "sort"

func main() {

}

func threeSum(num []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(num)
	for i := 0; i < len(num)-2; i++ {
		if i == 0 || i > 0 && num[i] != num[i-1] {
			lo := i + 1
			hi := len(num) - 1
			sum := 0 - num[i]
			for lo < hi {
				if num[lo]+num[hi] == sum {
					res = append(res, []int{num[i], num[lo], num[hi]})
					for lo < hi && num[lo] == num[lo+1] {
						lo++
					}
					for lo < hi && num[hi] == num[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if num[lo]+num[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return res
}
