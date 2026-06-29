// Day 8 练习题
// 主题：包、文件拆分、go run .、导出标识符
// 运行：
// 1. cd phase1-basics/day08/exercise
// 2. go run .
//
// 今天不要只运行 main.go。只运行 main.go 时，Go 不会自动带上同目录其他文件。

package main

import "fmt"

func main() {
	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
`

	students, err := parseStudents(rawInput)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	stats, err := calculateStats(students)
	if err != nil {
		fmt.Println("统计失败:", err)
		return
	}

	counts := gradeCounts(students)
	printReport(students, stats, counts)
}
