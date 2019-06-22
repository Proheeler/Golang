// go1
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go f()
	}
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

}
func f() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
