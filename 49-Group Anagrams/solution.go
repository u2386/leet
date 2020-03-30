package main

import (
	"fmt"
	"sort"
)

type SortBy []byte

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		w := append([]byte{}, str...)
		sort.Sort(SortBy(w))
		k := string(w)
		m[k] = append(m[k], str)
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(arr))
}
