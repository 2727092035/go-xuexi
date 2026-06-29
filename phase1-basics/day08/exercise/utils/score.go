package utils

import "fmt"

// ValidateScore 是导出函数。
// 练习：把函数名改成 validateScore，再运行 go run .，观察 parser.go 为什么无法调用。
func ValidateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("score %d out of range 0-100", score)
	}
	return nil
}
