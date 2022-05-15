//go:build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	v := s[:4]
	printSlice(v) // output: len=4 cap=4 [5 7 11 13]

	v[0] = 100
	printSlice(v) // output: len=4 cap=4 [100 7 11 13]
	printSlice(s) // output: len=2 cap=4 [100 7]

	v = append(v, 200)
	printSlice(v) // output: len=5 cap=6 [100 7 11 13 200]
	printSlice(s) // output: len=2 cap=4 [100 7]

	v[0] = 1e9
	printSlice(v) // output: len=6 cap=6 [1000000000 7 11 13 200]
	printSlice(s) // output: len=2 cap=4 [100 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
