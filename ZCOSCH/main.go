package main

import (
	"fmt"
	"os"
)

func main() {
	var rank int64
	fmt.Fscan(os.Stdin, &rank)
	if rank <= 50 {
		fmt.Println(100)
	} else if rank > 50 && rank <= 100 {
		fmt.Println(50)
	} else {
		fmt.Println(0)
	}
}
