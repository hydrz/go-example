//go:build OMIT

package main

import (
	"math"

	"github.com/hydrz/go-example/tour/moretypes/internal/pic"
)

type Fun func(x, y int) int

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dy)
	for i := range m {
		m[i] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			m[y][x] = uint8(fun2(x, y))
		}
	}
	return m
}

func fun1(x, y int) int {
	return (x + y) / 2
}

func fun2(x, y int) int {
	return x * y
}

func fun3(x, y int) int {
	return x ^ y
}

func fun4(x, y int) int {
	return int(float64(x) * math.Log(float64(y)))
}

func fun5(x, y int) int {
	return x % (y + 1)
}

func main() {
	pic.Show(Pic)
}
