// wrong
// Example Input

// 2
// 10 4
// 10 8

// Example Output

// 0
// 1

// Explanation

// Example case 1: Chef attacks with power 4
// , Darth's health becomes 6. Chef attacks with power 2, Darth's health becomes 4. Chef attacks with power 1 and Darth's health becomes 3, but Chef's attack power becomes 0

// .

// Example case 2: Chef attacks with power 8
// , Darth's health becomes 2. Chef attacks with power 4, Darth's health becomes 0. Chef kills Darth.

// If the attack power becomes 0 before Darth's health becomes 0 or less, Chef dies. Otherwise, Darth dies. You are given Darth's initial health H and the initial attack power P. Tell Chef if he can beat Darth or if he should escape.

package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		//log.Println("testCases -> ", i)
		var health, power int
		fmt.Scan(&health)
		fmt.Scan(&power)
		//log.Println("health -> ", health, " and power -> ", power)
		// for attack := power; attack >= 0; attack-- {

		// 	health = health - power
		// 	power = power / 2
		// 	log.Println("attack -> ", health, " power ", power)
		// 	if health >= 0 and power <= 0 {
		// 		log.Println("sol ", 0)
		// 		break
		// 	} else {
		// 		log.Println("sol ", 1)
		// 		break
		// 	}
		// }
		// for j := health; j > 0; j-- {
		// 	log.Println("health", health)
		// }
		for health >= 0 {
			//log.Println("health -> ", health)

			health = health - power
			power = power / 2
			//log.Println("health -> ", health, " power -> ", power)
			if power <= 0 {
				fmt.Println(0)
				break
			} else if health <= 0 && power > 0 {
				fmt.Println(1)
				break
			}
		}

	}
}
