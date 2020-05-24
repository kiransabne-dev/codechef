package main

import "fmt"

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for t := 0; t < testCases; t++ {
		var numLen int
		fmt.Scan(&numLen)
		//    fmt.Println("t => ", numLen)
		data := make([]int, numLen)
		//    fmt.Println("data => ", data )
		for i := 0; i < numLen; i++ {
			//      fmt.Println("numLen => ", i, numLen)
			fmt.Scan(&data[i])
		}
		//    fmt.Println("data => ", data )

		ans := true

		for i := 0; i < numLen/2; i++ {
			if data[i] == data[(numLen-i)-1] &&
				(data[i+1]-data[i]) <= 1 &&
				data[numLen/2] == 7 &&
				data[0] == 1 {
				ans = true
			} else {
				ans = false
				break
			}
		}
		if ans {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}

}
