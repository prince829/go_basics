package slice

import "fmt"

func Slice() {
	//Slice has dynamic length
	//[]int means a sequence of ints
	// can grow or shrink easily
	num := []int{10, 20, 30}
	num = append(num, 40)
	fmt.Println(num)
}
