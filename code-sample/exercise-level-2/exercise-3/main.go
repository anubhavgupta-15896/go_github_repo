package main

import "fmt"

const (
	a int     = 42
	b float64 = 23.32
	c string  = "james bond"
)

const (
	d = 42
	e = 23.32
	f = "james bond"
)

func main() {

	fmt.Println(a, b, c)
	fmt.Println(d, e, f)

}
