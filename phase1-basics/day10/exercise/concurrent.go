package main

import (
	"context"
	"errors"
)

// TODO: 实现 ParseStudentsConcurrent。
// 要求：
// 1. 用 splitStudentLines 拆出每一行。
// 2. 创建 jobs channel 和 results channel。
// 3. 启动 2 个 goroutine 并发调用 ParseStudent。
// 4. 用 context 支持取消。
// 5. 返回结果顺序要和输入行顺序一致。
func ParseStudentsConcurrent(ctx context.Context, input string) ([]Student, error) {
	return nil, errors.New("TODO: implement ParseStudentsConcurrent")
}
