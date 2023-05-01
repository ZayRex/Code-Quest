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
	e := 60
	f, g, h := false, 23.9, "random"
	fmt.Println(a, b, c, d) //println adds spaces as well!
	fmt.Println(e, f, g, h)

	/* User specified types */
	const a int32 = 12        // 32-bit integer
	const b float32 = 20.5    // 32-bit float
	var c complex128 = 1 + 4i // 128-bit complex number
	var d uint16 = 14         // 16-bit unsigned integer

	/* Default types */
	n := 42              // int
	pi := 3.14           // float64
	x, y := true, false  // bool
	z := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
