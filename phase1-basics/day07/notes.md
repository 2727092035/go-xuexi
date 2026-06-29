# Day 7 笔记：成绩统计 CLI v1

## 今天先记住一句话

小项目不是把所有代码都写进 `main`，而是把流程拆成多个职责清楚的函数：解析输入、校验数据、统计结果、输出报表。

## 1. 今天要掌握什么

Day 7 会把前面学过的知识组合起来：

1. 用 `struct` 表达学生和统计结果。
2. 用 `[]Student` 保存多个学生。
3. 用 `strings.Split` 解析多行输入和分数字符串。
4. 用 `strconv.Atoi` 把字符串分数转成整数。
5. 用 `error` 把非法输入返回给调用方。
6. 用 `map[string]int` 统计 A/B/C/D/F 各等级人数。
7. 让 `main` 只负责组织流程，不堆解析和统计细节。

## 2. 数据格式

今天的输入格式是多行字符串：

```text
Tom,90|80|70
Jerry,78|82
Lucy,100|96
```

每一行代表一个学生：

- 逗号 `,` 左边是姓名。
- 逗号 `,` 右边是分数列表。
- 多个分数用竖线 `|` 分隔。

解析后的目标结构是：

```go
Student{
	Name:   "Tom",
	Scores: []int{90, 80, 70},
}
```

## 3. 新语法和 API 说明

### 多行字符串

demo 里有这样的写法：

```go
rawInput := `
Tom,90|80|70
Jerry,78|82
`
```

反引号包起来的是原始字符串字面量。它可以直接跨多行，里面的换行会被保留下来。它适合模拟文件内容、用户粘贴的一段文本、SQL 等多行内容。

### strings.Split(input, "\n")

`strings.Split` 来自标准库 `strings` 包。

```go
lines := strings.Split(text, "\n")
```

这行代码拆开看：

- `strings`：标准库包名，负责字符串处理。
- `Split`：函数名，按分隔符拆字符串。
- `text`：第一个参数，要被拆分的字符串。
- `"\n"`：第二个参数，换行符。
- `lines`：返回值，类型是 `[]string`，每个元素是一行文本。

为什么这里要用它？因为 `parseStudents` 收到的是一整段多行输入，必须先切成一行一行，才能复用 `parseStudent`。

### map[string]int

`map[string]int` 表示“字符串到整数”的映射：

```go
counts := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
```

这行代码拆开看：

- `map`：Go 的映射类型，类似 JS 对象或 `Map`。
- `string`：key 的类型，这里用等级 `"A"`、`"B"`。
- `int`：value 的类型，这里用人数。
- `counts["A"]`：读取 A 等级人数。
- `counts[level]++`：把某个等级的人数加 1。

为什么这里要用 map？因为等级是有限但不适合写 5 个变量，map 可以用等级字符串直接定位对应人数。

### float64(total) / float64(count)

Go 里 `int / int` 会做整数除法：

```go
fmt.Println(5 / 2) // 2
```

如果要得到小数，必须先转成 `float64`：

```go
average := float64(total) / float64(count)
```

这行代码拆开看：

- `float64(total)`：把整数总分转成浮点数。
- `float64(count)`：把整数数量转成浮点数。
- `/`：浮点数除法，会保留小数。

## 4. demo 代码逐段详解

### 第 1 段：parseStudents

`parseStudents(input string) ([]Student, error)` 负责把一整段多行输入解析成学生列表。

流程：

1. `strings.TrimSpace(input)` 清理整段输入前后的空白。
2. 空输入直接返回错误。
3. `strings.Split(text, "\n")` 拆成多行。
4. 循环每一行，调用 `parseStudent(line)`。
5. 如果某一行失败，用 `fmt.Errorf("line %d invalid: %w", index+1, err)` 返回带行号的错误。

这里的重点是复用：批量解析不应该重新写姓名和分数解析逻辑，而是交给 `parseStudent`。

### 第 2 段：parseStudent

`parseStudent(line string) (Student, error)` 负责解析一行：

```text
Tom,90|80|70
```

流程：

1. 用 `strings.Split(line, ",")` 按逗号拆成姓名和分数列表。
2. `len(parts) != 2` 时说明格式不对。
3. 用 `strings.TrimSpace(parts[0])` 清理姓名。
4. 姓名为空时返回错误。
5. 用 `parseScores(parts[1])` 解析多个分数。
6. 返回 `Student{Name: name, Scores: scores}`。

### 第 3 段：calculateStats

`calculateStats(students []Student) (ClassStats, error)` 负责统计所有学生的所有分数。

它会计算：

- 班级平均分 `Average`
- 最高分 `Highest`
- 最低分 `Lowest`

这里有一个变量 `hasScore`：

```go
hasScore := false
```

它的作用是判断是否已经遇到第一个分数。第一次遇到分数时，把最高分和最低分都设成这个分数。这样比一开始随便写 `highest := 0`、`lowest := 100` 更稳。

### 第 4 段：gradeCounts

`gradeCounts(students []Student) map[string]int` 负责统计等级分布。

流程：

1. 先创建 `counts := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}`。
2. 遍历每个学生。
3. 用 `averageScore(student.Scores)` 算学生平均分。
4. 用 `grade(avg)` 得到等级。
5. 用 `counts[level]++` 给对应等级人数加 1。

## 5. 和 JS/TS 的关键差异

- Go 没有自动类型转换，字符串分数必须用 `strconv.Atoi` 转成 `int`。
- Go 的错误通过返回值显式传递，不是默认抛异常。
- Go 的 `map[string]int` 要声明 key 和 value 的类型。
- Go 的函数可以返回多个值，所以统计函数可以返回 `(ClassStats, error)`。
- Go 的 `main` 应该尽量只组织流程，复杂逻辑拆到小函数里。

## 6. 容易踩坑

- 忘记 `strings.TrimSpace`，导致空行或前后空格影响解析。
- `strings.Split("", "\n")` 会得到 `[]string{""}`，不是空切片。
- `int / int` 会丢掉小数，平均分要转成 `float64`。
- 统计最高分和最低分时，不要随便假设分数一定存在。
- 错误信息不要只写 `invalid input`，要写清楚哪一行、哪个值错了。
- `main` 不要堆太多业务逻辑，否则 Day 8 拆包会很难。

## 7. 练习任务和验收标准

基础练习：

- 写 `parseStudent(line string) (Student, error)`。
- 写 `parseStudents(input string) ([]Student, error)`。
- 写 `parseScores(input string) ([]int, error)`。
- 写 `parseScore(raw string) (int, error)` 和 `validateScore(score int) error`。

统计练习：

- 写 `calculateStats(students []Student) (ClassStats, error)`。
- 写 `averageScore(scores []int) float64`。
- 写 `grade(score float64) string`。
- 写 `gradeCounts(students []Student) map[string]int`。

加餐练习：

- 写 `printReport(students []Student, stats ClassStats, counts map[string]int)`。
- 错误输入里要能说明第几行出错。
- 尝试支持空行跳过，但要在注释里说明为什么跳过。

验收标准：

- 能运行 `go run phase1-basics/day07/demo/main.go`。
- 补全 `exercise/main.go` 里的 TODO 后，能运行 `go run phase1-basics/day07/exercise/main.go`。
- 能解释从原始字符串到报表输出的数据流。
- 能解释为什么解析、统计、输出要拆成不同函数。
- 能解释 `[]Student`、`map[string]int`、`error` 在项目里的作用。

## 8. 完成记录

（完成后记录日期、练习结果和踩坑点）
