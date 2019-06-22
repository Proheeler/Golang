// go1
package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	go f()
	// }
	arr()

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	max()
	fmt.Println(fib(30))
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)

	// fmt.Print("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)
	// output := input * 2
	// fmt.Println(output)

}
func f() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func arr() {
	x := []float64{98, 93, 77, 82, 83}
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	slice_x := x[:3]
	for i := 0; i < len(slice_x); i++ {
		fmt.Println(slice_x[i])
	}
	fmt.Println(average(x))
}
func max() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	max_el := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] > max_el {
			max_el = x[i]
		}
	}
	fmt.Println(max_el)
}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func fib(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}
func swap(xptr *int, yptr *int) {
	tmp := *yptr
	*yptr = *xptr
	*xptr = tmp
}
