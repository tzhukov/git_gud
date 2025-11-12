// Last Stone Weight
//
// You are given an array of integers stones where stones[i] is the weight of the ith stone.
// We are playing a game with the stones. On each turn, we choose the heaviest two stones and
// smash them together. Suppose the heaviest two stones have weights x and y with x <= y.
// The result of this smash is:
// - If x == y, both stones are destroyed
// - If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x
//
// At the end of the game, there is at most one stone left. Return the weight of the last
// remaining stone. If there are no stones left, return 0.
//
// Difficulty: Easy
//
// Problem:
// Simulate the stone smashing game using a max heap to efficiently find the heaviest stones.
//
// Example 1:
// Input: stones = [2, 7, 4, 1, 8, 1]
// Output: 1
// Explanation: We combine 7 and 8 to get 1, then 4 and 1 to get 3, then 2 and 3 to get 1, then 1 and 1 to get 0
//
// Example 2:
// Input: stones = [1]
// Output: 1
// Explanation: Only one stone, return its weight
//
// Example 3:
// Input: stones = [2, 2]
// Output: 0
// Explanation: 2 and 2 smash to nothing
//
// Constraints:
// - 1 <= stones.length <= 30
// - 1 <= stones[i] <= 1000

package main

import (
	"fmt"
	"reflect"
)

func solution(stones []int) int {
	// Your implementation will go here.
	// Hint: Use a max heap to efficiently get the two heaviest stones
	return 0
}

func main() {
	testCases := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "Example 1",
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			name:     "Example 2",
			stones:   []int{1},
			expected: 1,
		},
		{
			name:     "Example 3",
			stones:   []int{2, 2},
			expected: 0,
		},
		{
			name:     "All stones equal",
			stones:   []int{5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "Multiple rounds",
			stones:   []int{3, 7, 2, 9, 4},
			expected: 1,
		},
		{
			name:     "Large values",
			stones:   []int{100, 200, 300},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input: stones=%v\n", tc.stones)
		actual := solution(tc.stones)
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
