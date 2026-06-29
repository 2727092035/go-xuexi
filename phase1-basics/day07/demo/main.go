// Day 7 示例：成绩统计 CLI v1
// 运行：go run phase1-basics/day07/demo/main.go

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Student 是学生结构体，用来把“姓名 + 多个分数”放在同一个值里。
// type Student struct 是 Day 5 学过的结构体语法。
// Name 字段保存姓名，Scores 字段保存多个分数，所以类型是 []int。
type Student struct {
	Name   string
	Scores []int
}

// ClassStats 是班级统计结果。
// 把统计结果做成结构体，是为了让 calculateStats 只负责计算，
// printReport 只负责展示，main 只负责组织流程。
type ClassStats struct {
	Average float64
	Highest int
	Lowest  int
}

func main() {
	fmt.Println("===== 1. 从多行字符串解析学生列表 =====")
	// rawInput 模拟用户输入或文件内容。
	// 每一行是一个学生：姓名,分数列表。
	// 多个分数用竖线 | 分隔，例如 Tom,90|80|70。
	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
Jack,59
`

	students, err := parseStudents(rawInput)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Printf("解析结果: %+v\n", students)

	fmt.Println("\n===== 2. 统计班级平均分、最高分、最低分 =====")
	stats, err := calculateStats(students)
	if err != nil {
		fmt.Println("统计失败:", err)
		return
	}
	fmt.Printf("统计结果: %+v\n", stats)

	fmt.Println("\n===== 3. 用 map 统计等级人数 =====")
	counts := gradeCounts(students)
	fmt.Printf("等级人数: %+v\n", counts)

	fmt.Println("\n===== 4. 输出完整报表 =====")
	printReport(students, stats, counts)

	fmt.Println("\n===== 5. 错误输入演示 =====")
	// badInput 里 Bob 的分数是 abc，parseScore 会返回错误。
	// main 不负责修复错误，只负责把错误展示出来。
	badInput := "Tom,90\nBob,abc"
	_, err = parseStudents(badInput)
	fmt.Println("错误输入:", err)
}

func parseStudents(input string) ([]Student, error) {
	// strings.TrimSpace(input) 会去掉整段输入前后的空白和换行。
	// 这里必须先清理，因为多行字符串开头和结尾通常会带换行。
	text := strings.TrimSpace(input)
	if text == "" {
		return nil, errors.New("students input is empty")
	}

	// strings.Split(text, "\n") 按换行符拆分。
	// "\n" 是换行符字符串，表示每遇到一行结束就切开。
	// lines 的类型是 []string，里面每个元素都是一行学生数据。
	lines := strings.Split(text, "\n")

	// make([]Student, 0, len(lines)) 创建 Student 切片。
	// 当前长度是 0，容量预留为行数，后面用 append 逐个加入解析成功的学生。
	students := make([]Student, 0, len(lines))
	for index, line := range lines {
		student, err := parseStudent(line)
		if err != nil {
			// fmt.Errorf 里的 %w 会包装原始错误。
			// 这里加上 index+1，是为了告诉调用方第几行出错。
			return nil, fmt.Errorf("line %d invalid: %w", index+1, err)
		}
		students = append(students, student)
	}

	return students, nil
}

func parseStudent(line string) (Student, error) {
	// 一行学生数据格式是 "姓名,分数列表"。
	// strings.Split(line, ",") 用逗号拆成两段：姓名和分数列表。
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return Student{}, fmt.Errorf("invalid student line %q, want name,scores", line)
	}

	// strings.TrimSpace(parts[0]) 去掉姓名前后的空白。
	// 例如 " Tom " 会变成 "Tom"。
	name := strings.TrimSpace(parts[0])
	if name == "" {
		return Student{}, errors.New("student name is empty")
	}

	scores, err := parseScores(parts[1])
	if err != nil {
		return Student{}, err
	}

	return Student{Name: name, Scores: scores}, nil
}

func parseScores(input string) ([]int, error) {
	// 分数列表格式是 "90|80|70"。
	// strings.Split(input, "|") 用竖线切开，得到 []string{"90", "80", "70"}。
	parts := strings.Split(input, "|")
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

func parseScore(raw string) (int, error) {
	// raw 是原始分数字符串，可能带空格，例如 " 90 "。
	// TrimSpace 返回去掉前后空白的新字符串，不会修改 raw 本身。
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	// strconv.Atoi(text) 把字符串转成 int。
	// 它返回两个值：转换结果 score，以及失败原因 err。
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

func calculateStats(students []Student) (ClassStats, error) {
	// len(students) 返回切片里有多少个学生。
	// 如果没有学生，后面无法计算平均分、最高分、最低分，所以直接返回错误。
	if len(students) == 0 {
		return ClassStats{}, errors.New("students is empty")
	}

	total := 0
	count := 0
	highest := 0
	lowest := 0
	hasScore := false

	for _, student := range students {
		for _, score := range student.Scores {
			if !hasScore {
				highest = score
				lowest = score
				hasScore = true
			}
			if score > highest {
				highest = score
			}
			if score < lowest {
				lowest = score
			}
			total += score
			count++
		}
	}

	if count == 0 {
		return ClassStats{}, errors.New("scores is empty")
	}

	// float64(total) 是类型转换，把 int 转成 float64。
	// 如果直接用 total / count，会做整数除法，小数部分会丢失。
	average := float64(total) / float64(count)
	return ClassStats{
		Average: average,
		Highest: highest,
		Lowest:  lowest,
	}, nil
}

func averageScore(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func grade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func gradeCounts(students []Student) map[string]int {
	// map[string]int 表示 key 是字符串、value 是整数。
	// 这里 key 是等级 "A"、"B"、"C"、"D"、"F"，value 是对应等级的人数。
	counts := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
	for _, student := range students {
		level := grade(averageScore(student.Scores))
		counts[level]++
	}
	return counts
}

func printReport(students []Student, stats ClassStats, counts map[string]int) {
	fmt.Println("学生成绩：")
	for _, student := range students {
		avg := averageScore(student.Scores)
		fmt.Printf("- %s scores=%v average=%.1f grade=%s\n", student.Name, student.Scores, avg, grade(avg))
	}

	fmt.Printf("班级平均分: %.1f\n", stats.Average)
	fmt.Println("最高分:", stats.Highest)
	fmt.Println("最低分:", stats.Lowest)
	fmt.Printf("等级分布: A=%d B=%d C=%d D=%d F=%d\n",
		counts["A"], counts["B"], counts["C"], counts["D"], counts["F"])
}
