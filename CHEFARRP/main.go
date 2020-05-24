package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		var size int
		fmt.Fscan(os.Stdin, &size)
		line := make([]int, size)
		for j, _ := range line {
			fmt.Fscan(os.Stdin, &line[j])
		}
		fmt.Println("line -> ", line)
		ans := getCount(line, size)
		fmt.Println("ans -> ", ans)
	}
}

func getCount(input []int, size int) int {
	count := 0
	for i := 1; i <= size; i++ {
		fmt.Println("i -> ", i)
		sum := 0
		product := 1

		for j := i; j > 0; j-- {
			fmt.Println("j -> ", j)
			sum = sum + input[j-1]
			product = product * input[j-1]

			if sum == product {
				count++
			}
		}
	}
	return count
}
