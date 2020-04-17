package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)

	var sb strings.Builder
	b := prev[0]
	cnt := 1
	for i := 1; i < len(prev); i++ {
		if prev[i] == b {
			cnt++
		} else {
			sb.WriteString(strconv.Itoa(cnt))
			sb.WriteByte(b)
			cnt = 1
			b = prev[i]
		}
	}

	sb.WriteString(strconv.Itoa(cnt))
	sb.WriteByte(b)
	return sb.String()
}

func main() {
	fmt.Println(countAndSay(5))
}
