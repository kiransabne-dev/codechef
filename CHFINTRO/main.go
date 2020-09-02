package main

import (
	"fmt"
)

func main() {

	var N, r int
	fmt.Scan(&N)
	fmt.Scan(&r)
	for i := 1; i <= N; i++ {

		var R int
		fmt.Scan(&R)

		if R >= r {
			fmt.Println("Good boi")
		} else {
			fmt.Println("Bad boi")
		}
	}
}
