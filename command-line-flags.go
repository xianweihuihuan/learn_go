package main

// import (
// 	"fmt"
// 	"flag"
// 	"os"
// )

// func main(){
// 	wordPtr := flag.String("word","foo","a string")

// 	numPtr := flag.Int("numb",42,"an int")

// 	forkPtr := flag.Bool("fork",false,"a bool")

// 	var svar string
// 	flag.StringVar(&svar,"svar","bar","a string var")

// 	flag.Parse()

// 	p := fmt.Println
// 	p("word:",*wordPtr)
// 	p("numb:",*numPtr)
// 	p("fork:",*forkPtr)
// 	p("svar:",svar)
// 	p("tail:",flag.Args())

	
// 	p("total",os.Args)
// }