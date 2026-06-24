// Day 6 练习题
// 主题：error、输入校验、strings、strconv
// 运行：go run phase1-basics/day06/exercise/main.go

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Age    int
	Scores []int
}

func main() {
	for _, raw := range []string{"90", " 80 ", "abc", "-1", "101", ""} {
		score, err := parseScore(raw)
		fmt.Printf("parseScore %s: %d, %v\n", raw, score, err)
	}

	for i, score := range []int{90, 80, 70, 60, 50, 40, 30, 20, 10, 0, -1, 101} {
		err := validateScore(score)
		fmt.Printf("validateScore 下标 %d 分数 %d: %v\n", i, score, err)
	}

	for _, raw := range []string{"90,80,70", "90, abc, 70", "100,0", ""} {
		scores, err := parseScores(raw)
		fmt.Printf("字符串转换为分数数组 %s: %v, %v\n", raw, scores, err)
	}

	fmt.Println("--------------------------------")

	for _, item := range [][2]int{{10, 2}, {10, 0}, {9, 4}} {
		result, err := divide(item[0], item[1])
		fmt.Printf("divide %d, %d: %d, %v\n", item[0], item[1], result, err)
	}

	for _, raw := range []string{
		"Tom,18,90|80|70",
		"Jerry,20,100|95",
		"Lucy,x,90|80",
		"Bob,19,90|abc",
		"Diana,18,90",
	} {
		student, err := parseStudent(raw)
		fmt.Printf("解析学生信息 %s: %v, %v\n", raw, student, err)
	}

	fmt.Println("--------------------------------")
	students, err := parseStudents("Tom,18,90|80|70;Jerry,20,100|95")
	fmt.Printf("批量解析学生信息: %v, %v\n", students, err)

	students, err = parseStudents("Tom,18,90|80|70;Bad,x,90")
	fmt.Printf("批量解析学生信息: %v, %v\n", students, err)
}

func parseScore(raw string) (int, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return 0, errors.New("empty string")
	}

	score, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid score: %w", err)
	}

	if err := validateScore(score); err != nil {
		return 0, err
	}

	return score, nil
}

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return errors.New("score out of range")
	}
	return nil
}

func parseScores(input string) ([]int, error) {
	text := strings.TrimSpace(input)
	if text == "" {
		return nil, errors.New("empty string")
	}

	parts := strings.Split(text, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		score, err := parseScore(part)
		if err != nil {
			return nil, err
		}
		result = append(result, score)
	}
	return result, nil
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// parseStudent 解析一行学生信息，格式为 "name,age,score1|score2|score3"。
func parseStudent(input string) (Student, error) {
	fields := strings.Split(input, ",")
	if len(fields) != 3 {
		return Student{}, errors.New("invalid input")
	}

	name := strings.TrimSpace(fields[0])
	if name == "" {
		return Student{}, errors.New("name is required")
	}

	age, err := strconv.Atoi(strings.TrimSpace(fields[1]))
	if err != nil {
		return Student{}, fmt.Errorf("invalid age: %w", err)
	}
	if age < 0 {
		return Student{}, errors.New("age out of range")
	}

	scoreParts := strings.Split(fields[2], "|")
	result := make([]int, 0, len(scoreParts))
	for _, part := range scoreParts {
		score, err := parseScore(part)
		if err != nil {
			return Student{}, err
		}
		result = append(result, score)
	}

	return Student{Name: name, Age: age, Scores: result}, nil
}

// parseStudents 批量解析学生信息，多行之间用分号分隔。
func parseStudents(input string) ([]Student, error) {
	text := strings.TrimSpace(input)
	if text == "" {
		return nil, errors.New("empty string")
	}

	lines := strings.Split(text, ";")
	students := make([]Student, 0, len(lines))
	for index, line := range lines {
		student, err := parseStudent(line)
		if err != nil {
			return nil, fmt.Errorf("student line %d invalid: %w", index+1, err)
		}
		students = append(students, student)
	}
	return students, nil
}
