// Distributed Systems Graph Reachability
//
// Given a distributed system represented as a directed graph where nodes represent services and
// edges represent dependencies between services, determine if a target service is reachable from
// a source service. The graph may contain cycles, and you need to handle network partitions.
//
// Difficulty: Medium
//
// Problem:
// You are given a map of service dependencies where each key is a service name and the value is
// a list of services it depends on. Determine if there is a path from the source service to the
// target service.
//
// Example 1:
// Input: dependencies = {"api": ["db", "cache"], "db": ["storage"], "cache": ["storage"],
//        "storage": []}, source = "api", target = "storage"
// Output: true
// Explanation: api -> db -> storage and api -> cache -> storage are both valid paths
//
// Example 2:
// Input: dependencies = {"service_a": ["service_b"], "service_b": ["service_c"],
//        "service_c": []}, source = "service_a", target = "service_d"
// Output: false
// Explanation: service_d is not reachable from service_a
//
// Example 3:
// Input: dependencies = {"a": ["b"], "b": ["c"], "c": ["a"]}, source = "a", target = "b"
// Output: true
// Explanation: There is a cycle but b is still reachable from a
//
// Constraints:
// - 1 <= number of services <= 1000
// - 0 <= dependencies per service <= 100
// - Service names are non-empty strings with max length 50
// - The graph may contain cycles
// - Source and target services may or may not exist in the graph
