package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	// 首页
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "<h1>欢迎来到我的网站</h1>")
// 		w.Header().Set("test","xianwei")
// 		fmt.Fprintf(w, `<a href="/about">关于我们</a> | <a href="/contact">联系我们</a>`)
// 	})

// 	// 关于页面
// 	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "<h1>关于我们</h1>")
// 		fmt.Fprintf(w, "<p>我们是一个很棒的公司！</p>")
// 		fmt.Fprintf(w, `<a href="/">返回首页</a>`)
// 	})

// 	// 联系页面
// 	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "<h1>联系我们</h1>")
// 		fmt.Fprintf(w, "<p>电话：123-456-7890</p>")
// 		fmt.Fprintf(w, `<a href="/">返回首页</a>`)
// 	})
// 	fmt.Println("服务器启动在 http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }



