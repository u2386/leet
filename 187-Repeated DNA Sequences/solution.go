package main

func findRepeatedDnaSequences(s string) []string {
	ret := make([]string, 0)
	counter := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		if v, ok := counter[s[i:i+10]]; ok && v == 1 {
			ret = append(ret, s[i:i+10])
		}
		counter[s[i:i+10]]++
	}
	return ret
}

func main() {

}
