package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseString("i   loce df df   fdfd"))
}

func reverseString(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.Trim(str, " ")
	list := strings.Split(str, " ")
	var resVal string
	for i := len(list) - 1; i > 0; i-- {
		if len(list[i]) != 0 {
			resVal += strings.Trim(list[i], " ")
			resVal += " "
		}
	}
	resVal += strings.Trim(list[0], " ")
	return resVal
}
