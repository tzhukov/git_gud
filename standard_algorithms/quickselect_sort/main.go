package main

import "fmt"

func main() {
	array := []int{2, 7, 4, 7, 3, 7}

	quicksort(array, 0, len(array)-1)
	fmt.Println(array)

}

func quicksort(array []int, start int, end int) {
	if start < end {
		pivot := partition(array, start, end)
		quicksort(array, start, pivot-1)
		quicksort(array, pivot+1, end)
	}

}

func partition(A []int, start int, end int) int {
	i := start + 1
	pivot := A[start]

	for j := start + 1; j <= end; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[start], A[i-1] = A[i-1], A[start]
	return i - 1
}
