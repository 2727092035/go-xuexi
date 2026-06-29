package main

import (
	"errors"
	"strconv"
	"strings"

	"go-learn/phase1-basics/day08/exercise/utils"
)

// TODO: 实现 parseStudents。
// 要求：TrimSpace、按换行拆分、逐行调用 parseStudent，错误里带行号。
func parseStudents(input string) ([]Student, error) {
	return nil, errors.New("TODO: implement parseStudents")
}

// TODO: 实现 parseStudent。
// 输入格式："Tom,90|80|70"。
func parseStudent(line string) (Student, error) {
	return Student{}, errors.New("TODO: implement parseStudent")
}

// TODO: 实现 parseScores。
// 要求：按 | 拆分，循环调用 parseScore。
func parseScores(input string) ([]int, error) {
	return nil, errors.New("TODO: implement parseScores")
}

// TODO: 实现 parseScore。
// 要求：TrimSpace、strconv.Atoi、utils.ValidateScore。
func parseScore(raw string) (int, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	score, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}

	if err := utils.ValidateScore(score); err != nil {
		return 0, err
	}
	return score, nil
}
