//go:build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)

	p = &v1
	p.X = 1e9
	fmt.Println(v1) // output {1000000000 2}

	v4 := v2
	v4.X = 1e9
	fmt.Println(v2) // output {1 0}

}
