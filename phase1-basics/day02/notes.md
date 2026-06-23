# Day 2 笔记：条件与循环

## 今天先记住一句话

Go 的流程控制很少：判断用 `if` / `switch`，重复执行统一用 `for`。

## 1. if / else

Go 的 `if` 条件必须是 `bool`，不能像 JS 一样用 truthy / falsy。

```go
score := 86

if score >= 90 {
	fmt.Println("A")
} else if score >= 80 {
	fmt.Println("B")
} else {
	fmt.Println("C")
}
```

注意两点：

- 条件外面不写小括号。
- 大括号必须写。

## 2. if 可以带初始化语句

这种写法常用于“先算一个临时值，再判断”：

```go
if n := len("hello"); n > 3 {
	fmt.Println(n)
}
```

`n` 只在这个 `if/else` 内部有效。离开这段代码后，`n` 就不能再用了。

## 3. for 的三种写法

Go 没有 `while`，所有循环都用 `for`。

三段式，适合知道循环次数的场景：

```go
for i := 1; i <= 10; i++ {
	fmt.Println(i)
}
```

类 while，适合条件满足时一直循环：

```go
count := 3
for count > 0 {
	fmt.Println(count)
	count--
}
```

死循环，通常配合 `break`：

```go
for {
	if done {
		break
	}
}
```

## 4. switch

Go 的 `switch` 默认不会继续执行下一个 `case`，所以通常不需要写 `break`。

```go
day := "六"

switch day {
case "六", "日":
	fmt.Println("周末")
default:
	fmt.Println("工作日")
}
```

也可以不写 `switch` 后面的变量，用它替代很长的 `if/else`：

```go
switch {
case score >= 90:
	fmt.Println("A")
case score >= 80:
	fmt.Println("B")
default:
	fmt.Println("C")
}
```

## 5. range

`range` 用来遍历集合。

遍历切片：

```go
nums := []int{10, 20, 30}

for i, v := range nums {
	fmt.Println(i, v)
}
```

只需要值时，用 `_` 丢掉索引：

```go
for _, v := range nums {
	fmt.Println(v)
}
```

遍历字符串时，`range` 取到的是 Unicode 字符，也就是 `rune`：

```go
for index, r := range "Go语言" {
	fmt.Println(index, r)
}
```

这里的 `index` 是字节位置，不是第几个字符。这一点后面处理中文字符串时很重要。

## 和 JS/TS 的关键差异

- Go 没有 `while`，统一用 `for`。
- Go 的 `if` 条件必须是 `bool`。
- Go 的 `switch` 默认自动 `break`。
- Go 遍历字符串时，值是 `rune`，但索引是字节位置。
- `range map` 的顺序不固定，不能依赖输出顺序。

## 容易踩坑

- `if count {}` 这种写法不行，`count` 不是 `bool`。
- `for i := 0; i < 10; i++` 中间是分号，不是逗号。
- `range` 返回两个值，通常是 `索引/键` 和 `值`。
- 不用的返回值要用 `_` 接住，否则可能触发未使用变量错误。

## 练习时重点看什么

- 成绩等级练习：看 `if/else` 条件顺序是否正确。
- FizzBuzz：看 `%` 取余和多条件判断是否正确。
- 九九乘法表：看嵌套 `for`。
- 字符串统计：看 `range string` 对中文字符是否能正确处理。

## 完成记录

- 2026-06-18：Day 2 已完成。已能运行 if/else、for、switch、range，并完成成绩等级、FizzBuzz、九九乘法表、字符串统计、最大值统计练习。
