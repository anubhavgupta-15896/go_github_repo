package main

import "fmt"

const (
	a         = 42
	b         = 42.32
	c         = "James Bond"
	d int     = 42
	e float64 = 42.32
	f string  = "James Bond"
)

const g = 42
const h int = 23

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
