package main

import "fmt"

//GLOBAL SCOPE VARRIABLE DECLATION
var y = 32
var i int = 23
var j int

func main() {
	//LICAL SCOPE VARRIABLE DECLATION
	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	fmt.Println(y)
	fmt.Println(i)
	fmt.Println(j)

	z := "James Bond"
	fmt.Println(z)
	bar()
}

func bar() {
	fmt.Println(y)
	fmt.Println(i)
}
