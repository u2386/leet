package main

import (
	"fmt"
)

type node struct {
	word string
	next map[rune]*node
}

func search(node *node, s string) bool {
	if len(node.word) > 0 && len(s) == 0 {
		return true
	}

	if (node.next == nil && len(s) > 0) || (node.next != nil && len(s) == 0) {
		return false
	}

	if _, ok := node.next[rune(s[0])]; !ok {
		return false
	}
	return search(node.next[rune(s[0])], s[1:])
}

func divide(s string, n *node) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	for i < len(s) {
		if search(n, s[:i+1]) && divide(s[i+1:], n) {
			return true
		}
		i++
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	root := &node{}
	for _, word := range wordDict {
		n := root
		for _, ch := range word {
			if n.next == nil {
				n.next = make(map[rune]*node)
			}
			v, ok := n.next[ch]
			if !ok {
				v = &node{}
				n.next[ch] = v
			}
			n = v
		}
		n.word = word
	}

	return divide(s, root)
}

func main() {
	s := "abcd"
	wd := []string{"a", "abc", "b", "cd"}
	fmt.Println(wordBreak(s, wd))
}
