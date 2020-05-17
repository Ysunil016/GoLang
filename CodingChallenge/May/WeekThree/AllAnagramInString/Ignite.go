package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, p))
}
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	result := []int{}
	primaryArray := make([]int, 26)
	for i := 0; i < len(p); i++ {
		primaryArray[p[i]-'a']++
	}
	secandoryArray := make([]int, 26)
	for i := range s {
		secandoryArray[s[i]-'a']++
		if i >= len(p) {
			secandoryArray[s[i-len(p)]-'a']--
		}
		if compareArray(secandoryArray, primaryArray) == true {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

func compareArray(p, s []int) bool {
	for i := 0; i < 26; i++ {
		if p[i] != s[i] {
			return false
		}
	}
	return true
}
