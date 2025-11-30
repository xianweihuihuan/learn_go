package main

// import (
// 	"fmt"
// 	"io"
// 	"os/exec"
// )

// func main() {

// 	dateCmd := exec.Command("date")

// 	dateOut, err := dateCmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("> date")
// 	fmt.Println(string(dateOut))

// 	grepCmd := exec.Command("grep", "hello")

// 	grepIn, _ := grepCmd.StdinPipe()
// 	grepOut, _ := grepCmd.StdoutPipe()
// 	grepCmd.Start()
// 	grepIn.Write([]byte("hello grep\ngoodbye grep"))
// 	grepIn.Close()
// 	grepBytes, _ := io.ReadAll(grepOut)
// 	grepCmd.Wait()

// 	fmt.Println("> grep hello")
// 	fmt.Println(string(grepBytes))

// 	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
// 	lsOut, err := lsCmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("> ls -a -l -h")
// 	fmt.Println(string(lsOut))
// }

// // 注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递一个命令行字符串。 如果你想使用一个字符串生成一个完整的命令，那么你可以使用 bash 命令的 -c 选项：
// // Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。 我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。