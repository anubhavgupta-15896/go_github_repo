package main

import "fmt"

func main() {
	favSport := "Hello"
	switch favSport {
	case "Hello":
		fmt.Println("Hello")
		fallthrough
	case "Hii":
		fmt.Println("Hii")
	case "Good Morning":
		fmt.Println("Good Morning")
	case "Bye":
		fmt.Println("Bye")
	}
}
