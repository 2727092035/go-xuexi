package main

import "fmt"

func PrintReport(students []Student) {
	stats, err := CalculateStats(students)
	if err != nil {
		fmt.Println("统计失败:", err)
		return
	}
	counts := GradeCounts(students)

	for _, student := range students {
		avg := student.Average()
		fmt.Printf("- %s scores=%v average=%.1f grade=%s\n", student.Name, student.Scores, avg, Grade(avg))
	}
	fmt.Printf("班级平均分: %.1f 最高分: %d 最低分: %d\n", stats.Average, stats.Highest, stats.Lowest)
	fmt.Printf("等级分布: A=%d B=%d C=%d D=%d F=%d\n",
		counts["A"], counts["B"], counts["C"], counts["D"], counts["F"])
}
