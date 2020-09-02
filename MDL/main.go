// Example Input
// 2
// 3
// 2 5 3
// 10
// 38 9 102 10 96 7 46 28 88 13
// Example Output
// 2 5
// 102 7
// Explanation
// Example case 1: We remove the element 3 because it is the median of (2,5,3). The final sequence is (2,5).
package main

import (
	"fmt"
	"os"
)

func main() {
	//log.Println("hellow")
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for t := 0; t < testCases; t++ {
		var N int
		fmt.Fscan(os.Stdin, &N)
		line := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(os.Stdin, &line[i])
		}

		min, max, minIndex, maxIndex := MinMax(line)
		//		log.Println(min, max, minIndex, maxIndex)

		if maxIndex > minIndex {
			fmt.Println(min, max)
		} else {
			fmt.Println(max, min)
		}
	}
}

// 109 108 107 106 105 104 103 102 101 100 99 98 97 96 95 94 93 92 91 90 89 88 87 86 85 84 83 82 81 80 79 78 77 76 75 74 73 72 71 70 69 68 67 66 65 64 63 62 61 60 59 58 57 56 55 54 53 52 51 50 49 48 47 46 45 44 43 42 41 40 39 38 37 36 35 34 33 32 31 30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0
func MinMax(array []int) (int, int, int, int) {
	var max int = array[0]
	var min int = array[0]
	var minIndex, maxIndex int
	for index, value := range array {
		if max < value {
			max = value
			maxIndex = index
		}
		if min > value {
			min = value
			minIndex = index
		}
	}
	return min, max, minIndex, maxIndex
}
