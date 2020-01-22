package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	if n < 1 {
		return ""
	}

	buf := []byte{}
	var x, rem int
	for n > 0 {
		x = n / 26
		rem = n - x*26
		if rem == 0 {
			rem = 26
			x--
		}
		buf = append(buf, byte(rem+64))
		n = x
	}

	i, j := 0, len(buf)-1
	for i < j {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}

	return string(buf)
}

func main() {
	fmt.Println(convertToTitle(701))
}
