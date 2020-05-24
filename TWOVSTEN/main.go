package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		var x int64
		fmt.Fscan(os.Stdin, &x)
		if x%10 == 0 {
			fmt.Println(0)
		} else if x%5 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(-1)
		}
	}
}
