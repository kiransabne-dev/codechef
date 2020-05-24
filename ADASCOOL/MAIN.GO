package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i := 0; i < testCases; i++ {
		var rows int
		var columns int
		fmt.Fscan(os.Stdin, &rows)
		fmt.Fscan(os.Stdin, &columns)

		if (rows*columns)%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
