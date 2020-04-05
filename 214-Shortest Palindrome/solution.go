package main

import (
	"fmt"
	"strings"
)

func shortestPalindrome(s string) string {
	r := []byte(s)
	for i := 0; i < len(r)>>1; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	rs := string(r)
	
	var sb strings.Builder
	var i int
	for i = 0; i < len(s); i++ {
		if strings.Compare(s[:len(s)-i], rs[i:]) == 0 {
			break
		}
		sb.WriteByte(rs[i])
	}
	sb.WriteString(s)
	return sb.String()
}

func main() {
	s := "aacecaaa"
	fmt.Println(shortestPalindrome(s))
}
