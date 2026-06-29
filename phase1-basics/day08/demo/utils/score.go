package utils

import "fmt"

// ValidateScore 首字母大写，所以 demo/parser.go 可以通过 utils.ValidateScore 调用它。
func ValidateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("score %d out of range 0-100", score)
	}
	return nil
}
