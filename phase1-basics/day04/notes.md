# Day 4 笔记：数组、切片、append 与 map

## 今日目标

- 理解数组和切片的区别。
- 会用 `append` 动态构建切片。
- 会用 `len`、`cap` 观察切片长度和容量。
- 会用 `map` 做查询、计数、修改和删除。
- 会写 `value, ok := m[key]` 判断 key 是否存在。

## 核心知识点

### 1. 数组 array

数组长度固定，长度也是类型的一部分：

```go
var nums [3]int = [3]int{1, 2, 3}
```

`[3]int` 和 `[4]int` 是不同类型。日常业务中更常用切片。

### 2. 切片 slice

切片长度可以变化：

```go
scores := []int{80, 90}
scores = append(scores, 100)
```

- `len(slice)`：当前元素数量。
- `cap(slice)`：底层数组容量。
- `append` 可能复用原底层数组，也可能创建新的底层数组。

### 3. 切片截取

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(nums[:2])  // 1, 2
fmt.Println(nums[1:4]) // 2, 3, 4
fmt.Println(nums[2:])  // 3, 4, 5
```

### 4. map

```go
prices := map[string]int{
    "apple": 5,
}

prices["banana"] = 3
prices["apple"] = 6
delete(prices, "banana")
```

查询时推荐使用 `ok`：

```go
price, ok := prices["apple"]
```

如果只看 `price`，无法区分 key 不存在和 value 本身就是零值。

### 5. map 遍历顺序

`range map` 的顺序不固定，不能依赖输出顺序写业务逻辑。

## 今日练习

- 分数统计：总分、平均分、最高分、最低分。
- append 构建切片并观察 `len` / `cap`。
- 字符计数：`map[rune]int`。
- 商品价格表增删改查。
- 加餐：去重、最高分学生、购物车总价。

## 和前端 JS/TS 的关键差异

- Go 数组长度固定，slice 更接近 JS Array 的日常用途。
- `append` 会返回新 slice，必须接收返回值。
- map 查询不存在 key 时返回 value 类型的零值，所以要配合 `ok`。
- map 遍历顺序不稳定，不能当作有序对象使用。

## 疑问记录

（学习过程中有不懂的记在这里）

## 完成记录

（完成后记录日期、练习结果和踩坑点）
