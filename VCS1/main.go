package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var noOfFile, noOfIgnoredFile, noOfTrackedFile int
		fmt.Scan(&noOfFile)
		fmt.Scan(&noOfIgnoredFile)
		fmt.Scan(&noOfTrackedFile)

		fileList := make(map[int][]string)

		for i := 1; i <= noOfFile; i++ {
			fileList[i] = append(fileList[i])
		}

		// ignoredFiles list
		for j := 1; j <= noOfIgnoredFile; j++ {
			var fileNo int
			fmt.Scan(&fileNo)
			fileList[fileNo] = append(fileList[fileNo], "Y")
		}

		for j := 1; j <= noOfTrackedFile; j++ {
			var fileNo int
			fmt.Scan(&fileNo)
			fileList[fileNo] = append(fileList[fileNo], "Y")
		}

		//map[1:[Y Y] 2:[Y] 3:[Y] 4:[Y Y] 5:[] 6:[Y Y] 7:[Y Y]]
		var trackedAndIgnored, untrackedAndUnignored int
		for _, value := range fileList {

			if len(value) == 2 {
				trackedAndIgnored = trackedAndIgnored + 1
			} else if len(value) == 0 {
				untrackedAndUnignored = untrackedAndUnignored + 1
			}
		}
		fmt.Println(trackedAndIgnored, untrackedAndUnignored)
	}
}
