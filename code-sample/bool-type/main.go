package main

import "fmt"

var flag bool

func main() {

	fmt.Println(flag)

	flag = true
	fmt.Println(flag)

	x := 2
	y := 42

	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x != y)

	x = 42
	y = 42

	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x != y)

}
