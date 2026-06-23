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

type Product struct {
	Name  string
	Price int
}

type CartItem struct {
	Product  Product
	Quantity int
}

type BankAccount struct {
	Owner   string
	balance int
}

func main() {
	fmt.Println("练习 1：Student 结构体")
	// 要求：
	// 1. 创建 Student，姓名 Tom，年龄 18，分数 80、90、100。
	// 2. 打印姓名、年龄、分数。
	// 在这里写代码 👇
	student := Student{
		Name:   "Tom",
		Age:    18,
		Scores: []int{80, 90, 100},
	}
	fmt.Println("学生:", student)
	fmt.Printf("学生详细信息: %+v\n", student)
	fmt.Println("姓名:", student.Name)
	fmt.Println("年龄:", student.Age)
	fmt.Println("分数:", student.Scores)

	fmt.Println("\n练习 2：Average 和 Level 方法")
	// 要求：
	// 1. 给 Student 实现 Average() float64。
	// 2. 给 Student 实现 Level() string。
	// 3. Level 规则：平均分 >=90 A，>=80 B，>=70 C，>=60 D，否则 F。
	// 4. 空分数平均分返回 0。
	// 在这里写代码 👇
	fmt.Printf("平均分: %.1f\n", student.Average())
	fmt.Println("等级:", student.Level())

	fmt.Println("\n练习 3：Birthday 指针接收者")
	// 要求：
	// 1. 给 Student 实现 Birthday()，年龄加 1。
	// 2. 必须使用指针接收者。
	// 3. 调用后打印年龄，确认原对象被修改。
	// 在这里写代码 👇
	student.Birthday()
	fmt.Println("生日后年龄:", student.Age)

	fmt.Println("\n加餐 1：AddScore")
	// 要求：
	// 1. 给 Student 实现 AddScore(score int) bool。
	// 2. 只接受 0 到 100 的分数。
	// 3. 合法分数加入 Scores 并返回 true，非法分数不加入并返回 false。
	// 在这里写代码 👇
	fmt.Println("添加 100:", student.AddScore(100))
	fmt.Println("添加 101:", student.AddScore(101))
	fmt.Println("添加 0:", student.AddScore(0))
	fmt.Println("添加 -1:", student.AddScore(-1))
	fmt.Println("当前分数:", student.Scores)

	fmt.Println("\n加餐 2：购物车")
	// 要求：
	// 1. 定义 Product，字段 Name string、Price int。
	// 2. 定义 CartItem，字段 Product Product、Quantity int。
	// 3. 给 CartItem 实现 Subtotal() int。
	// 4. 创建两个购物车项，计算总价。
	// 在这里写代码 👇
	keyboard := CartItem{
		Product: Product{
			Name:  "Keyboard",
			Price: 199,
		},
		Quantity: 2,
	}
	apple := CartItem{
		Product: Product{
			Name:  "Apple",
			Price: 10,
		},
		Quantity: 3,
	}
	total := keyboard.Subtotal() + apple.Subtotal()
	fmt.Println("购物车项 1:", keyboard)
	fmt.Println("购物车项 1 小计:", keyboard.Subtotal())
	fmt.Println("购物车项 2:", apple)
	fmt.Println("购物车项 2 小计:", apple.Subtotal())
	fmt.Println("购物车总价:", total)

	fmt.Println("\n加餐 3：银行账户")
	// 要求：
	// 1. 定义 BankAccount，字段 Owner string、balance int。
	// 2. 实现 Deposit(amount int) bool，金额大于 0 才能存入。
	// 3. 实现 Withdraw(amount int) bool，金额大于 0 且余额足够才可取出。
	// 4. 实现 Balance() int，返回余额。
	// 5. balance 字段保持小写，不让外部直接改。
	// 在这里写代码 👇
	account := BankAccount{
		Owner:   "John",
		balance: 1000,
	}
	fmt.Println("账户所有人:", account.Owner)
	fmt.Println("初始余额:", account.Balance())
	fmt.Println("存入 100:", account.Deposit(100))
	fmt.Println("取出 50:", account.Withdraw(50))
	fmt.Println("取出 2000:", account.Withdraw(2000))
	fmt.Println("最终余额:", account.Balance())

	fmt.Println("\n练习完成！")
}

// 在下面补充类型和方法。

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

func (s Student) Level() string {
	average := s.Average()
	switch {
	case average >= 90:
		return "A"
	case average >= 80:
		return "B"
	case average >= 70:
		return "C"
	case average >= 60:
		return "D"
	default:
		return "F"
	}
}

func (s *Student) Birthday() {
	s.Age++
}

func (s *Student) AddScore(score int) bool {
	if score < 0 || score > 100 {
		return false
	}

	s.Scores = append(s.Scores, score)
	return true
}

func (item CartItem) Subtotal() int {
	return item.Product.Price * item.Quantity
}

func (account *BankAccount) Deposit(amount int) bool {
	if amount <= 0 {
		return false
	}

	account.balance += amount
	return true
}

func (account *BankAccount) Withdraw(amount int) bool {
	if amount <= 0 || amount > account.balance {
		return false
	}

	account.balance -= amount
	return true
}

func (account BankAccount) Balance() int {
	return account.balance
}
