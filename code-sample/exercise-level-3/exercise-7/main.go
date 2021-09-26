package main

import "fmt"

func main() {
	x := 40
	if x == 42 {
		fmt.Println("x is 42")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else {
		fmt.Println("x is not 41 and 42")
	}
}
