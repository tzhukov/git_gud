// Prioritizing Job Queue
//
// You are implementing a job scheduler that processes jobs based on priority and arrival time.
// Each job has a priority level (1-10, where 10 is highest) and an arrival timestamp. Process
// jobs in order of priority, and for jobs with equal priority, process them in FIFO order.
//
// Difficulty: Easy
//
// Problem:
// Given a list of jobs where each job is [jobId, priority, arrivalTime], return the order in
// which jobs should be processed.
//
// Example 1:
// Input: jobs = [[1, 5, 100], [2, 10, 101], [3, 5, 102]]
// Output: [2, 1, 3]
// Explanation: Job 2 has highest priority. Jobs 1 and 3 have equal priority, process by arrival
//
// Example 2:
// Input: jobs = [[1, 3, 50], [2, 3, 60], [3, 7, 55]]
// Output: [3, 1, 2]
// Explanation: Job 3 has priority 7. Jobs 1,2 have priority 3, ordered by arrival time
//
// Example 3:
// Input: jobs = [[1, 10, 100]]
// Output: [1]
// Explanation: Only one job to process
//
// Constraints:
// - 1 <= jobs.length <= 1000
// - jobs[i].length == 3
// - 1 <= jobId <= 10000 (unique job identifiers)
// - 1 <= priority <= 10
// - 0 <= arrivalTime <= 1000000 (timestamp in seconds)
// - All jobIds are unique

package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return LessHelper(h[i], h[j])
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.([]int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LessHelper(a, b []int) bool {
	if a[1] > b[1] {
		return true
	} else if a[1] == b[1] {
		return a[2] < b[2]
	}
	return false
}

func solution(jobs [][]int) []int {
	pq := IntHeap{}
	heap.Init(&pq)

	r := []int{}

	for _, job := range jobs {
		heap.Push(&pq, job)
		fmt.Println(pq)
	}

	for pq.Len() > 0 {
		r = append(r, heap.Pop(&pq).([]int)[0])
	}

	return r
}

func main() {
	testCases := []struct {
		name     string
		jobs     [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			jobs:     [][]int{{1, 5, 100}, {2, 10, 101}, {3, 5, 102}},
			expected: []int{2, 1, 3},
		},
		{
			name:     "Example 2",
			jobs:     [][]int{{1, 3, 50}, {2, 3, 60}, {3, 7, 55}},
			expected: []int{3, 1, 2},
		},
		{
			name:     "Example 3",
			jobs:     [][]int{{1, 10, 100}},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input jobs: %v\n", tc.jobs)
		actual := solution(tc.jobs)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual:   %v\n", actual)

		if reflect.DeepEqual(actual, tc.expected) {
			fmt.Println("Result: PASS")
		} else {
			fmt.Println("Result: FAIL")
		}
		fmt.Println()
	}
}
