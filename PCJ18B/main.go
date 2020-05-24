package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		var N int
		fmt.Fscan(os.Stdin, &N)
		count := getCount(N)
		fmt.Println("count  -> ", count)
	}
}

func getCount(n int) int {
	count := 0
	for i := 1; i <= n; i = i + 2 {
		//fmt.Println("i -> ", i)
		k := n - i + 1
		count = count + (k * k)
	}
	return count
}
