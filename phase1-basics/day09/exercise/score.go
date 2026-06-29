package main

import "errors"

// TODO: 实现 ParseScore。
// 要求：TrimSpace、Atoi、ValidateScore；非法输入返回 error。
func ParseScore(raw string) (int, error) {
	return 0, errors.New("TODO: implement ParseScore")
}

// TODO: 实现 ValidateScore。
// 要求：分数必须在 0 到 100 之间。
func ValidateScore(score int) error {
	return errors.New("TODO: implement ValidateScore")
}

// TODO: 实现 Grade。
// 规则：>=90 A，>=80 B，>=70 C，>=60 D，否则 F。
func Grade(score float64) string {
	return "F"
}

// TODO: 实现 MaxAndCount。
// 要求：返回切片最大值，以及最大值出现次数；空切片返回 0, 0。
func MaxAndCount(nums []int) (int, int) {
	return 0, 0
}
