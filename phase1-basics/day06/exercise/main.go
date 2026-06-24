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
	fmt.Println("练习 1：parseScore")
	// 要求：
	// 1. 去掉输入字符串前后空白。
	// 2. 把字符串转成 int。
	// 3. 非数字返回 error。
	// 4. 分数必须在 0 到 100。
	for _, raw := range []string{"90", " 80 ", "abc", "-1", "101", ""} {
		score, err := parseScore(raw)
		fmt.Printf("parseScore(%q) => score=%d err=%v\n", raw, score, err)
	}

	fmt.Println("\n练习 2：validateScore")
	// 要求：
	// 1. score < 0 返回错误。
	// 2. score > 100 返回错误。
	// 3. 0 到 100 返回 nil。
	for _, score := range []int{-1, 0, 60, 100, 101} {
		fmt.Printf("validateScore(%d) => %v\n", score, validateScore(score))
	}

	fmt.Println("\n练习 3：parseScores")
	// 要求：
	// 1. 输入格式："90,80,70"。
	// 2. 使用 strings.Split 拆分。
	// 3. 每个分数复用 parseScore。
	// 4. 任意一个分数非法，就返回 error。
	for _, input := range []string{"90,80,70", "90, abc, 70", "100,0", ""} {
		scores, err := parseScores(input)
		fmt.Printf("parseScores(%q) => scores=%v err=%v\n", input, scores, err)
	}

	fmt.Println("\n加餐 1：divide")
	// 要求：
	// 1. 正常除法返回结果和 nil。
	// 2. 除数为 0 时返回 error。
	for _, item := range [][2]int{{10, 2}, {10, 0}, {9, 4}} {
		result, err := divide(item[0], item[1])
		fmt.Printf("divide(%d, %d) => result=%d err=%v\n", item[0], item[1], result, err)
	}

	fmt.Println("\n加餐 2：parseStudent")
	// 要求：
	// 1. 输入格式："Tom,18,90|80|70"。
	// 2. 解析出 Student。
	// 3. 年龄必须是合法整数且不能小于 0。
	// 4. 分数复用 parseScore。
	for _, line := range []string{
		"Tom,18,90|80|70",
		"Jerry,20,100|95",
		"Lucy,x,90|80",
		"Bob,19,90|abc",
		",18,90",
	} {
		student, err := parseStudent(line)
		fmt.Printf("parseStudent(%q) => student=%+v err=%v\n", line, student, err)
	}

	fmt.Println("\n加餐 3：统计学生平均分和等级")
	// 要求：
	// 1. 给 Student 实现 Average() float64。
	// 2. 给 Student 实现 Level() string。
	// 3. 复用 parseStudent 得到的数据。
	student, err := parseStudent("Alice,18,90|80|70")
	if err != nil {
		fmt.Println("解析学生失败:", err)
		return
	}
	fmt.Printf("%s average=%.1f level=%s\n", student.Name, student.Average(), student.Level())

	fmt.Println("\n练习完成！")
}

func parseScore(raw string) (int, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	score, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("invalid score %q", raw)
	}

	if err := validateScore(score); err != nil {
		return 0, err
	}

	return score, nil
}

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("score %d out of range 0-100", score)
	}
	return nil
}

func parseScores(input string) ([]int, error) {
	if strings.TrimSpace(input) == "" {
		return nil, errors.New("scores input is empty")
	}

	parts := strings.Split(input, ",")
	scores := make([]int, 0, len(parts))
	for _, part := range parts {
		score, err := parseScore(part)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func parseStudent(line string) (Student, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return Student{}, fmt.Errorf("invalid student line %q, want name,age,scores", line)
	}

	name := strings.TrimSpace(parts[0])
	if name == "" {
		return Student{}, errors.New("student name is empty")
	}

	ageText := strings.TrimSpace(parts[1])
	age, err := strconv.Atoi(ageText)
	if err != nil {
		return Student{}, fmt.Errorf("invalid age %q", parts[1])
	}
	if age < 0 {
		return Student{}, fmt.Errorf("age %d out of range", age)
	}

	scoreParts := strings.Split(parts[2], "|")
	scores := make([]int, 0, len(scoreParts))
	for _, part := range scoreParts {
		score, err := parseScore(part)
		if err != nil {
			return Student{}, err
		}
		scores = append(scores, score)
	}

	return Student{
		Name:   name,
		Age:    age,
		Scores: scores,
	}, nil
}

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

func (s Student) Level() string {
	average := s.Average()
	switch {
	case average >= 90:
		return "A"
	case average >= 80:
		return "B"
	case average >= 70:
		return "C"
	case average >= 60:
		return "D"
	default:
		return "F"
	}
}
