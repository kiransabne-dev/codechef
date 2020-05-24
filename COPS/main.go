// Input:
// 3
// 4 7 8
// 12 52 56 8
// 2 10 2
// 21 75
// 2 5 8
// 10 51

// Output:
// 0
// 18
// 9

package main

import (
	"fmt"
)

func main() {

	var testCases int
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		fmt.Println("t => ", t)

		dataSet := make([]int, 110)
		fmt.Println(dataSet)

		var policeHouses, housePerMin, minutes int
		fmt.Scan(&policeHouses)
		fmt.Scan(&housePerMin)
		fmt.Scan(&minutes)
		reach := housePerMin * minutes
		// police roomno Input
		var pRoomNo int
		for i := 0; i < policeHouses; i++ {
			fmt.Scan(&pRoomNo)
			dataSet[pRoomNo] = 1
			for j := 1; j <= reach && (pRoomNo+j) <= 100; j++ {
				dataSet[pRoomNo+j] = 1
			}
			for j := 1; j <= reach && (pRoomNo-j) >= 1; j++ {
				dataSet[pRoomNo-j] = 1
			}
		}
		fmt.Println(dataSet)

		answer := 0
		for k := 1; k <= 100; k++ {
			if dataSet[k] == 0 {
				answer++
			}
		}

		fmt.Println("answer => ", answer)
	}
}
