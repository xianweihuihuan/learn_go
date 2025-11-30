package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func test(i int) {
// 	fmt.Printf("Test %d\n", i)
// }

// func main() {
// 	i := rand.Intn(3)
// 	fmt.Print("write ",i," as ")
// 	switch i{
// 	case 0:
// 		fmt.Println("zero")
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Saturday,time.Sunday:
// 		fmt.Println("It's weekend")
// 	default:
// 		fmt.Println("It's weekday")
// 	}

// 	t := time.Now()
// 	switch{
// 	case t.Hour()<12:
// 		fmt.Println("It's before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}

// 	whatAmI := func(i interface{}){
// 		switch t := i.(type){
// 		case bool:
// 			fmt.Println("I am a bool")
// 		case int:
// 			fmt.Println("I am an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n",t)
// 		}
// 	}

// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("strdwadwa")

// 	switch tmp := rand.Int31n(36);tmp{
// 	case 0:
// 		fmt.Println("zero")
// 	default:
// 		fmt.Println("default")
// 	}
// }

