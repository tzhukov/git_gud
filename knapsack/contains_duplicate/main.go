// Contains Duplicate
//
// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
//
// Difficulty: Easy
//
// Problem:
// Given an array of integers, determine if the array contains any duplicate values.
// Return true if any value appears more than once, false if all elements are unique.
//
// Example 1:
// Input: nums = [1, 2, 3, 1]
// Output: true
// Explanation: The value 1 appears twice
//
// Example 2:
// Input: nums = [1, 2, 3, 4]
// Output: false
// Explanation: All elements are distinct
//
// Example 3:
// Input: nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
// Output: true
// Explanation: Multiple values appear more than once
//
// Constraints:
// - 1 <= nums.length <= 10^5
// - -10^9 <= nums[i] <= 10^9

package main

import (
	"fmt"
	"reflect"
)

func solution(nums []int) bool {
	// Your implementation will go here.
	return false
}

func main() {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "Two same elements",
			nums:     []int{5, 5},
			expected: true,
		},
		{
			name:     "Large array with duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input: nums=%v\n", tc.nums)
		actual := solution(tc.nums)
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
