package main

// import (
// 	"bytes"
// 	"fmt"
// 	//"bytes"
// 	"regexp"
// )

// func main(){
// 	match,_ := regexp.MatchString("p([a-z]+)ch","peach")
// 	fmt.Println(match)

// 	r,_ := regexp.Compile("p([a-z]+)ch")

// 	fmt.Println(r.MatchString("peach"))

// 	fmt.Println(r.FindString("peach punch"))

// 	fmt.Println("idx:",r.FindStringIndex("peach punch"))

// 	fmt.Println(r.FindStringSubmatch("peach punch"))

// 	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	
// 	fmt.Println(r.FindAllString("peach punch pinch",-1))

// 	fmt.Println("all:",r.FindAllStringSubmatchIndex("peach punch pinch",-1))

// 	fmt.Println(r.FindAllString("peach punch pinch",2))

// 	fmt.Println(r.Match([]byte("peach")))

	
// 	r = regexp.MustCompile("p([a-z]+)ch")

// 	fmt.Println("regexp:",r)

// 	fmt.Println(r.ReplaceAllString("a peach","<fruit>"))

// 	in := []byte("a peach")
// 	out := r.ReplaceAllFunc(in,bytes.ToUpper)

// 	fmt.Println(string(in),"->",string(out))
// }


// // =====================
// // 基础字符匹配
// // =====================
// // .           匹配任意单个字符（不包括换行）
// // \w          字母、数字、下划线
// // \W          非 \w
// // \d          数字 [0-9]
// // \D          非数字
// // \s          空白字符（空格、Tab、换行）
// // \S          非空白字符

// // =====================
// // 字符集（Character Class）
// // =====================
// // [abc]       匹配 a 或 b 或 c
// // [^abc]      匹配非 a/b/c
// // [a-z]       匹配 a 到 z
// // [A-Z]       匹配 A 到 Z
// // [0-9]       匹配数字
// // [\w\d]      组合使用

// // =====================
// // 数量词（Quantifiers）
// // =====================
// // a*          0 次或多次
// // a+          1 次或多次
// // a?          0 次或 1 次
// // a{n}        正好 n 次
// // a{n,}       至少 n 次
// // a{n,m}      n 到 m 次（包含 n 和 m）
// // a+?         非贪婪匹配（RE2 支持）

// // =====================
// // 锚点（Anchors）
// // =====================
// // ^           输入开头
// // $           输入结尾
// // \b          单词边界
// // \B          非单词边界

// // =====================
// // 分组与引用
// // =====================
// // (group)         捕获分组
// // (?:group)       【不支持：Go/RE2 不支持非捕获组】
// // (?P<name>...)   【不支持：Go 不支持命名捕获】

// // =====================
// // 替代（OR）
// // =====================
// // a|b         匹配 a 或 b

// // =====================
// // 转义字符
// // =====================
// // \.          匹配 .
// // \*          匹配 *
// // \+          匹配 +
// // \?          匹配 ?
// // \|          匹配 |
// // \\          匹配反斜杠 \
// // \^          匹配 ^
// // \$          匹配 $

// // =====================
// // 特殊语法（RE2 支持）
// // =====================
// // (?i)        忽略大小写
// // (?m)        多行模式（^ 和 $ 作用于每行）
// // (?s)        让 . 匹配换行符
// // (?U)        禁止贪婪模式（默认成为非贪婪）
// // (?i:expr)   表达式 expr 局部忽略大小写

// // =====================
// // Go / RE2 不支持的（重要）
// // =====================
// // (?=...)     正向肯定预查        <-- 不支持
// // (?!...)     正向否定预查        <-- 不支持
// // (?<=...)    反向肯定预查        <-- 不支持
// // (?<!...)    反向否定预查        <-- 不支持
// // (?(cond)...) 条件匹配          <-- 不支持
