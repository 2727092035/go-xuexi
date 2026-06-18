// Day 1 练习题
// 完成下面的练习，验证你对变量和类型的理解

package main

import "fmt"

func main() {
	// 练习 1：个人信息卡片
	// 用短声明创建以下变量并打印：
	// - 你的名字（string）
	// - 年龄（int）
	// - 身高，单位米（float64）
	// - 是否会 Go 语言（bool）
	// 用 Printf 格式化输出："我叫XX，今年XX岁，身高XX米，会Go：XX"

	// 在这里写代码 👇
	name := "panxin"
	age := 26
	height := 183.5
	isGo := false
	// fmt.Printf("打印 \n",name,age,height,isGo)
	fmt.Printf("我叫%s，今年%d岁，身高%.1f米，会Go：%t\n", name, age, height, isGo)

	// 练习 2：类型转换
	// 给定温度 celsius = 36.6（float64）
	// 转换为华氏度：fahrenheit = celsius * 9/5 + 32
	// 再将 fahrenheit 转为 int（截断小数）
	// 打印三个值

	// 在这里写代码 👇

	celsius := 36.6
	fahrenheit := celsius*9/5 + 32
	fahrenheitInt := int(fahrenheit)
	fmt.Printf("%.1f摄氏度 = %.1f华氏度 = %d\n", celsius, fahrenheit, fahrenheitInt)

	// 练习 3：零值探索
	// 声明以下变量（只声明不赋值），打印它们的零值：
	// - var a int
	// - var b string
	// - var c bool
	// - var d float64
	// 用 Printf 的 %v 和 %T 分别打印值和类型
	// 思考：和 JS 的 undefined 有什么区别？

	// 在这里写代码 👇
	var a int
	var b string
	var c bool
	var d float64
	fmt.Printf("a: %d, type: %T\n", a, a)
	fmt.Printf("b: %s, type: %T\n", b, b)
	fmt.Printf("c: %t, type: %T\n", c, c)
	fmt.Printf("d: %f, type: %T\n", d, d)

	// 练习 4：常量计算
	// 定义圆的半径 radius = 5（常量）
	// 定义 pi = 3.14159（常量）
	// 计算并打印圆的面积（pi * radius * radius）和周长（2 * pi * radius）

	// 在这里写代码 👇
	const radius = 5
	const pi = 3.14159
	var area float64 = pi * radius * radius
	var circumference float64 = 2 * pi * radius
	fmt.Printf("圆的面积: %f, 圆的周长: %f\n", area, circumference)

	fmt.Println("\n练习完成！")
}
