# Day 10 笔记：综合复盘与并发入门

## 今天先记住一句话

并发不是让代码变复杂的理由；先把顺序流程写清楚，再用 goroutine 和 channel 把独立任务并行化。

## 1. 今天要掌握什么

1. 复盘 Day 1-Day 9 的基础语法和小项目数据流。
2. 理解 goroutine 是轻量级并发执行单元。
3. 理解 channel 用来在 goroutine 之间传递数据。
4. 理解 context 用来传递取消信号和超时。
5. 会把“多行成绩解析”改成并发处理。
6. 知道并发结果如果要保持顺序，需要带上原始下标。

## 2. 运行方式

运行 demo：

```bash
go run ./phase1-basics/day10/demo
```

运行测试：

```bash
go test ./phase1-basics/day10/demo
```

运行练习：

```bash
go run ./phase1-basics/day10/exercise
```

## 3. 新概念说明

### goroutine

在函数调用前加 `go`，就会启动一个 goroutine：

```go
go parseWorker(ctx, jobs, results)
```

它会和当前函数并发执行。你不能假设 goroutine 的执行顺序，所以需要 channel 来传递数据。

### channel

channel 是 goroutine 之间通信的管道：

```go
jobs := make(chan parseJob)
results := make(chan parseResult)
```

发送数据：

```go
jobs <- parseJob{Index: index, Line: line}
```

接收数据：

```go
result := <-results
```

### context

`context.Context` 用来告诉 goroutine：任务取消了，应该尽快退出。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

`ctx.Done()` 是一个 channel。如果超时或取消，它会收到信号：

```go
select {
case <-ctx.Done():
	return ctx.Err()
case result := <-results:
	// 使用结果
}
```

### select

`select` 可以同时等待多个 channel：

```go
select {
case <-ctx.Done():
	return
case job := <-jobs:
	// 处理任务
}
```

它常用于“等数据”和“等取消信号”之间做选择。

## 4. demo 代码逐段详解

### parser.go：顺序解析

`ParseStudents` 是顺序版本：

1. `splitStudentLines` 拆出多行输入。
2. 循环每一行。
3. 调用 `ParseStudent`。
4. append 到结果切片。

这个版本最容易理解，也是并发版本的基础。

### concurrent.go：并发解析

`ParseStudentsConcurrent` 做了几件事：

1. 把每一行包装成 `parseJob`，带上 `Index`。
2. 启动多个 worker goroutine。
3. worker 从 `jobs` channel 取任务。
4. worker 调用 `ParseStudent`。
5. worker 把结果写入 `results` channel。
6. 主 goroutine 收集所有结果，并按 `Index` 放回对应位置。

为什么结果要带 `Index`？因为并发执行时，第二行可能比第一行先解析完成。如果直接 append，输出顺序就可能乱。

### main.go：context 超时

demo 里使用：

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

这表示整个并发解析最多等 2 秒。真实项目里，context 常用于 HTTP 请求、数据库查询、外部 API 调用等需要超时和取消的场景。

## 5. 和 JS/TS 的关键差异

- JS/TS 常见并发模型是 Promise、async/await；Go 常见模型是 goroutine + channel。
- Go 的 goroutine 启动成本很低，但不是免费，仍然需要控制数量。
- Go 用 context 显式传递取消信号；不要用全局变量控制 goroutine 停止。
- Go 的 channel 是类型安全的，例如 `chan parseJob` 只能传 `parseJob`。

## 6. 容易踩坑

- 启动 goroutine 后没有接收结果，导致 goroutine 卡住。
- 忘记关闭 `jobs`，worker 一直等待新任务。
- 并发结果直接 append，导致顺序不稳定。
- 没有处理 `ctx.Done()`，取消后 goroutine 仍然继续跑。
- 为了练并发而把简单逻辑写复杂。先写顺序版，再改并发版。

## 7. 练习任务和验收标准

基础练习：

- 补全 `exercise/concurrent.go` 的 `ParseStudentsConcurrent`。
- 删除 `concurrent_test.go` 里的 `t.Skip`。
- 让 `go test ./phase1-basics/day10/exercise` 通过。

加餐练习：

- 给并发解析增加 worker 数量参数。
- 构造一条非法输入，确认错误信息包含行号。
- 给 `summary.md` 写完 Day 1-Day 10 的复盘。

验收标准：

- 能运行 `go test ./...`。
- 能解释 goroutine、channel、context 各自解决什么问题。
- 能解释为什么并发解析结果要保留输入顺序。
- 能讲清成绩统计 CLI 从输入到报表的数据流。

## 8. 完成记录

（完成后记录日期、练习结果和踩坑点）
