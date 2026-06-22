# Day 5 笔记：结构体、方法、指针接收者

## 今日目标

- 会用 `struct` 表达业务对象。
- 会创建结构体字面量并访问字段。
- 会给结构体定义方法。
- 能区分值接收者和指针接收者。
- 初步理解小写字段的封装意义。

## 核心知识点

### 1. 结构体 struct

结构体用来把一组相关字段组织成一个业务对象：

```go
type Student struct {
    Name   string
    Age    int
    Scores []int
}
```

创建结构体时推荐写字段名：

```go
s := Student{
    Name:   "Tom",
    Age:    18,
    Scores: []int{80, 90},
}
```

### 2. 方法

方法是带接收者的函数：

```go
func (s Student) Average() float64 {
    return 0
}
```

调用方式：

```go
avg := s.Average()
```

### 3. 值接收者

值接收者拿到结构体的一份拷贝，适合只读方法：

```go
func (s Student) Level() string {
    return "A"
}
```

### 4. 指针接收者

需要修改结构体字段时，用指针接收者：

```go
func (s *Student) Birthday() {
    s.Age++
}
```

Go 调用时通常可以自动取地址：

```go
s.Birthday()
```

### 5. 字段大小写

- 大写字段：导出，其他包可访问。
- 小写字段：不导出，其他包不能直接访问。

后面拆包后会更明显。现在可以先用 `balance int` 练习封装意识。

## 今日练习

- 定义 `Student`。
- 实现 `Average()`、`Level()`、`Birthday()`。
- 加餐：`AddScore(score int) bool`。
- 加餐：`Product`、`CartItem` 和购物车总价。
- 加餐：`BankAccount` 的存款、取款、余额查询。

## 和前端 JS/TS 的关键差异

- Go 没有 class 关键字，常用 `struct + method` 组织对象行为。
- 方法不写在结构体内部，而是用接收者绑定到类型上。
- 修改结构体字段时要关注值拷贝和指针接收者。
- Go 用标识符首字母大小写控制包级可见性。

## 疑问记录

（学习过程中有不懂的记在这里）

## 完成记录

（完成后记录日期、练习结果和踩坑点）
