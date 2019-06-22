// go1
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

}
func files() {
	file, err := os.Open("/home/dev/camera_host/Readme.md")
	if err != nil {
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	} // read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
func test() {
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
	var fi FrameInfo
	fi.frameNum = 11111111
	fi.timeStamp = 23363737466
	fi = FrameInfo{234234, 2363}
	buffer := []uint8{1, 2, 3, 4, 5, 6}
	frame := Frame{fi, &buffer}
	fmt.Println(frame.fi.getTime())
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
func swap(xptr, yptr *int) {
	tmp := *yptr
	*yptr = *xptr
	*xptr = tmp
}

type FrameInfo struct {
	frameNum, timeStamp uint64
}
type Frame struct {
	fi     FrameInfo
	buffer *[]uint8
}

func (fi *FrameInfo) getTime() uint64 {
	return fi.frameNum
}
