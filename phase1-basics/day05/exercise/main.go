// Day 5 练习题
// 完成下面的练习，验证你对结构体、方法、指针接收者的理解。
// 运行：go run phase1-basics/day05/exercise/main.go

package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Scores []int
}

func main() {
	fmt.Println("练习 1：Student 结构体")
	// 要求：
	// 1. 创建 Student，姓名 Tom，年龄 18，分数 80、90、100。
	// 2. 打印姓名、年龄、分数。
	// 在这里写代码 👇

	fmt.Println("\n练习 2：Average 和 Level 方法")
	// 要求：
	// 1. 给 Student 实现 Average() float64。
	// 2. 给 Student 实现 Level() string。
	// 3. Level 规则：平均分 >=90 A，>=80 B，>=70 C，>=60 D，否则 F。
	// 4. 空分数平均分返回 0。
	// 在这里写代码 👇

	fmt.Println("\n练习 3：Birthday 指针接收者")
	// 要求：
	// 1. 给 Student 实现 Birthday()，年龄加 1。
	// 2. 必须使用指针接收者。
	// 3. 调用后打印年龄，确认原对象被修改。
	// 在这里写代码 👇

	fmt.Println("\n加餐 1：AddScore")
	// 要求：
	// 1. 给 Student 实现 AddScore(score int) bool。
	// 2. 只接受 0 到 100 的分数。
	// 3. 合法分数加入 Scores 并返回 true，非法分数不加入并返回 false。
	// 在这里写代码 👇

	fmt.Println("\n加餐 2：购物车")
	// 要求：
	// 1. 定义 Product，字段 Name string、Price int。
	// 2. 定义 CartItem，字段 Product Product、Quantity int。
	// 3. 给 CartItem 实现 Subtotal() int。
	// 4. 创建两个购物车项，计算总价。
	// 在这里写代码 👇

	fmt.Println("\n加餐 3：银行账户")
	// 要求：
	// 1. 定义 BankAccount，字段 Owner string、balance int。
	// 2. 实现 Deposit(amount int) bool，金额大于 0 才能存入。
	// 3. 实现 Withdraw(amount int) bool，金额大于 0 且余额足够才可取出。
	// 4. 实现 Balance() int，返回余额。
	// 5. balance 字段保持小写，不让外部直接改。
	// 在这里写代码 👇

	fmt.Println("\n练习完成！")
}

// 在下面补充类型和方法。

func (s Student) Average() float64 {
	// TODO
	return 0
}

func (s Student) Level() string {
	// TODO
	return ""
}

func (s *Student) Birthday() {
	// TODO
}

func (s *Student) AddScore(score int) bool {
	// TODO
	return false
}
