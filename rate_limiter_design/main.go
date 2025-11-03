// Rate Limiter Design
//
// Design a simple rate limiter that allows a maximum of N requests per time window. Implement
// a function that takes a request timestamp and returns whether the request should be allowed
// or denied based on the rate limit.
//
// Difficulty: Easy
//
// Problem:
// Implement a rate limiter using the sliding window log algorithm. Given a sequence of request
// timestamps and rate limit parameters (maxRequests, windowSizeSeconds), determine which
// requests should be allowed.
//
// Example 1:
// Input: requests = [1, 2, 3, 4, 5], maxRequests = 3, windowSize = 3
// Output: [true, true, true, false, false]
// Explanation: First 3 requests allowed, requests 4,5 exceed limit within window
//
// Example 2:
// Input: requests = [1, 5, 9, 10], maxRequests = 2, windowSize = 5
// Output: [true, true, true, false]
// Explanation: Request at t=9 starts new window. Request at t=10 is 2nd in window [5,10]
//
// Example 3:
// Input: requests = [1, 10, 20], maxRequests = 1, windowSize = 5
// Output: [true, true, true]
// Explanation: All requests are more than 5 seconds apart
//
// Constraints:
// - 1 <= requests.length <= 1000
// - 0 <= requests[i] <= 1000000 (timestamps in seconds)
// - Requests are given in chronological order (sorted)
// - 1 <= maxRequests <= 1000
// - 1 <= windowSize <= 3600
