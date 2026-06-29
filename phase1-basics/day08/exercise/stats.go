package main

import "errors"

type ClassStats struct {
	Average float64
	Highest int
	Lowest  int
}

// TODO: 实现 calculateStats。
// 要求：统计所有分数的平均分、最高分、最低分；无学生或无分数时返回 error。
func calculateStats(students []Student) (ClassStats, error) {
	return ClassStats{}, errors.New("TODO: implement calculateStats")
}

// TODO: 实现 grade。
// 规则：>=90 A，>=80 B，>=70 C，>=60 D，否则 F。
func grade(score float64) string {
	return "F"
}

// TODO: 实现 gradeCounts。
// 要求：用 map[string]int 统计 A/B/C/D/F 各多少人。
func gradeCounts(students []Student) map[string]int {
	return map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
}
