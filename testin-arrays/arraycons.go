package main

import "fmt"

func array() {
	var a int
	var i int
	fmt.Println("how long should teh array be")
	fmt.Scan(&a)
	if a != 0 {
		var numbs [a]int
	}

	if i < a {
		fmt.Println("give me ", i+1, " number of the array")
		fmt.Scan(&a)
		numbs[i] = a
		i = +1
	}

}
