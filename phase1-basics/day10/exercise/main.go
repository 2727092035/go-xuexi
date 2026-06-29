// Day 10 练习题
// 主题：基础阶段复盘、goroutine/channel/context 入门
// 运行：go run ./phase1-basics/day10/exercise

package main

import (
	"context"
	"fmt"
)

func main() {
	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
`

	students, err := ParseStudentsConcurrent(context.Background(), rawInput)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Println("解析结果:", students)
}
