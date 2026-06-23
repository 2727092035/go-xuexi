# Day 5 笔记：结构体、方法、指针接收者

## 今天先记住一句话

`struct` 用来描述一个业务对象，方法用来把行为绑定到这个对象上；要修改对象本身时，用指针接收者。

## 1. 为什么需要 struct

如果一个学生有姓名、年龄、分数，单独用几个变量会很散：

```go
name := "Tom"
age := 18
scores := []int{80, 90}
```

`struct` 可以把这些相关字段组织成一个整体：

```go
type Student struct {
	Name   string
	Age    int
	Scores []int
}
```

`Student` 是你自己定义的新类型。

## 2. 创建结构体

推荐写字段名，清楚且不容易写错顺序：

```go
s := Student{
	Name:   "Tom",
	Age:    18,
	Scores: []int{80, 90},
}
```

访问字段：

```go
fmt.Println(s.Name)
fmt.Println(s.Age)
```

修改字段：

```go
s.Age = 19
```

## 3. 方法是什么

方法是“带接收者的函数”。

```go
func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	sum := 0
	for _, score := range s.Scores {
		sum += score
	}
	return float64(sum) / float64(len(s.Scores))
}
```

调用方式：

```go
avg := s.Average()
```

把它读成：对这个 `Student`，调用它的 `Average` 方法。

## 4. 值接收者

值接收者拿到的是结构体的一份拷贝。

```go
func (s Student) Level() string {
	if s.Average() >= 90 {
		return "A"
	}
	return "B"
}
```

这种适合只读方法：只根据字段计算结果，不修改原对象。

## 5. 指针接收者

如果方法要修改结构体字段，就用指针接收者。

```go
func (s *Student) Birthday() {
	s.Age++
}
```

调用时通常还是这样写：

```go
s.Birthday()
```

Go 会帮你自动取地址，所以不需要每次都写 `(&s).Birthday()`。

## 6. 值接收者和指针接收者怎么选

先用这个规则：

- 只是读取和计算：用值接收者。
- 要修改字段：用指针接收者。
- 结构体很大，不想拷贝：也可以用指针接收者。

例子：

```go
func (s Student) Average() float64 {
	return 0
}

func (s *Student) AddScore(score int) {
	s.Scores = append(s.Scores, score)
}
```

`AddScore` 会修改 `Scores` 字段，所以用 `*Student`。

## 7. 字段大小写和可见性

Go 用首字母大小写控制包外是否可访问。

```go
type BankAccount struct {
	Owner   string
	balance int
}
```

- `Owner` 首字母大写，其他包可以访问。
- `balance` 首字母小写，其他包不能直接访问。

这是一种封装方式。余额不让外部随便改，而是通过方法控制：

```go
func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
}

func (a BankAccount) Balance() int {
	return a.balance
}
```

## 和 JS/TS 的关键差异

- Go 没有 `class` 关键字，常用 `struct + method` 组织对象和行为。
- 方法不写在结构体内部，而是写成带接收者的函数。
- 修改结构体字段时，要关注值拷贝和指针接收者。
- Go 用标识符首字母大小写控制包级可见性。

## 容易踩坑

- `func (s Student)` 是值接收者，修改 `s` 不会改到原对象。
- 要改变字段时，优先写 `func (s *Student)`。
- 方法名和字段名都要遵守大小写可见性规则。
- 结构体字面量推荐写字段名，不要依赖字段顺序。

## 练习时重点看什么

- `Student`：能不能把姓名、年龄、分数放进一个结构体。
- `Average`：空分数切片是否返回 0。
- `Birthday`：是否真的修改了原学生的年龄。
- `AddScore`：是否使用指针接收者并写回 `append` 结果。
- 购物车：`Product` 和 `CartItem` 分别表示什么。
- `BankAccount`：余额字段是否通过方法读写，而不是直接暴露。

## 疑问记录

（学习过程中有不懂的记在这里）

## 完成记录

（完成后记录日期、练习结果和踩坑点）
