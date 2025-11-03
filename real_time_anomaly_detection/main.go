// Real-Time Anomaly Detection
//
// You are building a real-time monitoring system that detects anomalies in metric streams. Given
// a stream of numerical metrics, detect when a value is anomalous based on the moving average
// and standard deviation of the previous N values.
//
// Difficulty: Medium
//
// Problem:
// Implement a function that processes a stream of metrics and returns indices of anomalous
// values. A value is considered anomalous if it deviates from the moving average by more than
// k standard deviations. Use a sliding window of size N for calculations.
//
// Example 1:
// Input: metrics = [10, 12, 11, 13, 12, 50, 11, 10], N = 5, k = 2.0
// Output: [5]
// Explanation: At index 5, value 50 is more than 2 std devs away from avg of [10,12,11,13,12]
//
// Example 2:
// Input: metrics = [100, 102, 101, 103, 102, 104, 103], N = 4, k = 3.0
// Output: []
// Explanation: All values are within 3 standard deviations of their moving average
//
// Example 3:
// Input: metrics = [5, 5, 5, 5, 100, 5, 5], N = 3, k = 1.5
// Output: [4]
// Explanation: Value 100 at index 4 is anomalous compared to previous window [5, 5, 5]
//
// Constraints:
// - 1 <= metrics.length <= 10000
// - -1000000 <= metrics[i] <= 1000000
// - 1 <= N <= 1000
// - N < metrics.length
// - 0.1 <= k <= 10.0
// - Return empty array if there are fewer than N values to establish baseline
