// Container Resource Allocation
//
// You are scheduling containers on a cluster with limited CPU and memory. Each container
// requires specific CPU and memory resources. Given the available resources and container
// requirements, determine the maximum number of containers that can be scheduled.
//
// Difficulty: Easy
//
// Problem:
// Given arrays of CPU and memory requirements for containers, and total available CPU and
// memory, find the maximum number of containers you can run simultaneously.
//
// Example 1:
// Input: cpu = [2, 3, 1, 4], memory = [4, 2, 1, 3], availableCPU = 7, availableMemory = 8
// Output: 3
// Explanation: Can run containers [2,1,4] using CPU=7 and memory=[4,1,3]=8
//
// Example 2:
// Input: cpu = [5, 5, 5], memory = [10, 10, 10], availableCPU = 10, availableMemory = 20
// Output: 2
// Explanation: Can run 2 containers (CPU=10, memory=20)
//
// Example 3:
// Input: cpu = [1, 1, 1], memory = [1, 1, 1], availableCPU = 10, availableMemory = 10
// Output: 3
// Explanation: All containers fit within available resources
//
// Constraints:
// - 1 <= containers <= 1000
// - cpu.length == memory.length
// - 1 <= cpu[i] <= 100 (CPU cores)
// - 1 <= memory[i] <= 1000 (GB)
// - 1 <= availableCPU <= 10000
// - 1 <= availableMemory <= 100000
// - You cannot partially allocate a container
