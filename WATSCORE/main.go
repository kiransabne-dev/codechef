package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		var submissions int
		fmt.Fscan(os.Stdin, &submissions)
		fmt.Println("submissions -> ", submissions)

		scoreBoard := make(map[int]int)
		for j := 0; j < submissions; j++ {
			var question int
			var score int
			fmt.Fscan(os.Stdin, &question)
			fmt.Fscan(os.Stdin, &score)
			if question < 9 {
				_, ok := scoreBoard[question]
				fmt.Println("ok -> ", ok)
				if ok {
					if scoreBoard[question] < score {
						scoreBoard[question] = score
					}
				} else {
					scoreBoard[question] = score
				}
			}
		}
		fmt.Println("scoreboard -> ", scoreBoard)
		count := 0
		for _, val := range scoreBoard {
			count = count + val
		}
		fmt.Println("marks -> ", count)
	}
}
