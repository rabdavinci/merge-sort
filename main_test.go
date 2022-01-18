package main

import (
	"testing"
)

var is []int = generateIntSlice(10000)

func BenchmarkSequential(b *testing.B) {
	mergeSort(is)
}
func BenchmarkParallel(b *testing.B) {
	mergeSortParallel(is)
}
