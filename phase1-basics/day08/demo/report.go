package main

import "fmt"

func PrintReport(students []Student, stats ClassStats, counts map[string]int) {
	fmt.Println("学生成绩：")
	for _, student := range students {
		avg := student.Average()
		fmt.Printf("- %s scores=%v average=%.1f grade=%s\n", student.Name, student.Scores, avg, Grade(avg))
	}

	fmt.Printf("班级平均分: %.1f\n", stats.Average)
	fmt.Println("最高分:", stats.Highest)
	fmt.Println("最低分:", stats.Lowest)
	fmt.Printf("等级分布: A=%d B=%d C=%d D=%d F=%d\n",
		counts["A"], counts["B"], counts["C"], counts["D"], counts["F"])
}
