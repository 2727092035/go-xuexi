// Day 4 练习题
// 完成下面的练习，验证你对数组、切片、append 与 map 的理解。
// 运行：go run phase1-basics/day04/exercise/main.go

package main

import "fmt"

func main() {
	fmt.Println("练习 1：分数统计")
	// 要求：
	// 1. 用 slice 保存一组分数：88、92、76、100、67。
	// 2. 实现 scoreStats(scores []int) (sum int, avg float64, max int, min int)。
	// 3. 打印总分、平均分、最高分、最低分。
	// 4. 空切片返回 0, 0, 0, 0。
	// 在这里写代码 👇
	scores := []int{88, 92, 76, 100, 67}
	sum, avg, max, min := scoreStats(scores)
	fmt.Println(sum, avg, max, min)

	fmt.Println("\n练习 2：append 构建切片")
	// 要求：
	// 1. 从空切片开始，用 append 加入 1 到 10。
	// 2. 打印每次 append 后的 len 和 cap。
	// 3. 最后打印完整切片。
	// 在这里写代码 👇
	slice := []int{}
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Println(len(slice), cap(slice))
	}
	fmt.Println(slice)

	fmt.Println("\n练习 3：字符计数")
	// 要求：
	// 1. 实现 charCount(s string) map[rune]int。
	// 2. 给定 "hello, 世界"，统计每个字符出现次数。
	// 3. 用 range 遍历结果并打印。
	// 提醒：map 遍历顺序不固定。
	// 在这里写代码 👇
	s := "hello, 世界"
	count := charCount(s)
	for k, v := range count {
		fmt.Printf("%q %d\n", k, v)
	}

	fmt.Println("\n练习 4：商品价格表")
	// 要求：
	// 1. 用 map[string]int 保存商品价格：apple=5、banana=3、orange=4。
	// 2. 查询 banana，打印价格和是否存在。
	// 3. 修改 banana 价格为 6。
	// 4. 删除 apple。
	// 5. 查询 pear，打印价格和是否存在。
	// 在这里写代码 👇

	prices := map[string]int{"apple": 5, "banana": 3, "orange": 4}
	price, ok := prices["banana"]
	fmt.Println(price, ok)
	prices["banana"] = 6
	delete(prices, "apple")
	price, ok = prices["pear"]
	fmt.Println(price, ok)

	fmt.Println("\n加餐 1：整数去重")
	// 要求：
	// 1. 实现 unique(nums []int) []int。
	// 2. 输入 []int{3, 1, 3, 2, 1, 5}，输出去重后的结果。
	// 3. 保持第一次出现的顺序。
	// 在这里写代码 👇

	nums := []int{3, 1, 3, 2, 1, 5}
	unique := unique(nums)
	fmt.Println(unique)

	fmt.Println("\n加餐 2：最高分学生")
	// 要求：
	// 1. 实现 topScore(students map[string]int) (name string, score int)。
	// 2. 输入 map[string]int{"Tom": 80, "Lucy": 95, "Jack": 88}。
	// 3. 输出最高分学生姓名和分数。
	// 4. 空 map 返回 "", 0。
	// 在这里写代码 👇

	students := map[string]int{"Tom": 80, "Lucy": 95, "Jack": 88}
	name, score := topScore(students)
	fmt.Println(name, score)

	fmt.Println("\n加餐 3：简单购物车")
	// 要求：
	// 1. prices := map[string]int{"apple": 5, "banana": 3, "orange": 4}
	// 2. cart := map[string]int{"apple": 2, "orange": 3}
	// 3. 实现 cartTotal(prices, cart map[string]int) int。
	// 4. 只计算价格表里存在的商品。
	// 在这里写代码 👇

	cartPrices := map[string]int{"apple": 5, "banana": 3, "orange": 4}
	cart := map[string]int{"apple": 2, "orange": 3}
	total := cartTotal(cartPrices, cart)
	fmt.Println(total)

	fmt.Println("\n练习完成！")
}

// 在下面声明练习函数，先写签名，再补逻辑。

func scoreStats(scores []int) (sum int, avg float64, max int, min int) {
	if len(scores) == 0 {
		return
	}

	max = scores[0]
	min = scores[0]
	for _, score := range scores {
		sum += score
		if score > max {
			max = score
		} else if score < min {
			min = score
		}
	}
	avg = float64(sum) / float64(len(scores))
	return sum, avg, max, min
}

func charCount(s string) map[rune]int {
	count := map[rune]int{}
	for _, r := range s {
		count[r]++
	}
	return count
}

func contains(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}

func unique(nums []int) []int {
	// TODO
	unique := []int{}
	for _, num := range nums {
		if !contains(unique, num) {
			unique = append(unique, num)
		}
	}
	return unique
}

func topScore(students map[string]int) (name string, score int) {
	for studentName, studentScore := range students {
		if name == "" || studentScore > score {
			name = studentName
			score = studentScore
		}
	}
	return
}

func cartTotal(prices, cart map[string]int) int {
	total := 0
	for product, quantity := range cart {
		price, ok := prices[product]
		if ok {
			total += price * quantity
		}
	}
	return total
}
