package main

func main() {

}

func findDuplicates(nums []int) []int {
	arr := make([]int, len(nums)+1)
	for _, V := range nums {
		arr[V]++
	}

	counter := 0
	res := make([]int, (len(nums) + 1))
	for i, V := range arr {
		if V > 1 {
			res[counter] = i
			counter++
		}
	}
	return res[0:counter]
}
