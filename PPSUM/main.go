package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i := 0; i < testCases; i++ {
		var times, number int
		fmt.Fscan(os.Stdin, &times)
		fmt.Fscan(os.Stdin, &number)
		count := sum(times, number)
		fmt.Println("count -> ", count)

	}
}

func sum(d, n int) int {
	if d == 0 {
		return n
	}
	return sum(d-1, n*(n+1)/2)
}
