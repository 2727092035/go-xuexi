# Day 9 笔记：测试入门

## 今天先记住一句话

测试不是“证明代码永远正确”，而是把你已经想到的边界条件固定下来，防止以后改代码时悄悄改坏。

## 1. 今天要掌握什么

1. `_test.go` 文件只在 `go test` 时参与测试。
2. 测试函数必须写成 `func TestXxx(t *testing.T)`。
3. 表格驱动测试用一组测试数据覆盖多个输入。
4. `t.Run` 可以给每个用例单独命名。
5. `wantErr` 用来表达“这个输入应该返回错误”。
6. benchmark 用 `func BenchmarkXxx(b *testing.B)`。

## 2. 运行方式

运行 demo：

```bash
go run ./phase1-basics/day09/demo
```

运行测试：

```bash
go test ./phase1-basics/day09/demo
```

运行 benchmark：

```bash
go test -bench=. ./phase1-basics/day09/demo
```

## 3. 新语法和 API 说明

### testing 包

测试文件需要导入标准库 `testing`：

```go
import "testing"
```

`testing.T` 用在单元测试里，`testing.B` 用在 benchmark 里。

### 测试函数命名

```go
func TestGrade(t *testing.T) {
}
```

规则：

- 必须以 `Test` 开头。
- 后面接首字母大写的名字。
- 参数必须是 `t *testing.T`。
- 不需要返回值。

### 表格驱动测试

表格驱动测试就是把输入和期望结果放到一个 slice 里：

```go
tests := []struct {
	name string
	raw  string
	want int
}{
	{name: "valid", raw: "90", want: 90},
}
```

然后用 `for` 遍历每个用例：

```go
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		got, err := ParseScore(tt.raw)
		// 检查 got 和 err
	})
}
```

这种写法适合验证边界值，例如 `60`、`59.9`、`0`、`100`、`101`。

### t.Fatalf

`t.Fatalf` 表示测试失败，并停止当前测试用例：

```go
t.Fatalf("got %d, want %d", got, want)
```

它和 `fmt.Printf` 很像，但输出属于测试报告。

## 4. demo 代码逐段详解

### score.go：被测试的业务函数

`score.go` 里有四个核心函数：

- `ParseScore`
- `ValidateScore`
- `Grade`
- `MaxAndCount`

这些函数都比较小，输入明确、输出明确，很适合写单元测试。

### score_test.go：测试文件

`score_test.go` 和 `score.go` 在同一个目录、同一个 package 下，所以测试可以直接调用 `Grade`、`ParseScore`。

这也是 Go 测试最常见的组织方式：业务代码和测试代码放一起，文件名用 `_test.go` 区分。

### BenchmarkGrade

benchmark 用来粗略观察某个函数的性能：

```go
func BenchmarkGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Grade(88.5)
	}
}
```

`b.N` 由 Go 测试框架自动调整。你只要把要测的代码放进循环里。

## 5. 和 JS/TS 的关键差异

- Go 测试不需要额外安装 Jest/Vitest，标准库自带 `testing`。
- Go 用文件名 `_test.go` 识别测试文件。
- Go 没有 `expect(...).toBe(...)`，通常直接用 `if got != want { t.Fatalf(...) }`。
- Go 的表格驱动测试很常见，适合一次覆盖很多边界。

## 6. 容易踩坑

- 测试文件不以 `_test.go` 结尾，`go test` 不会识别。
- 测试函数写成 `testGrade` 或 `Testgrade`，不会被执行。
- 只测正常输入，不测空字符串、非法数字、越界值。
- `wantErr` 为 true 时还继续比较 `got`，导致测试逻辑混乱。
- benchmark 里忘记使用 `b.N` 循环。

## 7. 练习任务和验收标准

基础练习：

- 补全 `exercise/score.go` 的 `ParseScore`、`ValidateScore`、`Grade`、`MaxAndCount`。
- 删除 `score_test.go` 里的 `t.Skip`。
- 让 `go test ./phase1-basics/day09/exercise` 通过。

加餐练习：

- 给 Day 7 的 `parseStudents` 补测试。
- 至少覆盖空输入、非法数字、负数、超过 100、正常多学生。
- 给 `Grade` 增加 `89.9`、`79.9`、`69.9`、`59.9` 边界用例。

验收标准：

- 能运行 `go test ./...`。
- 能解释 `_test.go` 为什么不会进入普通 `go run`。
- 能解释表格驱动测试里的 `name`、`input`、`want`、`wantErr`。
- 能说明为什么边界值比随便写几个正常值更有价值。

## 8. 完成记录

（完成后记录日期、练习结果和踩坑点）
