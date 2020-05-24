package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		var A int
		var B int
		fmt.Fscan(os.Stdin, &A)
		fmt.Fscan(os.Stdin, &B)
		count := make([]int, 2)
		start := 1
		var winner string

		for true {
			if start%2 == 1 {
				count[0] = count[0] + start
				// fmt.Println("limak turn")
				if count[0] > A {
					winner = "Bob"
					break
				}
			} else {
				// fmt.Println("Bob turn")
				count[1] = count[1] + start
				if count[1] > B {
					winner = "Limak"
					break
				}
			}
			start++
		}
		fmt.Println(winner)
	}
}
