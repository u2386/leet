package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	i := 0
	carry := 0
	var c byte
	buf := []byte{}
	for i < len(a) && i < len(b) {
		x, y := a[len(a)-1-i], b[len(b)-1-i]
		if x == '0' && y == '0' {
			c = '0'
			if carry == 1 {
				c = '1'
				carry = 0
			}
		} else if x == '1' && y == '1' {
			c = '0'
			if carry == 1 {
				c = '1'
			}
			carry = 1
		} else {
			c = '1'
			if carry == 1 {
				c = '0'
				carry = 1
			}
		}
		i++

		buf = append(buf, c)
	}

	for i < len(a) {
		c := a[len(a)-1-i]
		if c == '0' {
			if carry == 1 {
				c = '1'
				carry = 0
			}
		} else {
			if carry == 1 {
				c = '0'
			}
		}
		i++

		buf = append(buf, c)
	}
	for i < len(b) {
		c := b[len(b)-1-i]
		if c == '0' {
			if carry == 1 {
				c = '1'
				carry = 0
			}
		} else {
			if carry == 1 {
				c = '0'
			}
		}
		i++

		buf = append(buf, c)
	}
	if carry == 1 {
		buf = append(buf, '1')
	}

	for i := 0; i < len(buf)>>1; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}

	var sb strings.Builder
	for _, c := range buf {
		sb.WriteByte(c)
	}
	return sb.String()
}

func main() {
	a := "1"
	b := "111"
	fmt.Println(addBinary(a, b))
}
