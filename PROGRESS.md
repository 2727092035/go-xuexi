# Go 学习进度

## 学习计划总览

| 阶段     | 内容     | 周期      | 状态   |
| -------- | -------- | --------- | ------ |
| 第一阶段 | 语法基础 | Day 1-14  | 进行中 |
| 第二阶段 | Web 开发 | Day 15-35 | 未开始 |
| 第三阶段 | 工程化   | Day 36-56 | 未开始 |
| 第四阶段 | 进阶实战 | Day 57-90 | 未开始 |

## 当前进度

- **当前阶段：** 第一阶段 · 语法基础
- **当前天数：** Day 4（数组、切片、append 与 map）
- **累计学习天数：** 3
- **最近学习日期：** 2026-06-22
- **今日状态：** Day 3 练习已完成，下一步进入 Day 4
- **当前计划文件：** `phase1-basics/next-5-days-plan.md`（已调整为 Day 3-Day 10 强化计划）
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
