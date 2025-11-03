// Top K Frequent Elements
//
// You are analyzing logs from a distributed system and need to find the K most frequently
// occurring error codes. Given a list of error codes and an integer k, return the k most
// frequent error codes. If multiple codes have the same frequency, return any k of them.
//
// Difficulty: Easy
//
// Problem:
// Given an array of error codes (integers) and an integer k, return the k most frequent error
// codes in any order.
//
// Example 1:
// Input: errorCodes = [404, 500, 404, 503, 404, 500, 200], k = 2
// Output: [404, 500]
// Explanation: 404 appears 3 times, 500 appears 2 times (most frequent)
//
// Example 2:
// Input: errorCodes = [200, 200, 201, 201, 202], k = 3
// Output: [200, 201, 202] (or any permutation)
// Explanation: All codes have similar frequencies, return any k
//
// Example 3:
// Input: errorCodes = [500], k = 1
// Output: [500]
// Explanation: Only one unique error code
//
// Constraints:
// - 1 <= k <= number of unique error codes <= errorCodes.length
// - 1 <= errorCodes.length <= 10000
// - 100 <= errorCodes[i] <= 599 (valid HTTP status codes)
// - It is guaranteed that the answer is unique (for the top k frequencies)
