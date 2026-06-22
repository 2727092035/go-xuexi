// Day 1：变量与类型
// 目标：理解 Go 的变量声明方式和基本数据类型
//
// 在开始之前，先理解每个 Go 文件的固定结构：
//
//   package main    ← 声明这个文件属于哪个包。main 包是程序入口
//   import "fmt"    ← 导入需要的包。fmt 是格式化输出包（format 缩写）
//   func main() {}  ← 程序从这里开始执行，类似 JS 里的立即执行

package main

import "fmt"

func main() {

	// =====================================================
	// 1. 变量声明的三种方式
	// =====================================================
	//
	// 前端只有 let / const / var 三个关键字
	// Go 也有三种声明变量的方式，但语法和规则完全不同

	fmt.Println("===== 1. 变量声明的三种方式 =====")

	// ----- 方式一：短声明 := -----
	// 语法：变量名 := 值
	// 特点：Go 自动推断类型，不需要你写类型
	// 限制：只能在函数内部使用（不能写在函数外面）
	// 这是最常用的方式，80% 的变量都用这种

	name := "Tom"     // Go 看到 "Tom" 是字符串，自动推断 name 是 string
	age := 25         // Go 看到 25 是整数，自动推断 age 是 int
	height := 175.5   // Go 看到 175.5 有小数，自动推断 height 是 float64
	isStudent := true // Go 看到 true，自动推断 isStudent 是 bool

	// 前端对比：
	//   let name = "Tom"     ← JS 也能推断，但 JS 变量可以随时改类型
	//   name := "Tom"        ← Go 推断后类型就固定了，不能改
	//   name = 123           ← ❌ Go 编译报错！string 不能赋 int
	//   name = "Jerry"       ← ✅ 同类型可以重新赋值

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("height:", height)
	fmt.Println("isStudent:", isStudent)

	// ----- 方式二：var 显式声明 -----
	// 语法：var 变量名 类型 = 值
	// 特点：你自己指定类型
	// 适用：需要在函数外部声明时（包级变量），或者想明确指定类型时

	var city string = "Beijing" // 明确告诉 Go：city 是 string
	var score int = 90          // 明确告诉 Go：score 是 int

	// 什么时候用 var 而不是 := ？
	// 1) 在函数外面声明变量时（:= 只能在函数里用）
	// 2) 想声明一个类型但暂时不赋值时（见方式三）
	// 3) 想强制指定一个不同的类型时，比如：
	var smallNum int8 = 100 // 用 := 会推断为 int，但你只想用 int8（省内存）

	fmt.Println("city:", city)
	fmt.Println("score:", score)
	fmt.Println("smallNum:", smallNum)

	// ----- 方式三：只声明不赋值（零值）-----
	// 语法：var 变量名 类型
	// 特点：Go 会给一个默认值，叫「零值」
	// 前端对比：JS 的 let x; 得到 undefined，Go 永远不会是 undefined

	var count int     // 零值：0（整数的默认值是 0）
	var text string   // 零值：""（字符串的默认值是空字符串，不是 null！）
	var flag bool     // 零值：false（布尔的默认值是 false）
	var price float64 // 零值：0（浮点数的默认值是 0）

	fmt.Println("\n--- 零值演示 ---")
	fmt.Println("int 零值:", count)     // 0
	fmt.Println("string 零值:", text)   // （空，什么都不显示）
	fmt.Println("bool 零值:", flag)     // false
	fmt.Println("float64 零值:", price) // 0

	// 为什么 Go 要有零值？
	// JS：let x; → x 是 undefined → x + 1 = NaN → 运行时才发现 bug
	// Go：var x int → x 是 0 → x + 1 = 1 → 永远不会出现 undefined 的问题
	// 这就是 Go 追求的「安全性」：变量永远有一个确定的值

	// ----- 常见错误 -----

	// ❌ 错误 1：var 和 := 混用
	// var x int := 10    // 编译报错！两种方式只能选一种

	// ❌ 错误 2：:= 重复声明同一个变量
	// name := "Tom"
	// name := "Jerry"    // 编译报错！name 已经声明过了
	// name = "Jerry"     // ✅ 用 = 重新赋值就行

	// ❌ 错误 3：声明了变量但没使用
	// var unused int     // 编译报错！Go 不允许有未使用的变量
	//                    // JS 只是 warning，Go 直接不让你编译

	// =====================================================
	// 2. 基本数据类型
	// =====================================================
	//
	// JS 只有 number（不分整数和小数）
	// Go 把数字分得很细：int、int8、int16、int32、int64、float32、float64...
	// 你暂时只需要记住 4 个常用的：int、float64、string、bool

	fmt.Println("\n===== 2. 基本数据类型 =====")

	// ----- 整数 int -----
	// 最常用的整数类型，在 64 位系统上就是 64 位整数
	var num1 int = 42

	// 其他整数类型（了解即可，一般用 int 就够了）：
	var num2 int8 = 127     // 范围 -128 ~ 127（1 字节）
	var num3 int16 = 32767  // 范围 -32768 ~ 32767（2 字节）
	var num4 int64 = 999999 // 范围超大（8 字节）

	// 为什么要分这么多种？—— 省内存
	// 如果你存年龄（0~150），int8 就够了，不需要 int64
	// 但日常开发直接用 int 就行，不用纠结

	fmt.Println("int:", num1)
	fmt.Println("int8:", num2, "（最大 127，再大就溢出了）")
	fmt.Println("int16:", num3)
	fmt.Println("int64:", num4)

	// ----- 浮点数 float64 -----
	// 前端的 number 既能表示整数也能表示小数
	// Go 必须区分：整数是 int，小数是 float64

	var f1 float32 = 3.14              // 32 位浮点，精度低（约 7 位有效数字）
	var f2 float64 = 3.141592653589793 // 64 位浮点，精度高（约 15 位有效数字）
	autoFloat := 3.14                  // := 推断小数默认是 float64

	fmt.Println("float32:", f1)
	fmt.Println("float64:", f2)
	fmt.Println("自动推断:", autoFloat)
	// 结论：小数就用 float64，别用 float32（除非有特殊需求）

	// ----- 字符串 string -----
	var s1 string = "Hello, Go!" // 双引号（和 JS 一样）
	s2 := `这是反引号字符串
可以直接换行
类似 JS 的模板字符串但不支持 ${}`
	// 反引号 `` = 原始字符串，不处理转义符，可以换行
	// 双引号 "" = 普通字符串，支持 \n \t 等转义

	fmt.Println("双引号:", s1)
	fmt.Println("反引号:", s2)

	// 字符串拼接用 +（和 JS 一样）
	greeting := "Hello" + ", " + "World"
	fmt.Println("拼接:", greeting)

	// Go 没有模板字符串 ${}，要拼变量得用 fmt.Sprintf：
	info := fmt.Sprintf("我叫%s，今年%d岁", name, age)
	fmt.Println("格式化拼接:", info)

	// ----- 布尔 bool -----
	var b1 bool = true
	var b2Var bool = false

	fmt.Println("bool:", b1, b2Var)

	// 前端对比：
	// JS 有 truthy/falsy：0、""、null、undefined 都算 false
	// Go 没有这个概念：if 后面必须是 bool，不能写 if 0 或 if ""
	// if 0 { }    // ❌ 编译报错！0 不是 bool
	// if "" { }   // ❌ 编译报错！"" 不是 bool
	// if true { } // ✅ 必须是 bool 值

	// =====================================================
	// 3. 类型转换（Go 不会隐式转换！）
	// =====================================================
	//
	// 这是从 JS 转 Go 最容易踩的坑
	// JS 会偷偷帮你转类型，Go 完全不会

	fmt.Println("\n===== 3. 类型转换 =====")

	// ----- JS 的隐式转换（Go 不会这样做）-----
	// JS:  "5" + 3 = "53"     （数字变字符串）
	// JS:  "5" - 3 = 2        （字符串变数字）
	// JS:  true + 1 = 2       （bool 变数字）
	// Go:  以上全部编译报错！

	// ----- 数字之间的转换 -----
	var intNum int = 10
	var floatNum float64 = float64(intNum) // int → float64：用 float64() 包裹
	var backToInt int = int(floatNum)      // float64 → int：用 int() 包裹，截断小数

	fmt.Println("int 10 → float64:", floatNum) // 10
	fmt.Println("float64 → int:", backToInt)   // 10

	// 截断演示（不是四舍五入！）
	var pi2 float64 = 3.99
	var truncated int = int(pi2)            // 直接砍掉小数，不会变成 4
	fmt.Println("3.99 截断为 int:", truncated) // 3

	// ----- int 和 string 的转换 -----

	// ❌ 常见错误：
	// string(rune(65)) 得到的是 "A"，不是 "65"；直接写 string(65) 会被 go vet 提醒。
	wrong := string(rune(65))
	fmt.Println("string(rune(65)) =", wrong, "← 不是 \"65\"，是 ASCII 字符！")

	// ✅ 正确方式：用 fmt.Sprintf 或 strconv 包
	correct := fmt.Sprintf("%d", 65)    // 数字 → 字符串
	fmt.Println("Sprintf 转换:", correct) // "65"

	// strconv 包的方式（后续会学到，这里先了解）：
	// import "strconv"
	// str := strconv.Itoa(65)       // int → string: "65"
	// num, err := strconv.Atoi("65") // string → int: 65

	// ----- 为什么 Go 不做隐式转换？ -----
	// 因为隐式转换是 bug 的温床：
	//   JS: [] + {} = "[object Object]"   ← 这种行为你能预测吗？
	//   JS: {} + [] = 0                   ← 换个顺序结果还不一样！
	// Go 的哲学：宁可让你多写一个 float64()，也不让你被隐式转换坑到

	// =====================================================
	// 4. 常量 const
	// =====================================================

	fmt.Println("\n===== 4. 常量 =====")

	// 语法：const 名字 = 值
	// 前端对比：和 JS 的 const 类似，声明后不能改

	const pi = 3.14159
	const appName = "MyApp"
	const maxRetry = 3

	// pi = 3.14  // ❌ 编译报错！常量不能重新赋值

	fmt.Println("pi:", pi)
	fmt.Println("appName:", appName)
	fmt.Println("maxRetry:", maxRetry)

	// 批量声明常量（类似枚举）
	const (
		StatusPending  = 0
		StatusActive   = 1
		StatusInactive = 2
	)
	fmt.Println("状态:", StatusPending, StatusActive, StatusInactive)

	// iota：自动递增的常量（Go 特有）
	// 从 0 开始，每一行加 1
	const (
		Sunday    = iota // 0
		Monday           // 1（自动 iota+1）
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println("周日:", Sunday, "周一:", Monday, "周六:", Saturday)
	// 前端里你可能用 enum 或者对象来做，Go 用 iota 更简洁

	// Go 的 const 和 JS 的 const 有什么区别？
	// JS:  const obj = {a: 1}; obj.a = 2  ← ✅ 允许！const 只保证引用不变
	// Go:  const 只能用于基本类型（数字、字符串、布尔）
	//      不能 const 一个 struct 或 slice，因为它们是复杂类型

	// =====================================================
	// 5. fmt 包常用函数（你的 console.log）
	// =====================================================

	fmt.Println("\n===== 5. fmt 输出 =====")

	// ----- Println：最简单的打印 -----
	// 自动在末尾加换行，多个参数用空格隔开
	fmt.Println("Hello", "World", 123) // Hello World 123

	// ----- Printf：格式化打印 -----
	// 需要手动加 \n 换行
	// 用占位符 % 来插入变量（类似 C 语言的 printf）
	fmt.Printf("名字: %s\n", name)

	// 常用占位符一览：
	fmt.Println("\n--- 占位符大全 ---")
	fmt.Printf("%%s  → 字符串:   %s\n", "hello")
	fmt.Printf("%%d  → 整数:     %d\n", 42)
	fmt.Printf("%%f  → 浮点数:   %f\n", 3.14)       // 默认 6 位小数
	fmt.Printf("%%.2f → 保留2位:  %.2f\n", 3.14159) // 3.14
	fmt.Printf("%%t  → 布尔:     %t\n", true)
	fmt.Printf("%%v  → 万能占位: %v %v %v\n", "hi", 42, true) // 什么类型都能打印
	fmt.Printf("%%T  → 打印类型: %T %T %T\n", "hi", 42, 3.14) // string int float64

	// 前端对比：
	// JS:  console.log(`名字: ${name}, 年龄: ${age}`)
	// Go:  fmt.Printf("名字: %s, 年龄: %d\n", name, age)
	//
	// 没有 ${} 模板语法，只能用 %占位符

	// ----- Sprintf：格式化但不打印，返回字符串 -----
	// 和 Printf 一样的语法，但不是打印到控制台，而是返回一个字符串
	result := fmt.Sprintf("我叫%s，今年%d岁", name, age)
	fmt.Println("Sprintf 结果:", result)
	// 使用场景：拼接字符串时用 Sprintf，比 + 拼接更清晰

	// =====================================================
	// 总结对照表
	// =====================================================

	fmt.Println("\n===== 总结 =====")
	fmt.Println("JS let x = 1        →  Go x := 1")
	fmt.Println("JS let x            →  Go var x int（零值 0，不是 undefined）")
	fmt.Println("JS const x = 1      →  Go const x = 1")
	fmt.Println("JS console.log()    →  Go fmt.Println()")
	fmt.Printf("JS `模板${变量}`    →  Go fmt.Sprintf(\"模板%%s\", 变量)\n")
	fmt.Println("JS 隐式类型转换     →  Go 不存在，必须显式转换")
	fmt.Println("JS undefined        →  Go 不存在，每种类型都有零值")
}
