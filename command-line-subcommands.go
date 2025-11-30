package main

// import (
// 	"fmt"
// 	"os"
// 	"flag"
// )


// func main(){
// 	fooCmd := flag.NewFlagSet("foo",flag.ExitOnError)
// 	fooEnable := fooCmd.Bool("enable",false,"enable")
// 	fooName := fooCmd.String("name","","name")

// 	barCmd := flag.NewFlagSet("bar",flag.ExitOnError)
// 	barLevel := barCmd.Int("level",0,"level")

// 	if len(os.Args) < 2{
// 		fmt.Println("expected 'foo' or 'bar' subcommands")
// 		os.Exit(1)
// 	}

// 	switch os.Args[1]{

// 	case "foo":
// 		fooCmd.Parse(os.Args[2:])
// 		fmt.Println("subcommand 'foo'")
// 		fmt.Println("  enable:",*fooEnable)
// 		fmt.Println("  name:",*fooName)
// 		fmt.Println("  tail:",fooCmd.Args())
// 	case "bar":
// 		barCmd.Parse(os.Args[2:])
// 		fmt.Println("subcommand 'bar'")
// 		fmt.Println("  level:",*barLevel)
// 		fmt.Println("  tail:",barCmd.Args())
// 	default:
// 		fmt.Println("expect 'foo' or 'bar' subcommands")
// 		os.Exit(1)
// 	}
// }





// // | 值                      | 含义                |
// // | ---------------------- | ----------------- |
// // | `flag.ContinueOnError` | 出错时返回 error，你自己处理 |
// // | `flag.ExitOnError`     | 出错时打印信息并退出（常用）    |
// // | `flag.PanicOnError`    | 出错直接 panic        |
