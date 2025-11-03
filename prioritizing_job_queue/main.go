// Prioritizing Job Queue
//
// You are implementing a job scheduler that processes jobs based on priority and arrival time.
// Each job has a priority level (1-10, where 10 is highest) and an arrival timestamp. Process
// jobs in order of priority, and for jobs with equal priority, process them in FIFO order.
//
// Difficulty: Easy
//
// Problem:
// Given a list of jobs where each job is [jobId, priority, arrivalTime], return the order in
// which jobs should be processed.
//
// Example 1:
// Input: jobs = [[1, 5, 100], [2, 10, 101], [3, 5, 102]]
// Output: [2, 1, 3]
// Explanation: Job 2 has highest priority. Jobs 1 and 3 have equal priority, process by arrival
//
// Example 2:
// Input: jobs = [[1, 3, 50], [2, 3, 60], [3, 7, 55]]
// Output: [3, 1, 2]
// Explanation: Job 3 has priority 7. Jobs 1,2 have priority 3, ordered by arrival time
//
// Example 3:
// Input: jobs = [[1, 10, 100]]
// Output: [1]
// Explanation: Only one job to process
//
// Constraints:
// - 1 <= jobs.length <= 1000
// - jobs[i].length == 3
// - 1 <= jobId <= 10000 (unique job identifiers)
// - 1 <= priority <= 10
// - 0 <= arrivalTime <= 1000000 (timestamp in seconds)
// - All jobIds are unique
