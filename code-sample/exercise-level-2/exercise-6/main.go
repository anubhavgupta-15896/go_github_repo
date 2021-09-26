package main

import "fmt"

const (
	_     = iota
	year1 = 2017 + iota
	year2 = 2017 + iota
	year3 = 2017 + iota
	year4 = 2017 + iota
)

func main() {

	fmt.Println(year1, year2, year3, year4)

}
