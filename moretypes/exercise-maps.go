//go:build OMIT

package main

import "github.com/hydrz/go-example/moretypes/internal/wc"

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
