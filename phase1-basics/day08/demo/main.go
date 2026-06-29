// Day 8 示例：把 Day 7 成绩统计 CLI 拆成多个文件和一个 utils 包
// 运行：
// 1. cd phase1-basics/day08/demo
// 2. go run .
//
// 注意：今天要用 go run .，不要只运行 main.go。
// 因为 main.go 会调用同目录 student.go、parser.go、stats.go、report.go 里的函数。

package main

import "fmt"

func main() {
	fmt.Println("===== Day 8：多文件项目 =====")

	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
Jack,59
`

	students, err := parseStudents(rawInput)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	stats, err := CalculateStats(students)
	if err != nil {
		fmt.Println("统计失败:", err)
		return
	}

	counts := GradeCounts(students)
	PrintReport(students, stats, counts)

	fmt.Println("\n===== 导出标识符演示 =====")
	fmt.Println("Student、CalculateStats、PrintReport 首字母大写，所以可以被其他包使用。")
	fmt.Println("parseStudents 首字母小写，只能在 package main 内部使用。")
}
