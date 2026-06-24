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
	// strings 是 Go 标准库里的字符串工具包。
	// strings.TrimSpace(" 90 ") 会返回 "90"，作用是去掉字符串前后的空格、换行、Tab。
	// strings.Split("90,80,100", ",") 会按逗号拆分，返回 []string{"90", "80", "100"}。
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
	// raw 是调用方传进来的原始字符串，可能带空格，例如 " 88 "。
	// strings.TrimSpace(raw) 不会修改 raw 本身，而是返回一个去掉前后空白的新字符串。
	// 这里必须先清理空白，因为 strconv.Atoi(" 88 ") 会转换失败。
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	// strconv 是 Go 标准库里的字符串转换工具包。
	// Atoi 是 ASCII to integer 的缩写，可以把 "90" 转成 int 类型的 90。
	// 它返回两个值：score 是转换结果，err 是失败原因。
	// 只要 err != nil，就说明输入不是合法整数，不能继续使用 score。
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
	// strings.Split(input, ",") 会用逗号切开字符串。
	// 例如 input 是 "90, 80,100"，parts 就是 []string{"90", " 80", "100"}。
	// 注意第二段仍然带前导空格，所以后面还要交给 parseScore 再 TrimSpace。
	parts := strings.Split(input, ",")
	// make([]int, 0, len(parts)) 创建一个 int 切片。
	// 第 2 个参数 0 表示当前长度为 0，第 3 个参数 len(parts) 表示提前预留容量。
	// 后面每解析成功一个分数，就用 append 追加进去。
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
	// 约定一行学生数据是 "姓名,年龄,分数列表"。
	// 先用 strings.Split(line, ",") 按逗号拆成三段。
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return Student{}, fmt.Errorf("invalid student line %q, want name,age,scores", line)
	}

	// parts[0] 是姓名字段。用 TrimSpace 是为了允许 " Tom " 这种输入。
	name := strings.TrimSpace(parts[0])
	if name == "" {
		return Student{}, errors.New("student name is empty")
	}

	// parts[1] 是年龄字段。先 TrimSpace，再用 strconv.Atoi 转成 int。
	ageText := strings.TrimSpace(parts[1])
	age, err := strconv.Atoi(ageText)
	if err != nil {
		return Student{}, fmt.Errorf("invalid age %q", parts[1])
	}
	if age < 0 {
		return Student{}, fmt.Errorf("age %d out of range", age)
	}

	// parts[2] 是分数字段，格式是 "90|80|70"。
	// 这里用竖线 | 再拆一次，得到每个单独的分数字符串。
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
