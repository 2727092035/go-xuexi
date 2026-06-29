// Day 10 示例：基础阶段复盘 + goroutine/channel/context 入门
// 运行：go run ./phase1-basics/day10/demo

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("===== Day 10：并发解析成绩数据 =====")

	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
Jack,59
`

	fmt.Println("顺序解析：")
	students, err := ParseStudents(rawInput)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	PrintReport(students)

	fmt.Println("\n并发解析：")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	students, err = ParseStudentsConcurrent(ctx, rawInput)
	if err != nil {
		fmt.Println("并发解析失败:", err)
		return
	}
	PrintReport(students)
}
