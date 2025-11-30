package main


// import (
// 	"fmt"
// 	"net"
// 	"net/url"
// )

// func main(){
// 	s := "postgres://user:pass@host.com:5432/path?k=v#f"
// 	u,err := url.Parse(s)
// 	if err != nil{
// 		panic(err)
// 	}

// 	fmt.Println(u.Scheme)
// 	fmt.Println(u.User)
// 	fmt.Println(u.User.Username())
// 	p,_ := u.User.Password()
// 	fmt.Println(p)

// 	fmt.Println(u.Host)
// 	host,port,_ := net.SplitHostPort(u.Host)

// 	fmt.Println(host)
// 	fmt.Println(port)

// 	fmt.Println(u.Path)
// 	fmt.Println(u.Fragment)

// 	fmt.Println(u.RawQuery)
// 	m,_ := url.ParseQuery(u.RawQuery)
// 	fmt.Println(m)
// 	fmt.Println(m["k"][0])

// }

// // 字段	符号	必须	示例	作用
// // 方案	://	✅	https	协议类型
// // 用户信息	@	❌	user:pass	身份验证
// // 主机名	无	✅	example.com	服务器地址
// // 端口	:	❌	:8080	服务端口
// // 路径	/	❌	/path/file	资源位置
// // 查询	?	❌	?key=value	传递参数
// // 片段	#	❌	#section	页面锚点
