//go:build OMIT

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println((*p).X) // output 1000000000
	// without the explicit dereference
	fmt.Println(p.X) // output 1000000000

}
