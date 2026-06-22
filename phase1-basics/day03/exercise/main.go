// Day 3 练习题
// 完成下面的练习，验证你对函数、作用域、多返回值的理解。
// 运行：go run phase1-basics/day03/exercise/main.go

package main

import "fmt"

func main() {
	fmt.Println("练习 1：成绩等级函数")
	// 要求：
	// 1. 实现 grade(score int) string。
	// 2. 用它打印 95、82、76、61、30 的等级。
	// 规则：>=90 A，>=80 B，>=70 C，>=60 D，否则 F。
	// 在这里写代码 👇

	scores := []int{95, 82, 76, 61, 30}
	for _, score := range scores {
		fmt.Println(grade(score))
	}

	fmt.Println("\n练习 2：温度转换函数")
	// 要求：
	// 1. 实现 fahrenheit(celsius float64) (float64, int)。
	// 2. 返回华氏温度小数值，以及截断后的整数值。
	// 3. 测试 36.6 和 0。
	// 公式：fahrenheit = celsius * 9 / 5 + 32。
	// 在这里写代码 👇
	fmt.Println(fahrenheit(36.6))
	fmt.Println(fahrenheit(0))

	fmt.Println("\n练习 3：最大值和出现次数")
	// 要求：
	// 1. 实现 maxAndCount(nums []int) (max int, count int)。
	// 2. 给定 []int{3, 7, 2, 9, 4, 9, 1}，输出最大值和出现次数。
	// 3. 空切片先返回 0, 0。
	// 在这里写代码 👇
	nums := []int{3, 7, 2, 9, 4, 9, 1}
	max, count := maxAndCount(nums)
	fmt.Println(max, count)

	fmt.Println("\n练习 4：字符串统计")
	// 要求：
	// 1. 实现 countChars(s string) (letters int, digits int, others int)。
	// 2. 给定 "Hello, 世界! 123"，输出字母、数字、其他字符数量。
	// 3. 字母判断可以沿用 Day 2 的写法。
	// 在这里写代码 👇
	letters, digits, others := countChars("Hello, 世界! 123")
	fmt.Println(letters, digits, others)

	fmt.Println("\n练习 5：FizzBuzz 函数拆分")
	// 要求：
	// 1. 实现 fizzBuzz(n int) string。
	// 2. 用 for 循环打印 1 到 30 的结果。
	// 3. main 里只负责循环和打印，判断逻辑必须放到函数里。
	// 在这里写代码 👇
	for i := 1; i <= 30; i++ {
		fmt.Println(fizzBuzz(i))
	}

	fmt.Println("\n加餐 1：质数判断")
	// 要求：
	// 1. 实现 isPrime(n int) bool。
	// 2. 打印 1 到 100 之间所有质数。
	// 提示：小于 2 不是质数；只需要检查 2 到 i*i <= n。
	// 在这里写代码 👇
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	fmt.Println("\n加餐 2：区间求和")
	// 要求：
	// 1. 实现 sumRange(start, end int) int。
	// 2. 支持 start > end 时自动交换。
	// 3. 测试 sumRange(1, 100) 和 sumRange(10, 1)。
	// 在这里写代码 👇

	fmt.Println(sumRange(1, 100))
	fmt.Println(sumRange(10, 1))
	fmt.Println("\n练习完成！")
}

// 在下面声明练习函数（只写签名，逻辑由你自己补全）。

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func fahrenheit(celsius float64) (float64, int) {
	value := celsius*9/5 + 32
	return value, int(value)
}

func maxAndCount(nums []int) (max int, count int) {
	if len(nums) == 0 {
		return 0, 0
	}

	max = nums[0]
	count = 0
	for _, num := range nums {
		if num > max {
			max = num
			count = 1
		} else if num == max {
			count++
		}
	}
	return max, count
}

func countChars(s string) (letters int, digits int, others int) {
	letters = 0
	digits = 0
	others = 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits++
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '一' && r <= '鿿') {
			letters++
		} else {
			others++
		}
	}
	return letters, digits, others
}

func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumRange(start, end int) int {
	if start > end {
		start, end = end, start
	}
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}
