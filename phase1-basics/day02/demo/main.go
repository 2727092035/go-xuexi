// Day 2 示例：条件与循环（if / for / switch / range）
// 运行：go run main.go

package main

import "fmt"

func main() {
	// ========== 1. if ==========
	// 条件不用括号，但大括号必须有
	age := 20
	if age >= 18 {
		fmt.Println("成年")
	} else if age >= 12 {
		fmt.Println("青少年")
	} else {
		fmt.Println("儿童")
	}

	// if 带初始化语句：n 只在 if/else 块内有效
	// 对比 JS：相当于把 const 写进了 if 的作用域里，不污染外层
	if n := len("hello"); n > 3 {
		fmt.Printf("字符串长度 %d，大于 3\n", n)
	}
	// 这里访问 n 会编译报错，因为 n 出了块就消失了

	// ========== 2. for ==========
	// (a) 经典三段式
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("1+2+3+4+5 =", sum)

	// (b) 当 while 用：省略初始化和后置语句
	count := 3
	for count > 0 {
		fmt.Print(count, " ")
		count--
	}
	fmt.Println()

	// (c) 死循环 + break / continue
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // 跳过偶数
		}
		if i > 9 {
			break // 超过 9 就停
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// ========== 3. switch ==========
	// 默认不穿透，每个 case 自动 break
	day := "六"
	switch day {
	case "六", "日": // 多个值用逗号，不用堆叠 case
		fmt.Println("周末")
	default:
		fmt.Println("工作日")
	}

	// 无表达式的 switch，当 if-else 链用，更清晰
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// fallthrough：强制穿透到下一个 case（很少用）
	switch n := 1; n {
	case 1:
		fmt.Println("一")
		fallthrough // 继续执行下一个 case，不再判断条件
	case 2:
		fmt.Println("二（被 fallthrough 带进来的）")
	case 3:
		fmt.Println("三") // 不会执行，fallthrough 只穿透一层
	}

	// ========== 4. range ==========
	// 遍历切片：i 是索引，v 是值
	fruits := []string{"苹果", "香蕉", "橙子"}
	for i, v := range fruits {
		fmt.Printf("[%d] %s\n", i, v)
	}

	// 只要值，索引用 _ 丢弃（声明不用会编译报错）
	total := 0
	for _, v := range []int{10, 20, 30} {
		total += v
	}
	fmt.Println("总和:", total)

	// 遍历 map：k 是键，v 是值（注意 map 遍历顺序是随机的）
	prices := map[string]int{"苹果": 5, "香蕉": 3}
	for k, v := range prices {
		fmt.Printf("%s: %d 元\n", k, v)
	}

	// 遍历字符串：range 按 rune（字符）走，中文也能正确处理
	for i, r := range "Go语言" {
		fmt.Printf("位置 %d: %c\n", i, r)
	}
}
