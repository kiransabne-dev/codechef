package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var testCases int
	var words string
	fmt.Fscan(os.Stdin, &testCases)
	for i := 0; i < testCases; i++ {
		fmt.Fscan(os.Stdin, &words)
		answer := main1(words)
		if answer {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main1(str string) bool {
	if len(str)%2 == 0 {

	}
	s1 := make([]int, 1000)
	s2 := make([]int, 1000)

	for i := 0; i < (len(str) / 2); i++ {
		s1 = append(s1, int(str[i]))
	}
	for i := (len(str) + 1) / 2; i < len(str); i++ {
		s2 = append(s2, int(str[i]))
	}

	sort.Ints(s1)
	sort.Ints(s2)
	ret := true
	for j := 0; j < len(s1); j++ {
		if s1[j] != s2[j] {
			ret = false
		}
	}
	return ret
}
