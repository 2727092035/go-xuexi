# Day 6 笔记：错误处理、输入校验、strings/strconv

## 今天先记住一句话

Go 的错误处理不是 `try/catch`，而是把可能失败的原因显式返回：成功时 `err == nil`，失败时 `err != nil`，调用方必须自己判断。

## 1. 今天要掌握什么

今天的重点不是“背 API”，而是学会一套 Go 里非常常见的数据处理流程：

1. 拿到字符串输入。
2. 用 `strings` 清理和拆分。
3. 用 `strconv` 做类型转换。
4. 用 `error` 表达失败原因。
5. 失败时提前返回，成功时继续处理。

这套流程以后会反复出现在命令行程序、HTTP 参数解析、JSON 字段校验、数据库数据处理里。

## 2. error 是什么

`error` 是 Go 里的内置接口类型。现在先把它理解成“一个能描述失败原因的值”。

常见函数签名：

```go
func parseScore(raw string) (int, error)
```

这个函数返回两个值：

- `int`：解析成功后的分数。
- `error`：失败原因。

调用方式：

```go
score, err := parseScore("90")
if err != nil {
	return err
}
fmt.Println(score)
```

约定：

- `err == nil`：没有错误，可以继续用结果。
- `err != nil`：发生错误，不要继续假装成功。

## 3. errors.New 和 fmt.Errorf

创建错误有两种常见方式。

固定错误文案用 `errors.New`：

```go
return errors.New("score is empty")
```

需要把变量写进错误信息时，用 `fmt.Errorf`：

```go
return fmt.Errorf("invalid score %q", raw)
```

`%q` 会把字符串带引号输出，例如 `"abc"`，比 `%s` 更适合错误提示。

## 4. strconv.Atoi：字符串转整数

`strconv.Atoi` 的作用是把字符串转成 `int`。

```go
score, err := strconv.Atoi("90")
```

它也返回两个值：

- 转换成功：`score == 90`，`err == nil`
- 转换失败：`score == 0`，`err != nil`

不要只看 `score`，因为失败时也可能返回 `0`。必须看 `err`。

```go
score, err := strconv.Atoi("abc")
if err != nil {
	return 0, fmt.Errorf("invalid score %q", "abc")
}
```

## 5. strings.TrimSpace 和 strings.Split

`strings.TrimSpace` 来自标准库 `strings` 包。它的作用是去掉字符串前后的空白字符，包括空格、换行、Tab。

```go
text := strings.TrimSpace(raw)
```

这行代码要拆开看：

- `strings`：标准库包名，专门放字符串处理函数。
- `TrimSpace`：函数名，意思是裁掉字符串两端的空白。
- `raw`：传给函数的参数，也就是原始输入字符串。
- `text`：函数返回的新字符串。

它不会修改 `raw`，而是返回一个新的清理结果：

```go
text := strings.TrimSpace(" 90 ")
fmt.Println(text) // "90"
```

为什么这里必须用它？因为 `strconv.Atoi(" 90 ")` 会失败，`strconv.Atoi("90")` 才能成功。

`strings.Split` 也来自 `strings` 包。它按指定分隔符拆字符串，返回一个字符串切片 `[]string`：

```go
parts := strings.Split("90,80,70", ",")
fmt.Println(parts) // []string{"90", "80", "70"}
```

这行代码也要拆开看：

- 第一个参数 `"90,80,70"`：要被拆分的原始字符串。
- 第二个参数 `","`：分隔符，表示遇到逗号就切开。
- 返回值 `parts`：字符串切片，也就是多个字符串组成的列表。

如果输入是 `"90, 80,100"`，按逗号拆完会得到：

```go
[]string{"90", " 80", "100"}
```

注意 `" 80"` 前面还有空格，所以每一段还要继续用 `strings.TrimSpace` 清理。

拆完以后通常配合 `for range` 遍历：

```go
for _, part := range parts {
	score, err := parseScore(part)
	if err != nil {
		return nil, err
	}
	scores = append(scores, score)
}
```

## 6. demo 代码逐段详解

### 第 1 段：error 的基本写法

`divide(a, b int) (int, error)` 表示这个函数可能失败。

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

关键点：

- 除数为 0 时，返回结果的零值 `0` 和一个错误。
- 成功时，返回真实结果和 `nil`。
- `error` 通常作为最后一个返回值。

### 第 2 段：parseScore

`parseScore` 把一个字符串解析成分数。

流程：

1. `strings.TrimSpace(raw)` 去掉原始输入 `raw` 前后的空白，返回清理后的 `text`。
2. 空字符串直接返回错误。
3. `strconv.Atoi(text)` 把清理后的字符串转成整数，并返回转换错误。
4. `validateScore(score)` 校验范围。
5. 全部成功才返回 `score, nil`。

这就是 Go 里很常见的“每一步都检查错误，失败就提前返回”。

关键代码：

```go
text := strings.TrimSpace(raw)
```

`raw` 是函数参数，代表外部传入的原始字符串。`TrimSpace` 返回的新值赋给 `text`，后面都使用 `text`，避免空格影响数字转换。

```go
score, err := strconv.Atoi(text)
```

`Atoi` 会尝试把字符串转成 `int`。它不是只返回一个数字，而是返回两个值：第一个是转换结果，第二个是错误。`err != nil` 时说明转换失败，必须提前返回。

### 第 3 段：parseScores

`parseScores("90,80,70")` 会返回 `[]int{90, 80, 70}`。

```go
parts := strings.Split(input, ",")
scores := make([]int, 0, len(parts))
```

`strings.Split(input, ",")` 的意思是：把 `input` 按逗号拆成多个字符串。逗号不会保留在结果里。

这里 `make([]int, 0, len(parts))` 的意思是：

- 长度先是 0。
- 容量预留为 `len(parts)`。
- 后面用 `append` 逐个加入分数。

这样写不是必须的，但比 `scores := []int{}` 更明确，也能少一些扩容。

### 第 4 段：parseStudent

输入格式：

```text
Tom,18,90|80|70
```

解析目标：

```go
Student{
	Name: "Tom",
	Age: 18,
	Scores: []int{90, 80, 70},
}
```

它会先用逗号拆成三段：姓名、年龄、分数字符串。分数字符串再用 `|` 拆成多个分数。

这段代码的重点是复用：

- 年龄自己用 `strconv.Atoi`。
- 分数复用 `parseScore`。
- 分数范围校验复用 `validateScore`。

不要把所有逻辑堆在一个函数里，否则 Day 7 的 CLI 会很难维护。

## 7. 和 JS/TS 的关键差异

- JS/TS 常见做法是 `try/catch` 或抛异常；Go 更常见的是返回 `error`。
- Go 的错误是普通返回值，不是隐藏控制流。
- Go 不会自动把 `"90"` 当成数字，必须用 `strconv.Atoi`。
- Go 函数可以返回多个值，所以 `(value, error)` 是常见模式。
- Go 里如果你拿到了 `err` 但不用，编译器不会报错；但代码习惯上必须检查。

## 8. 容易踩坑

- 忘记检查 `err`，继续使用错误结果。
- `strconv.Atoi(" 90 ")` 会失败，所以要先 `strings.TrimSpace`。
- `strings.Split("", ",")` 会得到 `[]string{""}`，不是空切片。
- 遇到错误时返回 `nil, err`，不要返回半成品数据。
- 错误信息不要只写 `invalid input`，最好带上具体值，例如 `invalid score "abc"`。
- `fmt.Errorf` 只是创建错误，不会自动打印；要返回或处理它。

## 9. 练习任务和验收标准

基础练习：

- 写 `parseScore(s string) (int, error)`。
- 写 `validateScore(score int) error`。
- 写 `parseScores(input string) ([]int, error)`。

加餐练习：

- 写 `divide(a, b int) (int, error)`。
- 写 `parseStudent(line string) (Student, error)`，输入格式：`"Tom,18,90|80|70"`。
- 写 `parseStudents(input string) ([]Student, error)`，输入格式：`"Tom,18,90|80|70;Jerry,20,100|95"`。
- `parseStudents` 要复用 `parseStudent`，并在某一行失败时返回带行号的错误。

验收标准：

- 能运行 `go run phase1-basics/day06/demo/main.go`。
- 补全 `exercise/main.go` 里的 TODO 后，能运行 `go run phase1-basics/day06/exercise/main.go` 并看到正确错误提示。
- 能解释为什么 Go 函数常返回 `(value, error)`。
- 能解释 `strings.Split`、`strings.TrimSpace`、`strconv.Atoi` 的作用。
- 能写出 `if err != nil { return ... }` 的提前返回逻辑。

## 10. 完成记录

（完成后记录日期、练习结果和踩坑点）
