package main

import (
	"math/rand"
	"sync"
	"time"
)

const MinSliceLenForParallel = 1 << 11

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	//fmt.Println(items)
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func mergeSortParallel(items []int) []int {
	if len(items) < 2 {
		return items
	}
	if len(items) <= MinSliceLenForParallel { // Sequential
		mergeSort(items)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	cfirst := make(chan []int)
	csecond := make(chan []int)
	go func() {
		defer wg.Done()
		cfirst <- mergeSortParallel(items[:len(items)/2])
	}()
	go func() {
		defer wg.Done()
		csecond <- mergeSortParallel(items[len(items)/2:])
	}()

	//fmt.Println(items)
	first := <-cfirst
	second := <-csecond
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	// fmt.Println(final)
	return final
}

func generateIntSlice(n int) (a []int) {
	rand.Seed(time.Now().Unix())
	a = rand.Perm(1000)
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return
}
