// Day 5 示例：结构体、方法、指针接收者
// 运行：go run phase1-basics/day05/demo/main.go

package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Scores []int
}

func main() {
	fmt.Println("===== 1. 结构体 struct =====")

	student := Student{
		Name:   "Tom",
		Age:    18,
		Scores: []int{80, 90, 100},
	}

	fmt.Println("学生:", student)
	fmt.Println("姓名:", student.Name)
	fmt.Println("年龄:", student.Age)
	fmt.Println("分数:", student.Scores)

	fmt.Println("\n===== 2. 函数 vs 方法 =====")

	fmt.Printf("函数计算平均分: %.1f\n", average(student))
	fmt.Printf("方法计算平均分: %.1f\n", student.Average())
	fmt.Println("等级:", student.Level())

	fmt.Println("\n===== 3. 值接收者 =====")

	student.TryBirthday()
	fmt.Println("调用 TryBirthday 后年龄:", student.Age)
	fmt.Println("值接收者拿到的是一份拷贝，改不到原结构体。")

	fmt.Println("\n===== 4. 指针接收者 =====")

	student.Birthday()
	fmt.Println("调用 Birthday 后年龄:", student.Age)
	fmt.Println("指针接收者能修改原结构体。")

	fmt.Println("\n===== 5. 结构体组合业务对象 =====")

	product := Product{Name: "Keyboard", Price: 199}
	item := CartItem{Product: product, Quantity: 2}
	fmt.Println("商品:", product)
	fmt.Println("购物车项小计:", item.Subtotal())
}

func average(s Student) float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

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
	avg := s.Average()
	switch {
	case avg >= 90:
		return "A"
	case avg >= 80:
		return "B"
	case avg >= 70:
		return "C"
	case avg >= 60:
		return "D"
	default:
		return "F"
	}
}

func (s Student) TryBirthday() {
	s.Age++
}

func (s *Student) Birthday() {
	s.Age++
}

type Product struct {
	Name  string
	Price int
}

type CartItem struct {
	Product  Product
	Quantity int
}

func (item CartItem) Subtotal() int {
	return item.Product.Price * item.Quantity
}
