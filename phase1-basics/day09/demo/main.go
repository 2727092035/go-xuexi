// Day 9 示例：testing、表格驱动测试、benchmark 入门
// 运行 demo：go run ./phase1-basics/day09/demo
// 运行测试：go test ./phase1-basics/day09/demo
// 运行 benchmark：go test -bench=. ./phase1-basics/day09/demo

package main

import "fmt"

func main() {
	fmt.Println("===== Day 9：测试入门 =====")

	for _, raw := range []string{"90", "abc", "-1", "101"} {
		score, err := ParseScore(raw)
		fmt.Printf("ParseScore(%q) => score=%d err=%v\n", raw, score, err)
	}

	for _, score := range []float64{95, 83, 72, 61, 30} {
		fmt.Printf("Grade(%.1f) => %s\n", score, Grade(score))
	}

	max, count := MaxAndCount([]int{90, 100, 80, 100})
	fmt.Printf("MaxAndCount => max=%d count=%d\n", max, count)

	fmt.Println("\n现在运行：go test ./phase1-basics/day09/demo")
}
