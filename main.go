package main

import (
	"fmt"
)

func main() {
	is := generateIntSlice(1000)
	fmt.Println(is)
	fmt.Println(mergeSort(is))
	fmt.Println(mergeSortParallel(is))
}
