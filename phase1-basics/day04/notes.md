# Day 4 笔记：数组、切片、append 与 map

## 今天先记住一句话

数组长度固定，切片才是日常最常用的动态列表；map 用来按 key 快速查 value。

## 1. 数组 array

数组是一段固定长度、固定类型的数据。

```go
var nums [3]int = [3]int{1, 2, 3}
```

这里的类型不是 `[]int`，而是 `[3]int`。

这意味着：

```go
var a [3]int
var b [4]int

// a = b // 编译失败，类型不同
```

数组长度是类型的一部分。日常业务里直接用数组不多，更多时候用切片。

## 2. 切片 slice

切片可以理解成“可以变长的列表”。

```go
scores := []int{80, 90}
scores = append(scores, 100)

fmt.Println(scores) // [80 90 100]
```

注意数组和切片的写法差别：

```go
[3]int{1, 2, 3} // 数组，长度固定
[]int{1, 2, 3}  // 切片，长度可变
```

## 3. len 和 cap

切片有两个常见概念：

- `len`：当前已经放了多少个元素。
- `cap`：底层数组最多还能容纳多少个元素。

```go
nums := []int{1, 2, 3}

fmt.Println(len(nums)) // 3
fmt.Println(cap(nums)) // 通常也是 3
```

`append` 时，如果容量够，Go 可能复用原来的底层数组；如果容量不够，Go 会创建一个更大的底层数组，再把旧元素拷过去。

所以 `append` 一定要接收返回值：

```go
nums = append(nums, 4)
```

不要只写：

```go
append(nums, 4) // 编译失败，返回值没用
```

## 4. 切片截取

切片可以从已有切片里截一段出来。

```go
nums := []int{1, 2, 3, 4, 5}

fmt.Println(nums[:2])  // [1 2]
fmt.Println(nums[1:4]) // [2 3 4]
fmt.Println(nums[2:])  // [3 4 5]
```

规则是左闭右开：包含左边位置，不包含右边位置。

`nums[1:4]` 取的是索引 1、2、3，不包含索引 4。

## 5. map

map 是 key-value 结构，适合做查询、计数、去重。

```go
prices := map[string]int{
	"apple":  5,
	"banana": 3,
}
```

读取：

```go
price := prices["apple"]
```

新增或修改：

```go
prices["orange"] = 4
prices["banana"] = 6
```

删除：

```go
delete(prices, "apple")
```

## 6. map 查询为什么要用 ok

如果 key 不存在，map 会返回 value 类型的零值。

```go
prices := map[string]int{
	"banana": 0,
}

price := prices["pear"]
fmt.Println(price) // 0
```

只看 `price`，你不知道这是“pear 不存在”，还是“pear 的价格就是 0”。

所以推荐这样写：

```go
price, ok := prices["pear"]
fmt.Println(price, ok)
```

- `ok == true`：key 存在。
- `ok == false`：key 不存在。

## 7. 用 map 做计数

统计字符出现次数时，可以用 `map[rune]int`。

```go
func charCount(s string) map[rune]int {
	count := map[rune]int{}
	for _, r := range s {
		count[r]++
	}
	return count
}
```

`count[r]++` 能成立，是因为不存在的 key 默认返回 `0`，第一次加一后就变成 `1`。

## 8. map 遍历顺序不固定

```go
for k, v := range count {
	fmt.Println(k, v)
}
```

每次输出顺序都可能不同。不要依赖 map 的遍历顺序写业务逻辑。

如果真的需要稳定顺序，后面可以先把 key 放进切片，再排序。

## 和 JS/TS 的关键差异

- Go 数组长度固定，slice 更接近 JS Array 的日常用途。
- `append` 会返回新 slice，必须接收返回值。
- map 查询不存在 key 时返回零值，所以要配合 `ok`。
- map 遍历顺序不稳定，不能当作有序对象使用。

## 容易踩坑

- 空切片没有第 0 个元素，访问 `scores[0]` 前要先判断 `len(scores)`。
- `append` 后要写回原变量：`nums = append(nums, x)`。
- `map` 使用前要初始化：`m := map[string]int{}`。
- 函数不能声明在另一个函数内部。
- 同一个作用域里不能重复 `:=` 声明同名变量。

## 练习时重点看什么

- 分数统计：空切片是否返回 `0, 0, 0, 0`。
- append 练习：观察 `len` 和 `cap` 的变化，不要求死记增长规则。
- 字符计数：先初始化 map，再用 `range` 遍历字符串。
- 商品价格表：查询时练习 `value, ok := m[key]`。
- 去重：用切片保持顺序，用 map 可以提升查重效率。
- 购物车总价：只计算价格表里存在的商品。

## 完成记录

### 2026-06-23

- **练习结果：** 已完成 Day 4 基础练习和加餐练习，`go run phase1-basics/day04/exercise/main.go` 可正常运行。
- **完成内容：** `scoreStats`、`charCount`、`unique`、`topScore`、`cartTotal`。
- **踩坑点：** `map` 要先初始化；空切片不能直接访问 `scores[0]`；函数不能写在 `main` 函数内部；同一作用域里变量名重复时不能继续用 `:=`。
