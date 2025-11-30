package main

// import (
// 	"os"
// 	"os/exec"
// 	"syscall"
// )

// func main() {

// 	binary, lookErr := exec.LookPath("ls")
// 	if lookErr != nil {
// 		panic(lookErr)
// 	}

// 	args := []string{"ls", "-a", "-l", "-h"}

// 	env := os.Environ()

// 	execErr := syscall.Exec(binary, args, env)
// 	if execErr != nil {
// 		panic(execErr)
// 	}
// }

// //操作系统的exec，会进行程序体的替换
// //Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量。
// //Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。 我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。