package main

import (
	"fmt"
	"log"
	s "strings"
)

func main() {
	fmt.Println("test")
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		log.Println("testXases -> ", testCases)
		var N, total int
		// var total int64
		fmt.Scan(&N)
		fmt.Scan(&total)
		ans := make([]string, N)

		for j := 1; j <= N; j++ {
			var amountRequested int
			fmt.Scan(&amountRequested)
			if amountRequested <= total {
				log.Println("1")
				total = total - amountRequested
				ans = append(ans, "1")
			} else {
				log.Println("0")
				ans = append(ans, "0")
			}
		}
		//log.Println("ans -> ", ans)
		fmt.Println(s.Join(ans, ""))

	}
}
