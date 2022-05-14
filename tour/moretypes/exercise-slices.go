//go:build OMIT

package main

import "github.com/hydrz/go-example/moretypes/internal/pic"

func Pic(dx, dy int) [][]uint8 {
	return [][]uint8{}
}

func main() {
	pic.Show(Pic)
}
