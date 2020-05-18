package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidaboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	primary := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		primary[s1[i]-'a']++
	}
	secondary := make([]int, 26)
	for i := 0; i < len(s2); i++ {
		secondary[s2[i]-'a']++
		if i >= len(s1) {
			secondary[s2[i-len(s1)]-'a']--
		}
		if compare(primary, secondary) {
			return true
		}
	}
	return false
}
func compare(x, y []int) bool {
	for i := 0; i < 26; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
