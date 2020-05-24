package main

import (
	"fmt"
	s "strings"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var n, k int
		fmt.Scan(&n)
		fmt.Scan(&k)
		fmt.Println("n ,k ", n, k)
		dictionary := make([]string, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&dictionary[i])
		}
		fmt.Println("dictionary => ", dictionary)
		//checking

		var mainslice []string
		list := make(map[string]string)

		fmt.Println(mainslice)

		for j := 0; j < k; j++ {
			var noOfDescription int
			fmt.Scan(&noOfDescription)

			//words list as per the noOfDescription
			phrases := make([]string, noOfDescription)
			for p := 0; p < noOfDescription; p++ {
				fmt.Scan(&phrases[p])
				mainslice = append(mainslice, phrases[p])
				list[phrases[p]] = phrases[p]
			}
			fmt.Println("phrases => ", list)

			//   for m := 0; m < n; m++{
			//   fmt.Println("dictionary[m] => ",dictionary[m])
			//   if dictionary[m] ==
			// }

		}

		//answers := make([]string, n)
		str := make([]string, n)
		fmt.Println("mainslice", mainslice)
		for i, v := range dictionary {
			_, ok := list[v]
			fmt.Println("found => ", ok)
			if ok {
				str[i] = "YES"
			} else {
				str[i] = "NO"
			}
		}
		fmt.Printf(s.Join(str, " ") + "\n")

	}
}

// Example
// Input:
// 2
// 3 2
// piygu ezyfo rzotm
// 1 piygu
// 6 tefwz tefwz piygu ezyfo tefwz piygu
// ---------------------------------
// 4 1
// kssdy tjzhy ljzym kegqz
// 4 kegqz kegqz kegqz vxvyj

// Output:
// YES YES NO
// NO NO NO YES
