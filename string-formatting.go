package main

// import (
// 	"fmt"
// 	"os"
// )

// type point struct {
// 	x, y int
// }

// func main() {

// 	p := point{1, 2}
// 	fmt.Printf("struct1: %v\n", p)

// 	fmt.Printf("struct2: %+v\n", p)

// 	fmt.Printf("struct3: %#v\n", p)

// 	fmt.Printf("type: %T\n", p)

// 	fmt.Printf("bool: %t\n", true)

// 	fmt.Printf("int: %d\n", 123)

// 	fmt.Printf("bin: %b\n", 14)

// 	fmt.Printf("char: %c\n", 33)

// 	fmt.Printf("hex: %x\n", 456)

// 	fmt.Printf("float1: %f\n", 78.9)

// 	fmt.Printf("float2: %e\n", 123400000.0)
// 	fmt.Printf("float3: %E\n", 123400000.0)

// 	fmt.Printf("str1: %s\n", "\"string\"")

// 	fmt.Printf("str2: %q\n", "\"string\"")

// 	fmt.Printf("str3: %x\n", "hex this")

// 	fmt.Printf("pointer: %p\n", &p)

// 	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

// 	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

// 	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

// 	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

// 	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

// 	s := fmt.Sprintf("sprintf: a %s", "string")
// 	fmt.Println(s)

// 	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
// }

// // // ======== 基本类型 ========

// // // 通用
// // %v      // 值的默认格式
// // %+v     // 显示结构体字段名和值
// // %#v     // Go 语法格式（可直接复制的代码形式）
// // %T      // 显示类型
// // %%      // 字面量百分号

// // // ======== 布尔值 ========
// // %t      // true 或 false

// // // ======== 整数 ========
// // %d      // 十进制整数
// // %b      // 二进制
// // %o      // 八进制
// // %x      // 十六进制（小写）
// // %X      // 十六进制（大写）
// // %c      // 显示对应的 Unicode 字符
// // %U      // Unicode 格式 U+1234

// // // ======== 浮点数与科学计数 ========
// // %f      // 小数形式（默认 6 位小数）
// // %.2f    // 指定小数位（例：保留 2 位）
// // %e      // 科学计数法（小写 e）
// // %E      // 科学计数法（大写 E）
// // %g      // 根据情况自动在 %e 和 %f 之间选择
// // %G      // 大写版本

// // // ======== 字符串 ========
// // %s      // 字符串
// // %q      // 带双引号的字符串（可转义）
// // %x      // 以十六进制输出字符串
// // % X     // 十六进制并用空格分隔（注意空格）

// // // ======== 指针 ========
// // %p      // 指针地址（十六进制）

// // // ======== 宽度与对齐 ========
// // %6d     // 右对齐，总宽度 6
// // %-6d    // 左对齐，总宽度 6
// // %06d    // 零填充，总宽度 6

// // // ======== 字符串/浮点宽度控制 ========
// // %8s     // 字符串宽度 8，右对齐
// // %-8s    // 左对齐
// // %.5s    // 截断字符串（只显示前 5 个字符）
// // %8.2f   // 总宽度 8，小数 2 位
