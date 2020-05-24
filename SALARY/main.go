package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	var sum int64
	var temp, SMALLWAGE int
	var n int
	for t := 0; t < testCases; t++ {

		sum = 0

		fmt.Scanf("%d", &n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &temp)
			sum += int64(temp)

			if i == 0 {
				SMALLWAGE = temp
			} else if temp < SMALLWAGE {
				SMALLWAGE = temp
			}
		}
		//fmt.Println("min => ", min)
		fmt.Println(sum - int64(n*SMALLWAGE))

		// }
	}
}
