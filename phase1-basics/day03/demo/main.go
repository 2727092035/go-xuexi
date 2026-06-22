// Day 3 示例：函数、作用域、多返回值
// 运行：go run phase1-basics/day03/demo/main.go

package main

import "fmt"

func main() {
	fmt.Println("===== 1. 基本函数 =====")

	score := 86
	level := grade(score)
	fmt.Printf("分数 %d 的等级是 %s\n", score, level)

	fmt.Println("\n===== 2. 参数和返回值 =====")

	celsius := 36.6
	fahrenheitValue, fahrenheitInt := fahrenheit(celsius)
	fmt.Printf("%.1f 摄氏度 = %.1f 华氏度，截断后是 %d\n", celsius, fahrenheitValue, fahrenheitInt)

	fmt.Println("\n===== 3. 多返回值 =====")

	nums := []int{3, 7, 2, 9, 4, 9, 1}
	max, count := maxAndCount(nums)
	fmt.Printf("最大值是 %d，出现 %d 次\n", max, count)

	letters, digits, others := countChars("Hello, 世界! 123")
	fmt.Printf("字母: %d，数字: %d，其他: %d\n", letters, digits, others)

	fmt.Println("\n===== 4. 作用域 =====")

	name := "Tom"
	if age := 18; age >= 18 {
		fmt.Printf("%s 已成年，年龄是 %d\n", name, age)
	}
	// 这里不能访问 age，因为 age 只在 if 代码块里有效。
	// fmt.Println(age) // 取消注释会编译失败

	fmt.Println("\n===== 5. 命名返回值 =====")

	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形面积: %d，周长: %d\n", area, perimeter)

	fmt.Println("\n===== 6. 函数拆分后的 FizzBuzz =====")

	for i := 1; i <= 20; i++ {
		fmt.Print(fizzBuzz(i), " ")
	}
	fmt.Println()
}

// grade 根据分数返回等级。
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

// fahrenheit 返回华氏温度的小数值和截断后的整数值。
func fahrenheit(celsius float64) (float64, int) {
	value := celsius*9/5 + 32
	return value, int(value)
}

// maxAndCount 返回切片中的最大值和最大值出现次数。
func maxAndCount(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}

	max := nums[0]
	count := 0
	for _, n := range nums {
		switch {
		case n > max:
			max = n
			count = 1
		case n == max:
			count++
		}
	}

	return max, count
}

// countChars 统计字符串中的字母、数字和其他字符数量。
func countChars(s string) (letters int, digits int, others int) {
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			digits++
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '一' && r <= '鿿'):
			letters++
		default:
			others++
		}
	}

	return letters, digits, others
}

// rectangle 演示命名返回值。
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

// fizzBuzz 把 Day 2 写在 main 里的逻辑拆成函数。
func fizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
