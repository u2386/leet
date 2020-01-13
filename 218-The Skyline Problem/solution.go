package main

func getSkyline(buildings [][]int) [][]int {
	skyline := [][]int{}
	if len(buildings) == 0 {
		return skyline
	}

	q := [][]int{}
	q = append(q, buildings[0])
	for i := 1; i < len(buildings); i++ {
		b := buildings[i]
		
		for len(q) > 0 {
			if b[0] < q[0][0] {
				break
			}
			q = q[1:]
		}
	}

	return skyline
}

func main() {

}
