// Window Max Latency Spikes
//
// You are analyzing request latencies for a web service. Given an array of latency measurements
// (in milliseconds) and a time window size, find the maximum latency spike within any window of
// that size. A spike is the difference between the maximum and minimum latency in the window.
//
// Difficulty: Medium
//
// Problem:
// Given an array of latency measurements and a window size k, find the maximum spike (max - min)
// across all possible windows of size k.
//
// Example 1:
// Input: latencies = [100, 200, 150, 300, 250, 400], k = 3
// Output: 150
// Explanation: Window [150, 300, 250] has spike of 300-150=150, which is the maximum
//
// Example 2:
// Input: latencies = [50, 60, 55, 58, 52], k = 2
// Output: 10
// Explanation: Window [50, 60] has the maximum spike of 10
//
// Example 3:
// Input: latencies = [1000, 100, 200, 150], k = 4
// Output: 900
// Explanation: Only one window [1000, 100, 200, 150] with spike 1000-100=900
//
// Constraints:
// - 1 <= k <= latencies.length <= 10000
// - 0 <= latencies[i] <= 1000000 (latencies in milliseconds)
// - All latency values are non-negative integers
