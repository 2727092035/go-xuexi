# Go 后续 7 天强化学习计划

背景：Day 1 已完成变量、基础类型、零值、类型转换和常量计算；Day 2 已完成条件、循环、`switch`、`range`、FizzBuzz、九九乘法表、字符串统计和最大值统计。你现在反馈“每天内容有点少”，所以后续节奏调整为：每天一个主知识点 + 一组基础练习 + 一组加餐练习 + 一个连续小项目。

建议每天投入：60 到 90 分钟。

## 每天固定节奏

1. 先运行当天 `demo/main.go`，看输出和现象。
2. 再读代码，把不懂的点写到当天 `notes.md`。
3. 完成 `exercise/main.go` 的基础题，确保 `go run phase1-basics/dayXX/exercise/main.go` 能跑通。
4. 完成 1 到 2 个加餐题，强迫自己把知识组合起来用。
5. 最后用 3 到 5 句话复盘：今天学了什么、卡在哪里、明天要补什么。

## Day 3：函数、作用域、多返回值

状态：已完成（2026-06-22）

主题：`func`、参数、返回值、多返回值、命名返回值、作用域

目标：
- 能把 Day 2 的练习拆成独立函数。
- 能写返回多个值的函数。
- 能理解变量块级作用域，避免 `:=` 重复声明冲突。

基础练习：
- `grade(score int) string`
- `fahrenheit(celsius float64) (float64, int)`
- `maxAndCount(nums []int) (max int, count int)`
- `countChars(s string) (letters int, digits int, others int)`

加餐练习：
- 把 Day 2 的 FizzBuzz 改成 `fizzBuzz(n int) string`。
- 写 `isPrime(n int) bool` 判断质数，再打印 1 到 100 的所有质数。
- 写 `sumRange(start, end int) int`，支持 start 大于 end 时自动交换。

验收：
- `main` 只负责组织流程，具体逻辑都放到函数里。
- 每个函数输入明确、输出明确，不依赖全局变量。

## Day 4：数组、切片、append 与 map

主题：array、slice、`append`、`len`、`cap`、map 增删改查

目标：
- 理解数组长度是类型的一部分，日常更常用 slice。
- 会用 `append` 动态添加元素。
- 会用 `map[string]int` 做计数和查询。
- 会写 `value, ok := m[key]` 并解释 `ok`。

基础练习：
- 用 slice 保存一组分数，计算总分、平均分、最高分、最低分。
- 用 map 统计字符串中每个字符出现次数。
- 用 map 做商品价格表，支持查询、修改、删除。

加餐练习：
- 写 `unique(nums []int) []int`，对整数切片去重。
- 写 `topScore(students map[string]int) (name string, score int)`，找最高分学生。
- 写一个简单购物车：商品名到数量，再结合价格表计算总价。

验收：
- 能区分 slice 和 array。
- 能说明 map 遍历顺序为什么不能依赖。

## Day 5：结构体、方法、指针接收者

主题：`struct`、结构体字面量、方法、值接收者、指针接收者

目标：
- 会定义结构体表达业务对象。
- 会给结构体写方法。
- 初步理解修改字段时为什么要用指针接收者。

基础练习：
- 定义 `Student`，包含 `Name`、`Age`、`Scores`。
- 写 `Average()`、`Level()` 方法。
- 写 `Birthday()` 方法，把年龄加 1，使用指针接收者。

加餐练习：
- 定义 `Product` 和 `CartItem`，计算购物车总价。
- 给 `Student` 增加 `AddScore(score int)` 方法，非法分数不加入。
- 定义 `BankAccount`，实现 `Deposit`、`Withdraw`、`Balance`。

验收：
- 能区分函数 `Average(s Student)` 和方法 `s.Average()`。
- 能说明值接收者和指针接收者的区别。

## Day 6：错误处理、输入校验、strings/strconv

主题：`error`、`errors.New`、`fmt.Errorf`、`strings`、`strconv`

目标：
- 理解 Go 不用异常作为主流程，常见模式是返回 `error`。
- 会检查错误并提前返回。
- 会把字符串转数字，并处理转换失败。

基础练习：
- 写 `parseScore(s string) (int, error)`。
- 写 `validateScore(score int) error`，限制 0 到 100。
- 写 `parseScores(input string) ([]int, error)`，输入如 `"90,80,abc"` 时返回错误。

加餐练习：
- 写 `divide(a, b int) (int, error)`，处理除以 0。
- 写 `parseStudent(line string) (Student, error)`，输入格式：`"Tom,18,90|80|70"`。
- 把错误信息写清楚，例如：`invalid score "abc"`。

验收：
- 能熟练写出 `if err != nil { return ... }`。
- 能说明为什么 Go 倾向显式处理错误。

## Day 7：小项目 1：成绩统计 CLI

主题：组合函数、slice、map、struct、error

项目目标：
- 输入一组学生姓名和分数。
- 校验分数范围必须是 0 到 100。
- 输出每个学生等级、全班平均分、最高分、最低分。
- 对非法输入返回清晰错误。

建议数据格式：
- `Tom,90`
- `Jerry,78`
- `Lucy,100`

最低要求：
- 至少有 `Student` 结构体。
- 至少有 4 个函数：解析、校验、统计、输出。
- 不把所有逻辑堆在 `main` 里。

加餐：
- 支持一个学生多个分数：`Tom,90|80|70`。
- 输出班级分数段统计：A/B/C/D/F 各多少人。
- 输出 Top 3 学生。

验收：
- 能从空输入、非法数字、越界分数里返回错误。
- 能清楚解释整个程序的数据流。

## Day 8：包、文件拆分、go mod

主题：`go mod init`、包名、文件拆分、导出标识符

目标：
- 理解一个目录通常对应一个 package。
- 会把项目拆成多个 `.go` 文件。
- 知道首字母大写代表导出。

基础练习：
- 把 Day 7 项目拆成 `main.go`、`student.go`、`stats.go`、`parser.go`，单独放在项目目录中。
- 初始化 `go mod init`。
- 保持 `go run .` 可以运行。

加餐练习：
- 新建 `utils` 包，放通用校验函数。
- 故意把导出/非导出名字改错，观察编译错误。

验收：
- 能说明 `go run exercises.go` 和 `go run .` 的区别。
- 能说明 package 名和目录组织的基本关系。

## Day 9：测试入门

主题：`testing`、表格驱动测试、边界用例

目标：
- 会给函数写单元测试。
- 会用表格驱动测试覆盖多组输入。
- 初步形成“写完逻辑就验证”的习惯。

基础练习：
- 给 `grade` 写测试。
- 给 `parseScore` 写测试。
- 给 `maxAndCount` 写测试。

加餐练习：
- 给 Day 7 成绩统计项目补测试。
- 至少覆盖：空输入、非法数字、负数、超过 100、正常多学生。

验收：
- 能运行 `go test ./...`。
- 能说明为什么测试文件以 `_test.go` 结尾。

## Day 10：综合复盘与重构

主题：复盘、重构、补漏洞、整理笔记

目标：
- 把 Day 7 项目重构到更清晰。
- 补齐测试和错误处理。
- 总结 Go 和前端 JS/TS 的关键差异。

任务：
- 整理一篇 `summary.md`：变量、条件循环、函数、slice/map、struct、error、测试。
- 重构成绩统计 CLI，让函数职责更清楚。
- 删除无用变量、无用注释和重复逻辑。

加餐：
- 给项目增加命令行参数。
- 支持从文本文件读取学生成绩。

验收：
- `go test ./...` 通过。
- 自己能从零复述项目结构和核心逻辑。

## 加量后的学习原则

- 每天不要只“看懂”，必须至少写 30 到 60 行代码。
- 遇到报错先读英文错误，不要马上问答案。
- 每天保留一个“加餐题”，可以不会，但要尝试 15 分钟。
- 每 3 天做一个小项目或重构，否则语法会很散。
- 当前阶段先别急着学 Web 框架，先把 Go 的基础语法、函数、结构体、错误处理打牢。
