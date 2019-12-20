package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringPattern(t *testing.T) {
	contains := func(container map[string][]string, target string) bool {
		_, ok := container[target]
		return ok
	}

	str := "hello"
	graph := make(map[string][]string)
	generateGraph(str, graph)

	assert.Equal(t, 5, len(graph))
	assert.True(t, contains(graph, "*ello"))
	assert.True(t, contains(graph, "h*llo"))
	assert.True(t, contains(graph, "he*lo"))
	assert.True(t, contains(graph, "hel*o"))
	assert.True(t, contains(graph, "hell*"))
}
