// Explanation

// Example case 1: Chef and Rick generate the optimal integers 3
// and 5 respectively. Each of them has 1

// digit, so Rick cheats and wins the game.

// Example case 2: Chef and Rick could generate e.g. 6877
// and 99 respectively. Since Rick's integer has 2 digits and Chef cannot generate an integer with less than 4

// digits, Rick wins.

// Example case 3: Chef and Rick could generate e.g. 86
// and 888 respectively. Chef's integer has 2 digits and Rick cannot generate an integer with less than 3 digits, so Chef wins.
// 1, 9, 18, 27, 36, ...
// 1, 9, 99, 999, 9999, 99999, ...

// Example Input

// 3
// 3 5
// 28 18
// 14 24

// Example Output

// 1 1
// 1 2
// 0 2

package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello")
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		// log.Println("testCases -> ", testCases)
		var pc, pr, digitC, digitR int
		fmt.Scan(&pc)
		fmt.Scan(&pr)
		if pc%9 == 0 {
			digitC = pc / 9
		} else {
			digitC = (pc / 9) + 1
		}

		if pr%9 == 0 {
			digitR = pr / 9
		} else {
			digitR = (pr / 9) + 1
		}

		if digitC == digitR || digitR < digitC {
			fmt.Println(1, " ", digitR)
		} else if digitC < digitR {
			fmt.Println(0, " ", digitC)
		}
	}
}
