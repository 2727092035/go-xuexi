# Day 3 笔记：函数、作用域、多返回值

## 核心知识点

### 1. 函数基本写法

```go
func 函数名(参数名 参数类型) 返回值类型 {
    return 返回值
}
```

示例：

```go
func grade(score int) string {
    if score >= 90 {
        return "A"
    }
    return "B"
}
```

### 2. 多返回值

Go 函数可以一次返回多个值：

```go
func fahrenheit(celsius float64) (float64, int) {
    value := celsius*9/5 + 32
    return value, int(value)
}
```

接收时：

```go
value, intValue := fahrenheit(36.6)
```

### 3. 命名返回值

```go
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return
}
```

命名返回值适合简单函数。逻辑复杂时，直接写明确的 `return value1, value2` 更清楚。

### 4. 作用域

- 函数内部声明的变量只能在函数内部使用。
- `if` / `for` / `switch` 代码块内部声明的变量，只能在当前代码块里使用。
- 同一个作用域里不能用 `:=` 重复声明同名变量。

示例：

```go
if n := len("hello"); n > 3 {
    fmt.Println(n)
}
// fmt.Println(n) // 编译失败：n 已经离开作用域
```

### 5. 函数拆分的目标

从 Day 3 开始，`main` 不应该堆太多业务逻辑。

推荐结构：

```go
func main() {
    result := grade(86)
    fmt.Println(result)
}

func grade(score int) string {
    // 具体判断逻辑
}
```

## 和前端 JS/TS 的关键差异

- Go 函数参数必须写类型。
- Go 可以原生返回多个值，不需要返回数组或对象。
- Go 没有默认参数。
- Go 没有函数重载，同一个 package 下不能有两个同名函数。
- Go 的作用域更严格，未使用变量会直接编译失败。

## 今日练习

- `grade(score int) string`
- `fahrenheit(celsius float64) (float64, int)`
- `maxAndCount(nums []int) (max int, count int)`
- `countChars(s string) (letters int, digits int, others int)`
- `fizzBuzz(n int) string`
- 加餐：`isPrime(n int) bool`
- 加餐：`sumRange(start, end int) int`

## 疑问记录

（学习过程中有不懂的记在这里）

## 完成记录

### 2026-06-22

- **练习结果：** 已完成 Day 3 所有基础练习和加餐练习，`go run phase1-basics/day03/exercise/main.go` 可正常运行。
- **完成内容：** `grade`、`fahrenheit`、`maxAndCount`、`countChars`、`fizzBuzz`、`isPrime`、`sumRange`。
- **踩坑点：** 切片不能用 `const` 声明；Go 的 `for` 循环条件之间用分号；函数签名声明了返回值时，所有路径都必须返回对应类型。
