// Merging Streamed Metrics
//
// You have multiple monitoring agents each sending sorted metric streams (by timestamp). Merge
// these k sorted streams into a single sorted stream efficiently. Each stream is sorted by
// timestamp in ascending order.
//
// Difficulty: Medium
//
// Problem:
// Given k sorted arrays of metrics (each metric is [timestamp, value]), merge them into one
// sorted array ordered by timestamp. If multiple metrics have the same timestamp, maintain their
// relative order from the original streams (stable merge).
//
// Example 1:
// Input: streams = [[[1, 10], [3, 30]], [[2, 20], [4, 40]], [[1, 15], [5, 50]]]
// Output: [[1, 10], [1, 15], [2, 20], [3, 30], [4, 40], [5, 50]]
// Explanation: Merged and sorted by timestamp, preserving order for timestamp 1
//
// Example 2:
// Input: streams = [[[1, 100]], [[2, 200]], [[3, 300]]]
// Output: [[1, 100], [2, 200], [3, 300]]
// Explanation: Each stream has one element, simple merge
//
// Example 3:
// Input: streams = [[[5, 50], [10, 100]]]
// Output: [[5, 50], [10, 100]]
// Explanation: Only one stream, return as is
//
// Constraints:
// - k == streams.length
// - 0 <= k <= 100 (number of streams)
// - 0 <= streams[i].length <= 1000
// - streams[i][j].length == 2
// - 0 <= timestamp <= 1000000
// - -1000000 <= value <= 1000000
// - Each individual stream is sorted by timestamp in ascending order
