package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	w := strings.Fields(s)
	m := make(map[string]int)

	for _, e := range w {
		v, exists := m[e]
		if exists {
			m[e] = v + 1
		} else {
			m[e] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
