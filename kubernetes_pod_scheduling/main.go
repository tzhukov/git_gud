// Kubernetes Pod Scheduling
//
// You are implementing a simplified Kubernetes scheduler. Given pods with resource requirements,
// node affinities, taints/tolerations, and available nodes, determine the optimal pod-to-node
// assignment that maximizes resource utilization while respecting all constraints.
//
// Difficulty: Hard
//
// Problem:
// Schedule pods to nodes considering: resource limits (CPU/memory), node affinity rules, taints
// and tolerations. Return a valid assignment or indicate if scheduling is impossible.
//
// Each pod has: {id, cpu, memory, affinity: [nodeLabels], tolerations: [taints]}
// Each node has: {id, cpu, memory, labels: [strings], taints: [strings]}
//
// Example 1:
// Input: pods = [{id: "p1", cpu: 2, memory: 4, affinity: ["ssd"], tolerations: []},
//                {id: "p2", cpu: 1, memory: 2, affinity: [], tolerations: ["spot"]}],
//        nodes = [{id: "n1", cpu: 5, memory: 10, labels: ["ssd"], taints: []},
//                 {id: "n2", cpu: 3, memory: 5, labels: [], taints: ["spot"]}]
// Output: {"p1": "n1", "p2": "n2"}
// Explanation: p1 needs ssd (n1), p2 tolerates spot (n2)
//
// Example 2:
// Input: pods = [{id: "p1", cpu: 10, memory: 20, affinity: [], tolerations: []}],
//        nodes = [{id: "n1", cpu: 5, memory: 10, labels: [], taints: []}]
// Output: null
// Explanation: Pod resources exceed node capacity, scheduling impossible
//
// Example 3:
// Input: pods = [{id: "p1", cpu: 1, memory: 1, affinity: ["gpu"], tolerations: []}],
//        nodes = [{id: "n1", cpu: 10, memory: 10, labels: ["cpu"], taints: []}]
// Output: null
// Explanation: No node matches affinity requirement
//
// Constraints:
// - 1 <= pods <= 100
// - 1 <= nodes <= 50
// - 1 <= cpu <= 100, 1 <= memory <= 1000
// - Pod cannot be placed on node with taints unless pod has matching toleration
// - If pod has affinity, node must have all specified labels
// - Maximize number of successfully scheduled pods
