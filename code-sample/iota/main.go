package main

import "fmt"

/*onst (
	a = iota
	b = iota
	c = iota
)*/

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
