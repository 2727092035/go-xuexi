// Day 6 示例：错误处理、输入校验、strings/strconv
// 运行：go run phase1-basics/day06/demo/main.go

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
	fmt.Println("===== 1. error 的基本写法 =====")
	// Go 通常不用 try/catch。函数如果可能失败，就把 error 作为最后一个返回值。
	// 调用方通过 err == nil 判断是否成功。
	result, err := divide(10, 2)
	fmt.Println("10 / 2:", result, "error:", err)

	result, err = divide(10, 0)
	fmt.Println("10 / 0:", result, "error:", err)

	fmt.Println("\n===== 2. strconv.Atoi：字符串转整数 =====")
	// strconv.Atoi 会返回两个值：转换后的 int 和 error。
	// 第二个返回值必须检查，否则遇到非法数字时会丢失失败原因。
	showParseScore("90")
	showParseScore(" 88 ")
	showParseScore("abc")

	fmt.Println("\n===== 3. strings：清理和拆分字符串 =====")
	// strings.TrimSpace 用来去掉前后空白。
	// strings.Split 用分隔符把一个字符串拆成 []string。
	input := "90, 80,100"
	scores, err := parseScores(input)
	fmt.Println("输入:", input)
	fmt.Println("分数:", scores, "error:", err)

	badInput := "90,abc,100"
	scores, err = parseScores(badInput)
	fmt.Println("输入:", badInput)
	fmt.Println("分数:", scores, "error:", err)

	fmt.Println("\n===== 4. 连续校验：失败就提前返回 =====")
	// parseStudent 会依次做：拆字段、清理空格、转年龄、转分数、校验范围。
	// 任何一步失败，都立刻返回清晰的 error。
	line := "Tom,18,90|80|70"
	student, err := parseStudent(line)
	fmt.Printf("输入: %q\n学生: %+v\nerror: %v\n", line, student, err)

	line = "Tom,18,90|abc|70"
	student, err = parseStudent(line)
	fmt.Printf("输入: %q\n学生: %+v\nerror: %v\n", line, student, err)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func showParseScore(raw string) {
	score, err := parseScore(raw)
	fmt.Printf("parseScore(%q) => score=%d err=%v\n", raw, score, err)
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
