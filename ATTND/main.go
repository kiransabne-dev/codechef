//ATTND
//  Example Input
//  1
// 4
//  hasan jaddouh
//  farhod khakimiyon
//  kerim kochekov
//  hasan khateeb
//
//  Example Output
//  hasan jaddouh
//  farhod
//  kerim
//  hasan khateeb

package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)
	fmt.Println("testCases -> ", testCases)
	for t := 0; t < testCases; t++ {
		var nameCount int
		fmt.Fscan(os.Stdin, &nameCount)
		fmt.Println("nameCount -> ", nameCount)
		nameSlice := make([][2]string, nameCount)
		nameCountMap := make(map[string]int)
		for i := 0; i < nameCount; i++ {
			fmt.Println("name -> ", nameSlice[i][0])
			fmt.Fscan(os.Stdin, &nameSlice[i][0])
			fmt.Fscan(os.Stdin, &nameSlice[i][1])
			_, present := nameCountMap[nameSlice[i][0]]
			fmt.Println("prs:", present)
			if present {
				nameCountMap[nameSlice[i][0]]++
			} else {
				nameCountMap[nameSlice[i][0]] = 1
			}

		}
		fmt.Println("nameslcie -> ", nameSlice, " ", nameCountMap)

		// range through slices for Outputu
		for i, v := range nameSlice {
			fmt.Println("i -> ", i, "v -> ", v)
			fmt.Println(nameCountMap[v[0]])
			if nameCountMap[v[0]] > 1 {
				fmt.Println("stdout -> ", v[0]+" "+v[1])
			} else {
				fmt.Println("stdout -> ", v[0])
			}

		}
	}
}

// 102  7
