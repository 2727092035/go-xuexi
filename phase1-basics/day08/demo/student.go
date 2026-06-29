package main

// Student 是导出的结构体。
// 首字母大写的标识符可以被其他 package 使用。
type Student struct {
	Name   string
	Scores []int
}

// Average 是导出方法，用来计算单个学生的平均分。
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
