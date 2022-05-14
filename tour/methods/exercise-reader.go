//go:build OMIT

package main

import "github.com/hydrz/go-example/methods/internal/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
