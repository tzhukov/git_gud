// Kth Largest Element in an Array
//
// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// You can assume k is always valid, 1 ≤ k ≤ array's length.
//
// Difficulty: Easy
//
// Problem:
// Find the kth largest element in an unsorted array using a heap-based approach.
// This is a classic heap problem that can be solved efficiently using a min heap of size k.
//
// Example 1:
// Input: nums = [3, 2, 1, 5, 6, 4], k = 2
// Output: 5
// Explanation: The 2nd largest element is 5
//
// Example 2:
// Input: nums = [3, 2, 3, 1, 2, 4, 5, 5, 6], k = 4
// Output: 4
// Explanation: The 4th largest element is 4
//
// Example 3:
// Input: nums = [1], k = 1
// Output: 1
// Explanation: Single element is the largest
//
// Constraints:
// - 1 <= k <= nums.length <= 10^5
// - -10^4 <= nums[i] <= 10^4

package main

import (
	"fmt"
	"reflect"
)

func solution(nums []int, k int) int {
	// Your implementation will go here.
	// Hint: Use a min heap of size k
	return 0
}

func main() {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5, 5},
			k:        3,
			expected: 5,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -5, -3, -2, -4},
			k:        2,
			expected: -2,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input: nums=%v, k=%d\n", tc.nums, tc.k)
		actual := solution(tc.nums, tc.k)
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
