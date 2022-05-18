//go:build OMIT

package main

import "fmt"

type A struct {
	Foo string
}

func (a A) Say() {
	fmt.Printf("A Say: %s\n", a.Foo)
}

func (a *A) Say2() {
	fmt.Printf("A Say2: %s\n", a.Foo)
}

type B struct {
	A
}

func (b B) Say() {
	fmt.Printf("B Say: %s\n", b.Foo)
}

type C struct {
	*A
}

type D struct {
	A A
}

type E struct {
	A *A
}

func main() {
	a := A{Foo: "foo"}
	b := B{A: a}
	c := C{A: &a}
	d := D{A: a}
	e := E{A: &a}

	a.Say()

	b.Say()
	b.Say2()

	c.Say()
	c.Say2()

	d.A.Say()
	d.A.Say2()

	e.A.Say()
	e.A.Say2()

}
