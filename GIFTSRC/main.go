// GIFTSRC
// 'L' means to go left, i.e. to the cell (x−1,y)
// 'R' means to go right, i.e. to the cell (x+1,y)
// 'U' means to go up, i.e. to the cell (x,y+1)
// 'D' means to go down, i.e. to the cell (x,y−1)

// Example Input
// 3
// 5
// LLLUR
// 7
// LLLRUUD
// 8
// LRULLUDU
// Example Output
// 0 1
// -1 1
// -2 2

// (0,0)→(−1,0)→(−1,0)→(−1,0)→(−1,1)→(0,1)
// (0,0) -> (-1,0) -> (-1,0) -> (-1,0), (-1,0), (-1,1)
// (0,0) -> (-1,0) -> (-1,0) -> (-1,1), (-2,1)-> (-2,2)

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hwllo world.")
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	for t := 0; t < testCases; t++ {
		fmt.Println(t)
		var instructionLen int
		fmt.Fscan(os.Stdin, &instructionLen)
		var instruction string
		fmt.Fscan(os.Stdin, &instruction)
		var x, y int
		fmt.Println("inst -> ", instruction, instructionLen)
		var previousInstruction byte

		//first instruction
		move(&x, &y, instruction[0])
		previousInstruction = instruction[0]
		//loop for instruction
		for i := 1; i < instructionLen; i++ {
			fmt.Println(previousInstruction, instruction[i])
			if ((instruction[i] == 'L' || instruction[i] == 'R') && (previousInstruction == 'L' || previousInstruction == 'R')) ||
				((instruction[i] == 'U' || instruction[i] == 'D') && (previousInstruction == 'U' || previousInstruction == 'D')) {
				previousInstruction = instruction[i]
				continue
			}
			previousInstruction = instruction[i]
			move(&x, &y, instruction[i])
		}
		//var one byte = instruction[0]
		//fmt.Println(one)
		fmt.Println(x, y)
	}
}

func move(x, y *int, instruction byte) {
	switch instruction {
	case 'L':
		*x--
		return
	case 'R':
		*x++
		return
	case 'U':
		*y++
		return
	case 'D':
		*y--
		return
	default:
		fmt.Println("done something wrong")
		return
	}
}
