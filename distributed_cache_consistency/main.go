// Distributed Cache Consistency
//
// You are designing a distributed cache with multiple nodes. Given a log of cache operations
// (GET, PUT, INVALIDATE) from different nodes with timestamps, determine the final consistent
// state of the cache considering eventual consistency and conflict resolution rules.
//
// Difficulty: Hard
//
// Problem:
// Given operations from multiple cache nodes, each with [timestamp, nodeId, operation, key,
// value], determine the final state of the cache. Use last-write-wins (LWW) for conflicts.
// INVALIDATE removes the key if timestamp is newer than last write.
//
// Example 1:
// Input: ops = [[1, "n1", "PUT", "k1", "v1"], [2, "n2", "PUT", "k1", "v2"],
//               [3, "n1", "GET", "k1", null]]
// Output: {"k1": "v2"}
// Explanation: PUT at t=2 overwrites PUT at t=1 due to LWW
//
// Example 2:
// Input: ops = [[1, "n1", "PUT", "k1", "v1"], [2, "n2", "INVALIDATE", "k1", null],
//               [3, "n3", "PUT", "k1", "v3"]]
// Output: {"k1": "v3"}
// Explanation: Invalidate at t=2, then PUT at t=3 sets new value
//
// Example 3:
// Input: ops = [[5, "n1", "PUT", "k1", "v5"], [3, "n2", "PUT", "k1", "v3"],
//               [7, "n1", "INVALIDATE", "k1", null]]
// Output: {}
// Explanation: PUT v5 at t=5 is latest write, INVALIDATE at t=7 removes it
//
// Constraints:
// - 1 <= operations <= 10000
// - 0 <= timestamp <= 1000000 (monotonically increasing per node)
// - Timestamps can be concurrent across nodes
// - 1 <= keys, values <= 100 characters
// - Operations: PUT, GET, INVALIDATE
// - GET operations don't change state (read-only)
// - For concurrent operations (same timestamp), use nodeId for tie-breaking (lexicographic)

package main

import (
	"fmt"
	"reflect"
)

func solution(ops [][]any) map[string]string {
	// Your implementation will go here.
	return nil
}

func main() {
	testCases := []struct {
		name     string
		ops      [][]any
		expected map[string]string
	}{
		{
			name: "Example 1",
			ops: [][]any{
				{1, "n1", "PUT", "k1", "v1"},
				{2, "n2", "PUT", "k1", "v2"},
				{3, "n1", "GET", "k1", nil},
			},
			expected: map[string]string{"k1": "v2"},
		},
		{
			name: "Example 2",
			ops: [][]any{
				{1, "n1", "PUT", "k1", "v1"},
				{2, "n2", "INVALIDATE", "k1", nil},
				{3, "n3", "PUT", "k1", "v3"},
			},
			expected: map[string]string{"k1": "v3"},
		},
		{
			name: "Example 3",
			ops: [][]any{
				{5, "n1", "PUT", "k1", "v5"},
				{3, "n2", "PUT", "k1", "v3"},
				{7, "n1", "INVALIDATE", "k1", nil},
			},
			expected: map[string]string{},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("--- Running Test Case: %s ---\n", tc.name)
		fmt.Printf("Input ops: %v\n", tc.ops)
		actual := solution(tc.ops)
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
