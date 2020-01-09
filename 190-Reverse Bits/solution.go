package main

import "fmt"

var lookup = [16]uint32{0, 2147483648, 1073741824, 3221225472, 536870912, 2684354560, 1610612736, 3758096384, 268435456, 2415919104, 1342177280, 3489660928, 805306368, 2952790016, 1879048192, 4026531840}

func reverseBits(num uint32) uint32 {
	if num < 16 {
		return lookup[num]
	}

	mask := uint32(15)
	ret := uint32(0)
	shift := 0
	for num > 0 {
		ret |= (lookup[num&mask] >> shift)
		num >>= 4
		shift += 4
	}
	return ret
}

func main() {
	fmt.Println(reverseBits(4294967293))
}
