package main

import (
	"fmt"

	"github.com/WhereIsMyLittleBear/dataStructure/sortfunc"
)

func main() {
	num := [...]int{9, 6, 11, 3, 5, 12, 8, 7, 15, 14, 4, 1, 13, 2}
	fmt.Println(num)
	sortfunc.QuickSort(num[:])
	fmt.Println(num)
}
