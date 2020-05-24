package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i := 0; i < testCases; i++ {
		var inputSize int
		fmt.Fscan(os.Stdin, &inputSize)
		frequency := make([]int, 101)
		data := make([]int, inputSize)
		for pos, _ := range data {
			fmt.Fscan(os.Stdin, &data[pos])
			frequency[data[pos]]++
		}
		sort.Ints(frequency)
		// fmt.Println("frequency -> ", frequency)
		// fmt.Println("data -> ", data)
		fmt.Println(inputSize - frequency[100])
	}

}
