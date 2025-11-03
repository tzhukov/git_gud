// Circuit Breaker State Machine
//
// Implement a circuit breaker pattern that transitions between states (CLOSED, OPEN, HALF_OPEN)
// based on failure rates. Track request outcomes and determine when to open/close the circuit
// to protect downstream services from cascading failures.
//
// Difficulty: Medium
//
// Problem:
// Given a sequence of request results (success=1, failure=0), failure threshold, and timeout
// period, simulate the circuit breaker and return the state after each request.
//
// States:
// - CLOSED: Requests flow normally. Opens if failures exceed threshold
// - OPEN: All requests fail fast. Transitions to HALF_OPEN after timeout
// - HALF_OPEN: Test request allowed. Success -> CLOSED, Failure -> OPEN
//
// Example 1:
// Input: requests = [1,1,0,0,0,1], failureThreshold = 3, windowSize = 5
// Output: ["CLOSED", "CLOSED", "CLOSED", "CLOSED", "OPEN", "OPEN"]
// Explanation: After 3 failures in window, circuit opens
//
// Example 2:
// Input: requests = [0,0,0,1,1,1], failureThreshold = 2, windowSize = 3
// Output: ["CLOSED", "CLOSED", "OPEN", "OPEN", "OPEN", "OPEN"]
// Explanation: Circuit opens after 2 failures, stays open
//
// Example 3:
// Input: requests = [1,1,1,1,1], failureThreshold = 3, windowSize = 5
// Output: ["CLOSED", "CLOSED", "CLOSED", "CLOSED", "CLOSED"]
// Explanation: No failures, circuit remains closed
//
// Constraints:
// - 1 <= requests.length <= 1000
// - requests[i] is either 0 (failure) or 1 (success)
// - 1 <= failureThreshold <= windowSize
// - 1 <= windowSize <= 100
// - For simplicity, ignore timeout transitions (focus on CLOSED/OPEN)
