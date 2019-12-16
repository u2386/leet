package main

import (
	"fmt"
	"strings"
)

type pair struct {
	word  string
	level int
}

func generatePattern(word string) (patterns []string) {
	patterns = []string{}
	for i := 0; i < len(word); i++ {
		sb := strings.Builder{}
		sb.WriteString(word[:i])
		sb.WriteString("*")
		sb.WriteString(word[i+1:])
		patterns = append(patterns, sb.String())
	}
	return
}

func updatePatterns(patterns map[string][]string, word string) {
	for _, pat := range generatePattern(word) {
		if _, ok := patterns[pat]; !ok {
			patterns[pat] = make([]string, 1)
		}
		patterns[pat] = append(patterns[pat], word)
	}
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	patterns := make(map[string][]string)

	exist := false
	for _, word := range wordList {
		if word == endWord {
			exist = true
		}
		updatePatterns(patterns, word)
	}
	if !exist {
		return 0
	}

	updatePatterns(patterns, beginWord)

	q := append(make([]pair, len(wordList)), pair{beginWord, 1})
	seen := make(map[string]bool)
	var node pair
	for len(q) > 0 {
		node, q = q[0], q[1:]
		word := node.word
		level := node.level

		for _, pat := range generatePattern(word) {
			for _, combo := range patterns[pat] {
				if combo == endWord {
					return level + 1
				}

				if _, ok := seen[combo]; ok {
					continue
				}
				seen[combo] = true
				q = append(q, pair{combo, level + 1})
			}
		}
	}
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
