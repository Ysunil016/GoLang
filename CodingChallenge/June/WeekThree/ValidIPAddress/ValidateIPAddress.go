package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(validIPAddress("172.16.254.1"))
	// fmt.Println(validIPAddress("256.256.256.256"))
	// fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	// fmt.Println(validIPAddress("2001:0db8:85a3:00000:0:8A2E:0370:7334"))
	fmt.Println(validIPAddress("1081:db8:85a3:01:-0:8A2E:0370:7334"))
}

func validIPAddress(IP string) string {
	bufferIPv4 := strings.Split(IP, ".")
	bufferIPv6 := strings.Split(IP, ":")
	if len(bufferIPv4) == 4 {
		O1, err := strconv.ParseInt(bufferIPv4[0], 10, 64)
		if err != nil || O1 >= 256 || O1 < 0 {
			return "Neither"
		}
		O2, err := strconv.ParseInt(bufferIPv4[1], 10, 64)
		if err != nil || O2 >= 256 || O2 < 0 {
			return "Neither"
		}
		O3, err := strconv.ParseInt(bufferIPv4[2], 10, 64)
		if err != nil || O3 >= 256 || O3 < 0 {
			return "Neither"
		}
		O4, err := strconv.ParseInt(bufferIPv4[3], 10, 64)
		if err != nil || O4 >= 256 || O4 < 0 {
			return "Neither"
		}

		S1 := strings.Split(bufferIPv4[0], "")
		if S1[0] == "0" && O1 > 0 {
			return "Neither"

		}
		S2 := strings.Split(bufferIPv4[1], "")
		if S2[0] == "0" && O2 > 0 {
			return "Neither"

		}
		S3 := strings.Split(bufferIPv4[2], "")
		if S3[0] == "0" && O3 > 0 {
			return "Neither"

		}
		S4 := strings.Split(bufferIPv4[3], "")
		if S4[0] == "0" && O4 > 0 {
			return "Neither"

		}
		if O1 < 10 {
			if len(bufferIPv4[0]) > 1 {
				return "Neither"
			}
		}
		if O2 < 10 {
			if len(bufferIPv4[1]) > 1 {
				return "Neither"
			}
		}
		if O3 < 10 {
			if len(bufferIPv4[2]) > 1 {
				return "Neither"
			}
		}
		if O4 < 10 {
			if len(bufferIPv4[3]) > 1 {
				return "Neither"
			}
		}

		return "IPv4"
	}
	if len(bufferIPv6) == 8 {
		O1, err := strconv.ParseInt(bufferIPv6[0], 16, 64)
		if err != nil || O1 >= 65535 || O1 < 0 || len(bufferIPv6[0]) > 4 || strings.Split(bufferIPv6[0], "")[0] == "-" {
			return "Neither"
		}
		O2, err := strconv.ParseInt(bufferIPv6[1], 16, 64)
		if err != nil || O2 >= 65535 || O2 < 0 || len(bufferIPv6[1]) > 4 || strings.Split(bufferIPv6[1], "")[0] == "-" {
			return "Neither"
		}
		O3, err := strconv.ParseInt(bufferIPv6[2], 16, 64)
		if err != nil || O3 >= 65535 || O3 < 0 || len(bufferIPv6[2]) > 4 || strings.Split(bufferIPv6[2], "")[0] == "-" {
			return "Neither"
		}
		O4, err := strconv.ParseInt(bufferIPv6[3], 16, 64)
		if err != nil || O4 >= 65535 || O4 < 0 || len(bufferIPv6[3]) > 4 || strings.Split(bufferIPv6[3], "")[0] == "-" {
			return "Neither"
		}
		O5, err := strconv.ParseInt(bufferIPv6[4], 16, 64)
		if err != nil || O5 >= 65535 || O5 < 0 || len(bufferIPv6[4]) > 4 || strings.Split(bufferIPv6[4], "")[0] == "-" {
			return "Neither"
		}
		O6, err := strconv.ParseInt(bufferIPv6[5], 16, 64)
		if err != nil || O6 >= 65535 || O6 < 0 || len(bufferIPv6[5]) > 4 || strings.Split(bufferIPv6[5], "")[0] == "-" {
			return "Neither"
		}
		O7, err := strconv.ParseInt(bufferIPv6[6], 16, 64)
		if err != nil || O7 >= 65535 || O7 < 0 || len(bufferIPv6[6]) > 4 || strings.Split(bufferIPv6[6], "")[0] == "-" {
			return "Neither"
		}
		O8, err := strconv.ParseInt(bufferIPv6[7], 16, 64)
		if err != nil || O8 >= 65535 || O8 < 0 || len(bufferIPv6[7]) > 4 || strings.Split(bufferIPv6[7], "")[0] == "-" {
			return "Neither"
		}
		return "IPv6"
	}
	return "Neither"
}
