# Day 3 笔记：函数、作用域、多返回值

## 今天先记住一句话

函数用来把一段逻辑起名字；作用域决定变量在哪里能用；多返回值是 Go 里很常见的表达结果方式。

## 1. 函数基本写法

Go 函数写法：

```go
func 函数名(参数名 参数类型) 返回值类型 {
	return 返回值
}
```

例子：

```go
func grade(score int) string {
	if score >= 90 {
		return "A"
	}
	return "B"
}
```

调用函数：

```go
result := grade(86)
fmt.Println(result)
```

先把函数理解成一个小工具：输入参数，执行逻辑，返回结果。

## 2. 参数和返回值都要写类型

Go 不会像 JS 一样让参数随便传。函数签名会明确告诉你：

```go
func add(a int, b int) int {
	return a + b
}
```

如果连续几个参数类型一样，可以合并：

```go
func add(a, b int) int {
	return a + b
}
```

## 3. 多返回值

Go 函数可以一次返回多个值。

```go
func fahrenheit(celsius float64) (float64, int) {
	value := celsius*9/5 + 32
	return value, int(value)
}
```

接收时也要接两个：

```go
value, intValue := fahrenheit(36.6)
```

多返回值常用于：

- 一个函数同时返回多个计算结果。
- 一个函数返回结果和是否成功。
- 后面会经常看到 `value, err := xxx()`。

## 4. 命名返回值

返回值也可以提前起名字：

```go
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}
```

这里 `area` 和 `perimeter` 已经是函数内部变量，所以最后可以直接写 `return`。

建议：

- 简单函数可以用命名返回值。
- 逻辑变复杂时，写 `return area, perimeter` 更清楚。

## 5. 作用域

作用域就是“变量在哪些地方能被访问”。

```go
func main() {
	name := "Tom"
	fmt.Println(name)
}
```

`name` 只在 `main` 函数内部有效。

代码块也会产生作用域：

```go
if n := len("hello"); n > 3 {
	fmt.Println(n)
}

// fmt.Println(n) // 编译失败，n 已经离开作用域
```

## 6. 从 main 里拆函数

从 Day 3 开始，不要把所有逻辑都堆在 `main` 里。

推荐结构：

```go
func main() {
	result := grade(86)
	fmt.Println(result)
}

func grade(score int) string {
	if score >= 90 {
		return "A"
	}
	return "B"
}
```

`main` 负责组织流程，具体逻辑放到函数里。

## 和 JS/TS 的关键差异

- Go 函数参数必须写类型。
- Go 可以原生返回多个值，不需要返回数组或对象。
- Go 没有默认参数。
- Go 没有函数重载，同一个 package 下不能有两个同名函数。
- Go 的作用域更严格，未使用变量会直接编译失败。

## 容易踩坑

- 函数声明了返回值类型，就必须保证所有路径都有返回值。
- 同一个 package 下不能有同名函数。
- `:=` 不能在同一个作用域里只重复声明旧变量。
- 命名返回值虽然方便，但复杂逻辑里可能降低可读性。

## 练习时重点看什么

- `grade`：条件顺序是否从高到低。
- `fahrenheit`：返回两个值时，调用方是否也接收两个值。
- `maxAndCount`：空切片怎么处理。
- `countChars`：字符串里的字母、数字和其他字符如何判断。
- 加餐函数：把小逻辑拆出来，不要全写进 `main`。

## 完成记录

### 2026-06-22

- **练习结果：** 已完成 Day 3 所有基础练习和加餐练习，`go run phase1-basics/day03/exercise/main.go` 可正常运行。
- **完成内容：** `grade`、`fahrenheit`、`maxAndCount`、`countChars`、`fizzBuzz`、`isPrime`、`sumRange`。
- **踩坑点：** 切片不能用 `const` 声明；Go 的 `for` 循环条件之间用分号；函数签名声明了返回值时，所有路径都必须返回对应类型。
