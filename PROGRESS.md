# Go 学习进度

## 学习计划总览

| 阶段     | 内容     | 周期      | 状态   |
| -------- | -------- | --------- | ------ |
| 第一阶段 | 语法基础强化 | Day 1-10  | 进行中 |
| 第二阶段 | Web API 与数据库 | Day 11-25 | 未开始 |
| 第三阶段 | 工程化与中间件 | Day 26-40 | 未开始 |
| 第四阶段 | 综合项目实战 | Day 41-60 | 未开始 |

## 当前进度

- **当前阶段：** 第一阶段 · 语法基础强化
- **当前天数：** Day 6（错误处理、输入校验、strings/strconv）
- **累计学习天数：** 5
- **最近学习日期：** 2026-06-23（Day 4 有记录；Day 5 为计划调整前已完成，具体日期未记录）
- **今日状态：** Day 5 练习已完成，下一步进入 Day 6
- **当前计划文件：** `LEARNING_PLAN_60_DAYS.md`（60 天总计划） + `phase1-basics/next-5-days-plan.md`（Day 3-Day 10 强化计划）
- **60 天计划位置：** 已完成 Day 1-Day 5，当前处于 Day 6，剩余 Day 6-Day 60 共 55 天
- **目录规范：** 每天拆分为 `demo/main.go`、`exercise/main.go`、`notes.md`，避免多个 `main` 冲突

---

## 学习日志

### Day 1 — 变量与类型

- **日期：** 2026-06-17
- **学习内容：** 变量声明三种方式、基本数据类型、零值、类型转换、常量、fmt 格式化输出
- **产出文件：** `phase1-basics/day01/`（demo/main.go 示例 + exercise/main.go 练习 + notes.md 笔记）
- **掌握程度：** 已完成
- **疑问/待解决：** —
- **下次计划：** Day 2 — 条件与循环（if/for/switch/range）

### Day 2 — 条件与循环

- **日期：** 2026-06-18
- **学习内容：** if/else、for 三种写法、break/continue、switch、range 遍历 slice/map/string
- **产出文件：** `phase1-basics/day02/`（demo/main.go 示例 + exercise/main.go 练习 + notes.md 笔记）
- **完成练习：** 成绩等级、FizzBuzz、九九乘法表、字符串统计、最大值与出现次数统计
- **掌握程度：** 已完成
- **疑问/待解决：** 后续需要继续熟悉 `range` 遍历字符串时 index 是字节位置、value 是 rune
- **下次计划：** Day 3 — 函数、作用域、多返回值

### Day 3 — 函数、作用域、多返回值

- **日期：** 2026-06-22
- **学习内容：** func、参数、返回值、多返回值、命名返回值、作用域、函数拆分
- **产出文件：** `phase1-basics/day03/`（demo/main.go 示例 + exercise/main.go 练习 + notes.md 笔记）
- **完成练习：** 成绩等级函数、温度转换、多返回值、最大值与次数统计、字符串统计、FizzBuzz 函数拆分、质数判断、区间求和
- **掌握程度：** 已完成
- **疑问/待解决：** 需要继续熟悉 Go 的 `for` 语法、切片变量声明、函数返回值签名与实际 return 的一致性
- **下次计划：** Day 4 — 数组、切片、append 与 map

### Day 4 — 数组、切片、append 与 map

- **日期：** 2026-06-23
- **学习内容：** array、slice、`append`、`len`、`cap`、map 增删改查、`value, ok`、map 计数和去重
- **产出文件：** `phase1-basics/day04/`（demo/main.go 示例 + exercise/main.go 练习 + notes.md 笔记）
- **完成练习：** `scoreStats`、`charCount`、`unique`、`topScore`、`cartTotal`
- **掌握程度：** 已完成
- **疑问/待解决：** 需要继续注意 map 初始化、空切片访问、`append` 写回、同一作用域 `:=` 重复声明
- **下次计划：** Day 5 — 结构体、方法、指针接收者

### Day 5 — 结构体、方法、指针接收者

- **日期：** 计划调整前已完成（具体日期未记录）
- **学习内容：** `struct`、结构体字面量、方法、值接收者、指针接收者、字段可见性
- **产出文件：** `phase1-basics/day05/`（demo/main.go 示例 + exercise/main.go 练习 + notes.md 笔记）
- **完成练习：** `Student`、`Average`、`Level`、`Birthday`、`AddScore`、`Product`/`CartItem`、`BankAccount`
- **掌握程度：** 已完成
- **疑问/待解决：** 后续在包拆分时继续熟悉字段和方法的大小写可见性
- **下次计划：** Day 6 — 错误处理、输入校验、strings/strconv

---

## 计划调整记录

### 2026-06-18 — 目录结构调整

- 已将 Day 1 和 Day 2 拆分为 `demo/main.go` 与 `exercise/main.go`。
- 已补充 `phase1-basics/day02/notes.md`。
- 这样 `go test ./...` 不会再因为同目录多个 `func main()` 失败。

### 2026-06-18 — 学习节奏加量

- 原计划偏轻，已调整为每天 60 到 90 分钟。
- 后续每天包含：主知识点、基础练习、加餐练习、验收标准。
- Day 7 增加成绩统计 CLI 小项目。
- Day 9 增加单元测试入门。
- Day 10 增加综合复盘与重构。

### 2026-06-22 — Day 3 学习完成

- 已完成 `phase1-basics/day03/exercise/main.go` 所有基础题和加餐题。
- 已验证 `go run phase1-basics/day03/exercise/main.go` 可正常运行。
- 当前学习进度推进到 Day 4。

### 2026-06-24 — 总计划压缩到 60 天

- 原 90 天路线压缩为 60 天，整体节奏调整为每天 90 到 120 分钟。
- 压缩天数不代表减少学习内容，而是提高每天学习密度。
- 第一阶段从 Day 1-Day 14 压缩为 Day 1-Day 10，语法学习和小项目合并推进。
- 第二阶段 Day 11-Day 25 进入 Web API 与数据库，不再继续拉长基础语法。
- 第三阶段 Day 26-Day 40 补工程化、中间件、测试、Docker。
- 第四阶段 Day 41-Day 60 完成综合项目实战，并覆盖安全、性能、部署、微服务入门等进阶主题。
- 新增总计划文件：`LEARNING_PLAN_60_DAYS.md`。
- 已在总计划中补充“基于当前进度的执行表”：从 Day 6 接着学，明确接下来 7 天、Day 11-Day 25、Day 26-Day 40、Day 41-Day 60 的执行重点。

### 2026-06-24 — 当前进度修正到 Day 6

- 你反馈在调整学习计划前已经完成到 Day 5。
- 已将当前进度修正为：Day 1-Day 5 已完成，下一步进入 Day 6。
- 60 天计划剩余范围同步修正为 Day 6-Day 60，共 55 天。

### 2026-06-24 — 后续 demo 加强语法教学

- 从 Day 6 开始，后续每天的 `demo/main.go` 要按语法点分段，并用中文注释说明语法目的、核心写法和输出含义。
- 后续每天的 `notes.md` 要包含语法详解、demo 逐段解析、常见坑、JS/TS 对比、练习验收。
- 项目日也不能只写项目代码，需要先解释当天用到的语法、标准库 API 或工程概念。
