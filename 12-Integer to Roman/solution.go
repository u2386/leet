package main

import (
	"fmt"
	"strings"
)

var m = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

var v = []string{
	"M",
	"CM",
	"D",
	"CD",
	"C",
	"XC",
	"L",
	"XL",
	"X",
	"IX",
	"V",
	"IV",
	"I",
}

func intToRoman(num int) string {
	var sb strings.Builder
	for num > 0 {
		for i := 0; i < len(m); i++ {
			if num >= m[i] {
				sb.WriteString(v[i])
				num -= m[i]
				break
			}
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(intToRoman(1994))
}
