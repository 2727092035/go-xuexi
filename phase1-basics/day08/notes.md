# Day 8 笔记：包、文件拆分、go mod

## 今天先记住一句话

Go 不是按文件组织程序，而是按 package 组织程序；同一个目录里的多个 `.go` 文件通常属于同一个 package，可以互相直接调用。

## 1. 今天要掌握什么

1. 一个目录通常对应一个 package。
2. 同目录下多个 `.go` 文件可以共同组成一个程序。
3. `go run file.go` 只运行指定文件，`go run .` 会运行当前目录整个 package。
4. 首字母大写的名字是导出的，可以被其他 package 调用。
5. 首字母小写的名字不导出，只能在当前 package 内部使用。
6. `go.mod` 记录模块名，跨 package import 时会用到模块路径。

## 2. 今天的目录结构

```text
phase1-basics/day08/
├── demo/
│   ├── main.go
│   ├── student.go
│   ├── parser.go
│   ├── stats.go
│   ├── report.go
│   └── utils/
│       └── score.go
└── exercise/
    ├── main.go
    ├── student.go
    ├── parser.go
    ├── stats.go
    ├── report.go
    └── utils/
        └── score.go
```

`demo` 里是完整示例，`exercise` 里是练习骨架。

运行 demo：

```bash
cd phase1-basics/day08/demo
go run .
```

运行练习：

```bash
cd phase1-basics/day08/exercise
go run .
```

## 3. 新概念说明

### package main

每个 `.go` 文件开头都要声明 package：

```go
package main
```

`package main` 表示这是一个可执行程序。只要这个 package 里有 `func main()`，Go 就能编译出一个可运行的命令。

同一个目录下的文件必须使用同一个 package 名。例如：

```text
demo/main.go      package main
demo/student.go   package main
demo/parser.go    package main
```

它们虽然在不同文件里，但仍然属于同一个 package，所以 `main.go` 可以直接调用 `parseStudents`、`CalculateStats`、`PrintReport`。

### go run file.go 和 go run .

`go run main.go` 只把 `main.go` 交给 Go 编译。

如果 `main.go` 调用了 `student.go` 里的 `Student`，单独运行就会报错：

```text
undefined: Student
```

`go run .` 表示运行当前目录的整个 package。Go 会把当前目录里同 package 的 `.go` 文件一起编译。

Day 8 开始，凡是一个目录里有多个文件，优先用：

```bash
go run .
```

### 导出标识符

Go 用首字母大小写控制能不能被其他 package 访问：

```go
func ValidateScore(score int) error
```

`ValidateScore` 首字母大写，是导出的。其他 package 可以这样调用：

```go
utils.ValidateScore(score)
```

如果写成：

```go
func validateScore(score int) error
```

它就只能在 `utils` package 内部使用，`parser.go` 不能调用它。

### go.mod 和 import 路径

仓库根目录的 `go.mod` 里有：

```go
module go-learn
```

所以 demo 里导入 utils 包时，路径要从模块名开始：

```go
import "go-learn/phase1-basics/day08/demo/utils"
```

这行代码的意思是：从当前模块 `go-learn` 下面找到 `phase1-basics/day08/demo/utils` 这个 package。

## 4. demo 代码逐段详解

### main.go：只组织流程

`main.go` 负责串起整个程序：

1. 准备原始输入。
2. 调用 `parseStudents` 解析学生。
3. 调用 `CalculateStats` 统计班级结果。
4. 调用 `GradeCounts` 统计等级人数。
5. 调用 `PrintReport` 输出报表。

`main.go` 不直接写解析和统计细节，这是拆文件的核心目的：让每个文件关注一类职责。

### student.go：学生结构和方法

`student.go` 放 `Student` 结构体和 `Average` 方法。

```go
type Student struct {
	Name   string
	Scores []int
}
```

结构体放在单独文件里，以后其他解析、统计、输出代码都可以复用它。

### parser.go：输入解析

`parser.go` 放所有字符串解析相关函数：

- `parseStudents`
- `parseStudent`
- `parseScores`
- `parseScore`

这些函数首字母小写，因为它们只在 `package main` 内部使用。它们不需要暴露给其他 package。

### stats.go：统计逻辑

`stats.go` 放统计相关类型和函数：

- `ClassStats`
- `CalculateStats`
- `Grade`
- `GradeCounts`

这些名字首字母大写，表示可以被其他 package 使用。今天先重点理解规则，后续项目里会更频繁地设计哪些函数该导出、哪些函数该隐藏。

### report.go：输出报表

`report.go` 只负责格式化输出，不负责计算。

这样拆的好处是：如果以后要把 CLI 输出换成 Web API 返回 JSON，只需要替换输出层，不需要重写解析和统计函数。

### utils/score.go：单独 package

`utils` 是一个单独目录，所以它是一个单独 package：

```go
package utils
```

`parser.go` 想调用它，必须 import：

```go
import "go-learn/phase1-basics/day08/demo/utils"
```

并且只能调用 `utils` 里导出的名字，比如 `ValidateScore`。

## 5. 和 JS/TS 的关键差异

- JS/TS 通常按文件 export/import；Go 是按 package 管理可见性。
- Go 不写 `export` 关键字，而是用首字母大小写控制导出。
- Go 同目录下多个文件不需要互相 import，它们天然属于同一个 package。
- Go 跨目录才需要 import。
- `go.mod` 类似项目的模块声明，import 本项目代码时会从 module 名开始。

## 6. 容易踩坑

- 多文件程序只运行 `go run main.go`，导致其他文件里的函数找不到。
- 同一个目录下混用 `package main` 和 `package utils`，Go 会直接报错。
- 跨 package 调用小写函数，例如 `utils.validateScore`，会报未导出错误。
- import 路径漏掉 module 名 `go-learn`。
- 把文件拆开后职责仍然混乱，比如 parser 里写报表输出。

## 7. 练习任务和验收标准

基础练习：

- 补全 `exercise/student.go` 的 `Average` 方法。
- 补全 `exercise/parser.go` 的 `parseStudents`、`parseStudent`、`parseScores`。
- 补全 `exercise/stats.go` 的 `calculateStats`、`grade`、`gradeCounts`。
- 补全 `exercise/report.go` 的 `printReport`。

加餐练习：

- 把 `utils.ValidateScore` 改成小写 `validateScore`，观察编译错误，再改回来。
- 故意执行 `go run main.go`，观察为什么找不到其他文件里的函数。
- 把 `Grade` 改成小写 `grade`，再判断哪些地方需要同步调整。

验收标准：

- 能运行 `go run ./phase1-basics/day08/demo` 或在 demo 目录执行 `go run .`。
- 补全练习后，能在 exercise 目录执行 `go run .`。
- 能解释为什么 Day 8 不能只运行 `main.go`。
- 能解释同目录多文件为什么不需要 import。
- 能解释为什么 `utils.ValidateScore` 必须首字母大写。

## 8. 完成记录

（完成后记录日期、练习结果和踩坑点）
