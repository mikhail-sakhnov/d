package main

import "github.com/soider/d"

func main() {
	a := 123
	b := "hello world"
	c := 3.1415926
	d_ := func(n int) bool { return n > 0 }(1)
	e := []int{1, 2, 3}
	f := []byte("goodbye world")
	g := e[1:]

	d.D(a, b, c, d_, e, f, g)
}
