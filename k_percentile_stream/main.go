package main

import (
	"container/heap"
	"fmt"
)

func main() {

	h := IntHeap{0}

	data := []int{}

	heap.Init(&h)

	data1 := []int{4, 1, 3, 2, 16, 4, 9, 10, 14, 8, 7}

	data2 := []int{7, 15, 5, 20, 3, 8, 7, 6, 12, 1, 18}
	data3 := []int{4, 6, 8, 2, 10, 12, 5, 14, 16}

	p95_1 := streamedPercentile(&h, &data, &data1, 95)
	fmt.Println("95th Percentile after data1:", p95_1)

	p95_2 := streamedPercentile(&h, &data, &data2, 95)
	fmt.Println("95th Percentile after data2:", p95_2)

	p95_3 := streamedPercentile(&h, &data, &data3, 95)
	fmt.Println("95th Percentile after data3:", p95_3)
}

func streamedPercentile(h *IntHeap, data *[]int, dataNew *[]int, percentile int) int {

	ln := len(*data) + len(*dataNew)

	pIndex := ln - (percentile*ln)/100 + 1

	fmt.Println("Current total data length:", ln, "Percentile index:", pIndex)

	for _, v := range *dataNew {
		if len(*h) < pIndex {
			heap.Push(h, v)
		} else {
			if (*h)[0] < v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
		*data = append(*data, v)
	}
	return (*h)[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	ret := old[n-1]
	*h = old[0 : n-1]
	return ret
}
func (h *IntHeap) Push(n any) {
	(*h) = append((*h), n.(int))
}
