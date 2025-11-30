package main

// import "fmt"
// import s "strings"

// var p = fmt.Println

// func main() {
// 	p("Contains:  ", s.Contains("test", "es"))
// 	p("Count:     ", s.Count("test", "t"))
// 	p("HasPrefix: ", s.HasPrefix("test", "te"))
// 	p("HasSuffix: ", s.HasSuffix("test", "st"))
// 	p("Index:     ", s.Index("test", "e"))
// 	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
// 	p("Repeat:    ", s.Repeat("a", 5))
// 	p("Replace:   ", s.Replace("foo", "o", "0", -1))
// 	p("Replace:   ", s.Replace("foo", "o", "0", 1))
// 	p("Split:     ", s.Split("a-b-c-d-e", "-"))
// 	p("ToLower:   ", s.ToLower("TEST"))
// 	p("ToUpper:   ", s.ToUpper("test"))
// 	p()

// 	p("Len: ", len("hello"))
// 	p("Char: ", "hello"[1])
// }
// 	// // 是否包含
// 	// fmt.Println(strings.Contains(s, "lo")) // true

// 	// // 是否以某前缀开头
// 	// fmt.Println(strings.HasPrefix(s, "he")) // true

// 	// // 是否以某后缀结尾
// 	// fmt.Println(strings.HasSuffix(s, "ld")) // true

// 	// // 查找子串位置（找不到返回 -1）
// 	// fmt.Println(strings.Index(s, "o")) // 4

// 	// // 按字符串切割
// 	// fmt.Println(strings.Split(s, " ")) // ["hello" "world"]

// 	// // 按空白字符切割（更智能）
// 	// fmt.Println(strings.Fields("  a   b   c  ")) // ["a" "b" "c"]

// 	// // 拼接字符串切片
// 	// fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // "a-b-c"

// 	// // 全部转小写
// 	// fmt.Println(strings.ToLower(s)) // "hello world"

// 	// // 全部转大写
// 	// fmt.Println(strings.ToUpper(s)) // "HELLO WORLD"

// 	// // 替换字符串
// 	// fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // "bbaa"（替换 2 次）
// 	// fmt.Println(strings.ReplaceAll("aaaa", "a", "b")) // "bbbb"（全部替换）

// 	// // 修剪字符串
// 	// fmt.Println(strings.Trim("  hello  ", " "))       // "hello"
// 	// fmt.Println(strings.TrimSpace("  hello world  ")) // "hello world"

// 	// // 统计子串出现次数
// 	// fmt.Println(strings.Count("banana", "na")) // 2

