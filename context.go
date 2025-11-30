package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func hello(w http.ResponseWriter, req *http.Request) {

// 	ctx := req.Context()
// 	fmt.Println("server: hello handler started")
// 	defer fmt.Println("server: hello handler ended")

// 	select {
// 	case <-time.After(10 * time.Second):
// 		fmt.Fprintf(w, "hello\n")
// 	case <-ctx.Done():

// 		err := ctx.Err()
// 		fmt.Println("server:", err)
// 		internalError := http.StatusInternalServerError
// 		http.Error(w, err.Error(), internalError)
// 	}
// }

// func main() {

// 	http.HandleFunc("/hello", hello)
// 	http.ListenAndServe(":8090", nil)
// }

// // Go 的 context 用于解决 goroutine 无法协作取消的问题。

// // 主要用途：
// // 1. 取消一组 goroutine（手动或自动）
// // 2. 超时控制
// // 3. 截止时间控制
// // 4. 传递少量元数据（如 trace id）

// // 四种类型：
// // - Background：根 context
// // - TODO：不确定时用
// // - WithCancel：手动取消
// // - WithTimeout：指定时间自动取消
// // - WithDeadline：到时间点自动取消
// // - WithValue：传递 key-value（少用）

// // 核心方法：
// // - ctx.Done()：用于监听取消信号
// // - ctx.Err()：取消原因
// // - ctx.Value()：取值
// // - ctx.Deadline()：获取截止时间

// // 重要原则：
// // - 不要滥用 WithValue
// // - 不要传 nil context
// // - 不要把 context 存起来
// // - 要总是调用 cancel()
// // - goroutine 必须检查 ctx.Done()

// // 典型场景：
// // - HTTP 请求超时
// // - 客户端断开连接自动取消
// // - 数据库、RPC 调用超时
// // - 管理一组 worker 协程
