package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

func main() {

	h := IntHeap{}

	heap.Init(&h)

	for i := 0; i < 5; i++ {
		num := rand.Intn(10)
		fmt.Print(num)
		heap.Push(&h, num)
	}

	fmt.Println()
	fmt.Println(h)

	for i := 0; i < 5; i++ {
		fmt.Print(heap.Pop(&h))
	}
	fmt.Println()

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
