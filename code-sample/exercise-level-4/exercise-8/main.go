package main

import "fmt"

func main() {
	strArr1 := []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`}
	strArr2 := []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`}
	strArr3 := []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`}

	//strArr4 := [][]string{strArr1, strArr2, strArr3}
	m := map[string][][]string{
		"last_first": {strArr1, strArr2, strArr3},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\tThis is the record for", i)
			for j, v3 := range v2 {
				fmt.Println("\t\t", j, v3)
			}
		}
	}
}
