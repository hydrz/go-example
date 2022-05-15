//go:build OMIT

package main

import (
	"strings"

	"github.com/hydrz/go-example/tour/moretypes/internal/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
