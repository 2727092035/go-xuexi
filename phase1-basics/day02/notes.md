# Day 2 笔记：条件与循环

## 核心知识点

### 1. if / else

- 条件不需要小括号。
- 大括号必须写。
- `if` 可以带初始化语句，例如：`if n := len(s); n > 3 { ... }`。
- 初始化语句里的变量只在 `if/else` 块内有效。

### 2. for 的三种写法

| 写法 | 示例 | 说明 |
|------|------|------|
| 三段式 | `for i := 0; i < 10; i++` | 最像前端里的 for |
| 类 while | `for count > 0` | Go 没有 while，用 for 代替 |
| 死循环 | `for { ... }` | 配合 `break` 退出 |

### 3. switch

- Go 的 `switch` 默认不会穿透，不需要手写 `break`。
- 一个 `case` 可以匹配多个值：`case "六", "日":`。
- 无表达式 `switch` 可以替代长 `if/else` 链。
- `fallthrough` 会强制进入下一个 case，但实际业务里很少用。

### 4. range

- 遍历 slice：`for i, v := range nums`。
- 只要值时用 `_` 丢弃索引：`for _, v := range nums`。
- 遍历 map 时顺序是随机的，不能依赖输出顺序。
- 遍历字符串时，index 是字节位置，value 是 rune。

## 和前端 JS/TS 的关键差异

- Go 没有 `while`，统一用 `for`。
- Go 的 `if` 条件必须是 bool，不能用 truthy/falsy。
- Go 的 `switch` 默认自动 break，和 JS 不同。
- Go 遍历字符串按 Unicode 字符处理，但索引仍然是字节位置。

## 完成记录

- 2026-06-18：Day 2 已完成。已能运行 if/else、for、switch、range，并完成成绩等级、FizzBuzz、九九乘法表、字符串统计、最大值统计练习。
