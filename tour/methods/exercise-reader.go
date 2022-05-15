//go:build OMIT

package main

import "github.com/hydrz/go-example/tour/methods/internal/reader"

type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
