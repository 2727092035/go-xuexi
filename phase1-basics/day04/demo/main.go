// Day 4 示例：数组、切片、append 与 map
// 运行：go run phase1-basics/day04/demo/main.go

package main

import "fmt"

func main() {
	fmt.Println("===== 1. 数组 array =====")

	var fixedScores [3]int = [3]int{80, 90, 100}
	fmt.Println("数组:", fixedScores)
	fmt.Println("数组长度:", len(fixedScores))

	// 数组长度是类型的一部分：[3]int 和 [4]int 是不同类型。
	// 日常业务里更常用 slice，因为长度可以动态变化。

	fmt.Println("\n===== 2. 切片 slice =====")

	scores := []int{78, 90, 66}
	fmt.Println("原始切片:", scores, "len:", len(scores), "cap:", cap(scores))

	scores = append(scores, 88)
	scores = append(scores, 95, 100)
	fmt.Println("append 后:", scores, "len:", len(scores), "cap:", cap(scores))

	total := 0
	for _, score := range scores {
		total += score
	}
	average := float64(total) / float64(len(scores))
	fmt.Printf("总分: %d，平均分: %.1f\n", total, average)

	fmt.Println("\n===== 3. 切片截取 =====")

	firstThree := scores[:3]
	middle := scores[1:4]
	fromTwo := scores[2:]
	fmt.Println("前三个:", firstThree)
	fmt.Println("中间段:", middle)
	fmt.Println("从下标 2 到末尾:", fromTwo)

	fmt.Println("\n===== 4. map 增删改查 =====")

	prices := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	prices["orange"] = 4
	prices["banana"] = 6
	delete(prices, "apple")
	fmt.Println("价格表:", prices)

	price, ok := prices["banana"]
	if ok {
		fmt.Println("banana 价格:", price)
	} else {
		fmt.Println("banana 不存在")
	}

	missingPrice, ok := prices["pear"]
	fmt.Println("pear 查询结果:", missingPrice, "是否存在:", ok)

	fmt.Println("\n===== 5. map 计数 =====")

	text := "hello go"
	counts := map[rune]int{}
	for _, r := range text {
		counts[r]++
	}
	fmt.Println("字符计数:", counts)

	fmt.Println("\n===== 6. map 遍历顺序 =====")

	for name, price := range prices {
		fmt.Printf("%s: %d\n", name, price)
	}
	fmt.Println("注意：map 遍历顺序不能依赖，每次运行都可能不同。")
}
