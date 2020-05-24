package main

import (
	"fmt"
	"os"
)

func main() {
	line := []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	//testcases loop
	for i := 0; i < testCases; i++ {
		var num1 int
		var num2 int
		fmt.Fscan(os.Stdin, &num1)
		fmt.Fscan(os.Stdin, &num2)
		sum := num1 + num2

		sizeOfSum := countDigitsRecur(sum)

		// loop over ans each digit
		matchesRequired := 0
		for j := sizeOfSum; j > 0; j-- {
			ans := sum % 10
			matchesRequired = matchesRequired + line[ans]
			sum = sum / 10
		}
		fmt.Println(matchesRequired)

	}

}

func countDigitsRecur(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + countDigitsRecur(number/10)
	}
}
