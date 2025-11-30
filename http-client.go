package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// 	"strings"
// )

// func main() {
// 	client := &http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	req,err := http.NewRequest("GET","https://gobyexample.com",nil)
// 	if err != nil{
// 		panic(err)
// 	}
// 	//req.Header.Set("test","xianwei")

// 	rsp,err := client.Do(req)
// 	if err != nil{
// 		panic(err)
// 	}
// 	defer rsp.Body.Close()
// 	body,_ := io.ReadAll(rsp.Body)
// 	fmt.Println(string(body))
// }

// func main() {
// 	// 最简单的 GET 请求 - 一行代码！
// 	time.Sleep(time.Millisecond)
// 	resp, err := http.Get("https://gobyexample.com")
// 	if err != nil {
// 		fmt.Println("请求失败:", err)
// 		return
// 	}
// 	defer resp.Body.Close() // 记得关闭！

// 	// 读取响应内容
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println("响应内容:", string(body))
// }

// func main() {
// 	// 发送 JSON 数据
// 	jsonData := `{"name": "张三", "age": 25}`
// 	time.Sleep(time.Millisecond)
// 	resp, err := http.Post(
// 		"https://api.example.com/users",
// 		"application/json",
// 		strings.NewReader(jsonData),
// 	)
// 	if err != nil {
// 		fmt.Println("请求失败:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println("创建用户结果:", string(body))
// }


// // Request{
// //     Method          // 请求方法，如 GET、POST
// //     URL             // 请求地址，包含 Path、Query 等
// //     Proto           // 协议版本，如 "HTTP/1.1"
// //     Header          // 请求头 map[string][]string
// //     Body            // 请求体 io.ReadCloser，只能读一次
// //     ContentLength   // Body 长度，未知为 -1
// //     Host            // Host 头，与 URL.Host 可不同
// //     Form            // GET + POST 表单数据（需 ParseForm）
// //     PostForm        // POST 表单数据（需 ParseForm）
// //     RemoteAddr      // 客户端地址，仅服务器端可用
// //     RequestURI      // 原始请求行 URI，客户端一般不用
// //     TLS             // HTTPS 信息

// // }

// // Response{
// //     Status          // "200 OK"
// //     StatusCode      // 200
// //     Proto           // "HTTP/1.1"
// //     Header          // 响应头 map[string][]string
// //     Body            // 响应体 io.ReadCloser，必须 Close()
// //     ContentLength   // 响应体大小，不确定为 -1
// //     TransferEncoding// 若为 chunked 则包含 "chunked"
// //     Trailer         // 尾部 header
// //     Request         // 对应的 Request
// // }
