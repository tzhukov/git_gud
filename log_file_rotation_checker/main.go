// Log File Rotation Checker
//
// You manage a system with rotating log files named with timestamps (e.g., app.log.20230101).
// Given a list of log file names and a retention period in days, determine which log files
// should be deleted to comply with the retention policy.
//
// Difficulty: Easy
//
// Problem:
// Given an array of log filenames with format "prefix.log.YYYYMMDD" and a retention period,
// return the filenames that should be deleted because they are older than the retention period
// from the current date.
//
// Example 1:
// Input: files = ["app.log.20230101", "app.log.20230115", "app.log.20230120"],
//        currentDate = "20230125", retentionDays = 7
// Output: ["app.log.20230101", "app.log.20230115"]
// Explanation: Files older than 7 days from 20230125 (before 20230118) should be deleted
//
// Example 2:
// Input: files = ["sys.log.20230601", "sys.log.20230610"], currentDate = "20230615",
//        retentionDays = 30
// Output: []
// Explanation: All files are within 30-day retention period
//
// Example 3:
// Input: files = ["error.log.20230101"], currentDate = "20230102", retentionDays = 1
// Output: ["error.log.20230101"]
// Explanation: File is exactly 1 day old, outside retention (exclusive)
//
// Constraints:
// - 1 <= files.length <= 1000
// - File format is always "prefix.log.YYYYMMDD"
// - 1 <= retentionDays <= 365
// - currentDate format is "YYYYMMDD"
// - All dates are valid (no need to validate date format)
// - Dates range from year 2020 to 2030
