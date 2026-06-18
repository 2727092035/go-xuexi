// Day 2 练习题
// 完成下面的练习，验证你对条件与循环的理解
// 运行：go run exercises.go

package main

import "fmt"

func main() {
	// 练习 1：成绩等级（if-else）
	// 给定 score = 78
	// 用 if-else 判断等级并打印：
	// >= 90 → "A"，>= 80 → "B"，>= 70 → "C"，>= 60 → "D"，否则 "F"
	// 在这里写代码 👇
	fmt.Println("练习 1：成绩等级")
	const score = 78
	if score >= 90 {
		fmt.Print("A")
	} else if score >= 80 {
		fmt.Print("B")
	} else if score >= 70 {
		fmt.Print("C")
	} else if score >= 60 {
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
	fmt.Println()

	// 练习 2：FizzBuzz（for + if + switch 任选）
	// 打印 1 到 20：
	// - 能被 3 整除 → "Fizz"
	// - 能被 5 整除 → "Buzz"
	// - 同时能被 3 和 5 整除 → "FizzBuzz"
	// - 其他 → 数字本身
	// 在这里写代码 👇
	fmt.Println("练习 2：FizzBuzz")
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()

	// 练习 3：九九乘法表（嵌套 for）
	// 用两层 for 循环打印九九乘法表
	// 提示：内层用 fmt.Printf("%d*%d=%d\t", ...)，每行结束 fmt.Println()
	// 在这里写代码 👇
	fmt.Println("练习 3：九九乘法表")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习 4：统计字符串（range）
	// 给定 s := "Hello, 世界! 123"
	// 用 range 遍历，统计：字母数量、数字数量、其他字符数量
	// 提示：字符 r 是 rune 类型，可以用 r >= '0' && r <= '9' 判断数字
	// 在这里写代码 👇
	fmt.Println("练习 4：统计字符串")
	s := "Hello, 世界! 123"
	letterCount := 0
	digitCount := 0
	otherCount := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digitCount++
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '一' && r <= '鿿') {
			letterCount++
		} else {
			otherCount++
		}
	}

	fmt.Println("字母数量:", letterCount)
	fmt.Println("数字数量:", digitCount)
	fmt.Println("其他字符数量:", otherCount)

	// 练习 5：找最大值（for + range + 条件）
	// 给定 nums := []int{3, 7, 2, 9, 4, 9, 1}
	// 找出最大值并打印；额外：统计最大值出现了几次
	// 在这里写代码 👇
	fmt.Println("练习 5：找最大值")
	nums := []int{3, 7, 2, 9, 4, 9, 1}
	max := nums[0]
	num := 0
	for _, n := range nums {
		switch {
		case n > max:
			max = n
			num = 1
		case n == max:
			num++
		}
	}
	fmt.Printf("最大值为%d，出现了%d次\n", max, num)

	fmt.Println("\n练习完成！")
}
