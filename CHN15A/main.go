package main

import "fmt"

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var noOfMinions, initalCharacter int
		var count int

		fmt.Scan(&noOfMinions)
		fmt.Scan(&initalCharacter)
		//fmt.Println("N, I => ", noOfMinions, initalCharacter)
		for j := 0; j < noOfMinions; j++ {
			var temp int
			fmt.Scan(&temp)
			//fmt.Println("temp => ", temp)
			if (temp+initalCharacter)%7 == 0 {
				count++
			}
		}
		fmt.Println(count)
	}

}
