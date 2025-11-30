package main

// import "fmt"
// import "time"

// func main(){
// 	time1 := time.NewTimer(2 * time.Second)

// 	<- time1.C
// 	fmt.Println("Timer 1 fired")

// 	time2 := time.NewTimer(5* time.Second)
// 	go func(){
// 		fmt.Println("begain")
// 		<- time2.C
// 		fmt.Println("Timer 2 fired")
// 	}()
// 	time.Sleep(2 * time.Second)
// 	stop2 := time2.Stop()
// 	if stop2{
// 		fmt.Println("Timer 2 stopped")
// 	}

// 	time.Sleep(2* time.Second)
// }