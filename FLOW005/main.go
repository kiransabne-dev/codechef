
package main

import (
	"fmt"
	"os"
)
//1, 2, 5, 10, 50, 100
func main(){
  var testCases int
  fmt.Fscan(os.Stdin, &testCases)
  for i := 0; i < testCases; i++{
    var amount int
    fmt.Fscan(os.Stdin, &amount)

        sum := 0;
        sum = sum + amount/100;
        amount=amount%100;
        fmt.Println(" ", sum)
        fmt.Println("amount1 ", amount)
        sum = sum + amount/50;
        amount=amount%50;
        fmt.Println(" ", sum)
        fmt.Println("amount2 ", amount)
        sum = sum + amount/10;
        amount=amount%10;
        fmt.Println(" ", sum)
        fmt.Println("amount3 ", amount)
        sum = sum + amount/5;
        amount=amount%5;
        sum = sum + amount/2;
        amount=amount%2;
        sum = sum + amount/1;
        amount=amount%1;

        fmt.Println("sum -> ", sum)
  }
}