package main

func main() {

}
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	result := 0
	for true {
		result += num % 10
		num = num / 10
		if num == 0 {
			if result < 10 {
				return result
			}
			num, result = result, 0
		}
	}
	return result
}
