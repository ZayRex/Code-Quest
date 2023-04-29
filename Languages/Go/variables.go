package main

import (
	"fmt"
)

var a int

var (
	b string
	c bool
	d float32
)

func main() {
	a = 42
	b, c, d = "Hello", true, 3.14
	fmt.Println(a, b, c, d)//println adds spaces as well!
}
