package main

import "fmt"

var y string
var z int

func main() {
	//DECLARE avarriable to be a certain type
	//and thnen ASSIGHN a VALUE of that type to the varriable

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "James Bond"
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
