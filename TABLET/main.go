package main

import (
	"fmt"
	"os"
)

func main() {
	var testCases int
	fmt.Fscan(os.Stdin, &testCases)

	for i := 0; i < testCases; i++ {
		fmt.Println("testCases -> ", testCases)
		//get itemCount and budget
		var items int
		var budget int
		fmt.Fscan(os.Stdin, &items)
		fmt.Fscan(os.Stdin, &budget)
		// looping over items
		var screenSize int
		for j := 0; j < items; j++ {
			var width, height, price int
			fmt.Fscan(os.Stdin, &width)
			fmt.Fscan(os.Stdin, &height)
			fmt.Fscan(os.Stdin, &price)
			if budget >= price {
				fmt.Println("inside budget")
				if (width * height) > screenSize {
					fmt.Println("screenSize -> ", width*height, " ", screenSize)
					screenSize = width * height
				}
			} else {
				fmt.Println("not inside budget")
			}
		}
		if screenSize != 0 {
			fmt.Println(screenSize)
		} else {
			fmt.Println("no tablet")
		}

	}
}
