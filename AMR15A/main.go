package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println("main")
	var N, noOfLucky, noOfOdd int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		log.Println("N -> ", N)
		var weapons int
		fmt.Scan(&weapons)
		if weapons%2 == 0 {
			//log.Println("even")
			noOfLucky = noOfLucky + 1
		} else {
			//log.Println("odd")
			noOfOdd = noOfOdd + 1
		}

	}
	log.Println(noOfLucky, noOfOdd)
	if noOfLucky > noOfOdd {
		fmt.Println("READY FOR BATTLE")
	} else {
		fmt.Println("NOT READY")
	}
}
