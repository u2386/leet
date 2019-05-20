package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	s := 0
	for _, n := range(A) {
		if isEven(n) {
			s += n
		}
	}

	var val, idx int
	answer := []int{}
	for _, item := range(queries) {
		val, idx = item[0], item[1]

		if isEven(val) {
			if isEven(A[idx]) {
				s += val
			}
		} else {
			if !isEven(A[idx]) {
				s += val + A[idx]
			} else {
				s -= A[idx]
			}
		}
		A[idx] += val
		answer = append(answer, s)
	}
	return answer
}


func isEven(n int) bool {
	return n % 2 == 0
}
