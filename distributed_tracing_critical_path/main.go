// Distributed Tracing Critical Path
//
// Analyze distributed traces to find the critical path (longest path) through a microservices
// call tree. Given trace spans with parent-child relationships, start/end times, and service
// names, identify the sequence of spans that contributes most to total request latency.
//
// Difficulty: Hard
//
// Problem:
// Given trace spans where each span has {id, parentId, service, startTime, endTime}, find the
// critical path - the sequence of spans from root to leaf with maximum cumulative duration.
// Return the path as an ordered list of span IDs and the total duration.
//
// Example 1:
// Input: spans = [
//   {id: "s1", parentId: null, service: "api", start: 0, end: 100},
//   {id: "s2", parentId: "s1", service: "auth", start: 10, end: 30},
//   {id: "s3", parentId: "s1", service: "db", start: 35, end: 95}
// ]
// Output: {path: ["s1", "s3"], duration: 160}
// Explanation: s1(100) + s3(60) = 160ms is longer than s1(100) + s2(20) = 120ms
//
// Example 2:
// Input: spans = [
//   {id: "s1", parentId: null, service: "gateway", start: 0, end: 50},
//   {id: "s2", parentId: "s1", service: "svc-a", start: 5, end: 25},
//   {id: "s3", parentId: "s2", service: "cache", start: 10, end: 20}
// ]
// Output: {path: ["s1", "s2", "s3"], duration: 80}
// Explanation: Only one path: s1(50) + s2(20) + s3(10) = 80ms
//
// Example 3:
// Input: spans = [{id: "s1", parentId: null, service: "api", start: 0, end: 100}]
// Output: {path: ["s1"], duration: 100}
// Explanation: Single span is the critical path
//
// Constraints:
// - 1 <= spans <= 1000
// - Each span has unique id
// - parentId is null for root spans, otherwise references another span
// - 0 <= startTime < endTime <= 1000000 (microseconds)
// - Child spans must be within parent's time range
// - The trace forms a valid tree structure (no cycles)
// - Multiple root spans possible (calculate critical path for each tree)
