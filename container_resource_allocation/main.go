// Container Resource Allocation
//
// You are scheduling containers on a cluster with limited CPU and memory. Each container
// requires specific CPU and memory resources. Given the available resources and container
// requirements, determine the maximum number of containers that can be scheduled.
//
// Difficulty: Easy
//
// Problem:
// Given arrays of CPU and memory requirements for containers, and total available CPU and
// memory, find the maximum number of containers you can run simultaneously.
//
// Example 1:
// Input: cpu = [2, 3, 1, 4], memory = [4, 2, 1, 3], availableCPU = 7, availableMemory = 8
// Output: 3
// Explanation: Can run containers [2,1,4] using CPU=7 and memory=[4,1,3]=8
//
// Example 2:
// Input: cpu = [5, 5, 5], memory = [10, 10, 10], availableCPU = 10, availableMemory = 20
// Output: 2
// Explanation: Can run 2 containers (CPU=10, memory=20)
//
// Example 3:
// Input: cpu = [1, 1, 1], memory = [1, 1, 1], availableCPU = 10, availableMemory = 10
// Output: 3
// Explanation: All containers fit within available resources
//
// Constraints:
// - 1 <= containers <= 1000
// - cpu.length == memory.length
// - 1 <= cpu[i] <= 100 (CPU cores)
// - 1 <= memory[i] <= 1000 (GB)
// - 1 <= availableCPU <= 10000
// - 1 <= availableMemory <= 100000
// - You cannot partially allocate a container

package main

import (
	"fmt"
	"reflect"
)

func solution(cpu []int, memory []int, availableCPU int, availableMemory int) int {
	// Your implementation will go here.
	return 0
}

func main() {
	testCases := []struct {
		name            string
		cpu             []int
		memory          []int
		availableCPU    int
		availableMemory int
		expected        int
	}{
		{
			name:            "Example 1",
			cpu:             []int{2, 3, 1, 4},
			memory:          []int{4, 2, 1, 3},
			availableCPU:    7,
			availableMemory: 8,
			expected:        3,
		},
		{
			name:            "Example 2",
			cpu:             []int{5, 5, 5},
			memory:          []int{10, 10, 10},
			availableCPU:    10,
			availableMemory: 20,
			expected:        2,
		},
		{
			name:            "Example 3",
			cpu:             []int{1, 1, 1},
			memory:          []int{1, 1, 1},
			availableCPU:    10,
			availableMemory: 10,
			expected:        3,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input: cpu=%v, mem=%v, availCPU=%d, availMem=%d\n", tc.cpu, tc.memory, tc.availableCPU, tc.availableMemory)
		actual := solution(tc.cpu, tc.memory, tc.availableCPU, tc.availableMemory)
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
