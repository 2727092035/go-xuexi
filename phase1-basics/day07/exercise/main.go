// Day 7 练习题
// 主题：成绩统计 CLI v1、组合函数、slice、map、struct、error
// 运行：go run phase1-basics/day07/exercise/main.go

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Scores []int
}

type ClassStats struct {
	Average float64
	Highest int
	Lowest  int
}

func main() {
	fmt.Println("练习 1：parseStudent")
	// 要求：
	// 1. 实现 parseStudent(line string) (Student, error)。
	// 2. 输入格式："Tom,90|80|70"。
	// 3. 用 strings.Split(line, ",") 拆出姓名和分数列表。
	// 4. 姓名不能为空。
	// 5. 分数列表复用 parseScores。
	// 在这里写代码 👇
	for _, line := range []string{"Tom,90|80|70", " Jerry ,78|82", ",90", "Bob,abc"} {
		student, err := parseStudent(line)
		fmt.Printf("parseStudent(%q) => student=%+v err=%v\n", line, student, err)
	}

	fmt.Println("\n练习 2：parseStudents")
	// 要求：
	// 1. 实现 parseStudents(input string) ([]Student, error)。
	// 2. 输入是多行字符串，每行一个学生。
	// 3. 用 strings.TrimSpace 清理整段输入。
	// 4. 用 strings.Split(input, "\n") 拆成多行。
	// 5. 每一行复用 parseStudent。
	// 6. 某一行出错时，错误信息里要带行号。
	// 在这里写代码 👇
	rawInput := `
Tom,90|80|70
Jerry,78|82
Lucy,100|96
`
	students, err := parseStudents(rawInput)
	fmt.Printf("parseStudents(valid) => students=%+v err=%v\n", students, err)

	badInput := "Tom,90\nBob,abc"
	students, err = parseStudents(badInput)
	fmt.Printf("parseStudents(invalid) => students=%+v err=%v\n", students, err)

	fmt.Println("\n练习 3：calculateStats")
	// 要求：
	// 1. 实现 calculateStats(students []Student) (ClassStats, error)。
	// 2. 统计所有分数的平均分、最高分、最低分。
	// 3. 空学生列表或没有任何分数时返回 error。
	// 4. 平均分要保留小数，注意 int 转 float64。
	// 在这里写代码 👇
	students = []Student{
		{Name: "Tom", Scores: []int{90, 80, 70}},
		{Name: "Jerry", Scores: []int{78, 82}},
		{Name: "Lucy", Scores: []int{100, 96}},
	}
	stats, err := calculateStats(students)
	fmt.Printf("calculateStats => stats=%+v err=%v\n", stats, err)

	fmt.Println("\n练习 4：grade 和 gradeCounts")
	// 要求：
	// 1. 实现 averageScore(scores []int) float64。
	// 2. 实现 grade(score float64) string。
	// 3. 实现 gradeCounts(students []Student) map[string]int。
	// 4. grade 规则：>=90 A，>=80 B，>=70 C，>=60 D，否则 F。
	// 在这里写代码 👇
	counts := gradeCounts(students)
	fmt.Printf("gradeCounts => %+v\n", counts)

	fmt.Println("\n加餐：printReport")
	// 要求：
	// 1. 实现 printReport(students []Student, stats ClassStats, counts map[string]int)。
	// 2. 输出每个学生的分数、平均分、等级。
	// 3. 输出班级平均分、最高分、最低分、等级分布。
	// 4. main 只负责组织流程，格式化输出放到 printReport。
	// 在这里写代码 👇
	printReport(students, stats, counts)

	fmt.Println("\n练习完成！")
}

// 在下面补全练习函数。
// 提示：完成这些函数时，需要按需把 errors、strconv、strings 加到 import 里。

func parseStudent(line string) (Student, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return Student{}, fmt.Errorf("invalid student line %q, want name,scores", line)
	}

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

func parseStudents(input string) ([]Student, error) {
	text := strings.TrimSpace(input)
	if text == "" {
		return nil, errors.New("students input is empty")
	}

	lines := strings.Split(text, "\n")
	students := make([]Student, 0, len(lines))
	for i, line := range lines {
		// 支持跳过空行，便于处理用户粘贴多行文本时夹带的空白行。
		if strings.TrimSpace(line) == "" {
			continue
		}

		student, err := parseStudent(line)
		if err != nil {
			return nil, fmt.Errorf("line %d invalid: %w", i+1, err)
		}
		students = append(students, student)
	}
	return students, nil
}

func parseScores(input string) ([]int, error) {
	if strings.TrimSpace(input) == "" {
		return nil, errors.New("scores is empty")
	}

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

func calculateStats(students []Student) (ClassStats, error) {
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
			total += score
			count++
			if score > highest {
				highest = score
			}
			if score < lowest {
				lowest = score
			}
		}
	}

	if count == 0 {
		return ClassStats{}, errors.New("scores is empty")
	}

	return ClassStats{Average: float64(total) / float64(count), Highest: highest, Lowest: lowest}, nil
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
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func gradeCounts(students []Student) map[string]int {
	counts := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
	for _, student := range students {
		avg := averageScore(student.Scores)
		counts[grade(avg)]++
	}
	return counts
}

func printReport(students []Student, stats ClassStats, counts map[string]int) {
	fmt.Println("学生 分数 平均分 等级")
	for _, student := range students {
		avg := averageScore(student.Scores)
		fmt.Printf("%s %v %.1f %s\n", student.Name, student.Scores, avg, grade(avg))
	}
	fmt.Printf("班级平均分 %.1f 最高分 %d 最低分 %d\n", stats.Average, stats.Highest, stats.Lowest)
	fmt.Println("等级分布", counts)
}
