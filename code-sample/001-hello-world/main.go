package main

import "fmt"

func main() {
	fmt.Println("Hello everyone")
	foo()
	fmt.Println("something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("i am in foo")
}

func bar() {
	fmt.Println("and then we exit")
}

/*control flow:
sequence
loop
conditional
*/
