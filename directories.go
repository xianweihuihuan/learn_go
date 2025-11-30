package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {
// 	err := os.Mkdir("subdir", 0755)
// 	check(err)

// 	defer os.RemoveAll("subdir")

// 	createEmtyFile := func(name string) {
// 		d := []byte("")
// 		check(os.WriteFile(name, d, 0644))
// 	}

// 	createEmtyFile("subdir/file1")
// 	err = os.MkdirAll("subdir/parent/child", 0755)
// 	check(err)

// 	createEmtyFile("subdir/parent/file2")
// 	createEmtyFile("subdir/parent/file3")
// 	createEmtyFile("subdir/parent/child/file4")

// 	c, err := os.ReadDir("subdir/parent")
// 	for _, entry := range c {
// 		fmt.Println(" ", entry.Name(), entry.IsDir())
// 	}

// 	err = os.Chdir("subdir/parent/child")
// 	check(err)
// 	c, err = os.ReadDir(".")
// 	check(err)

// 	fmt.Println("Listing subdir/parent/child")
// 	for _, entry := range c {
// 		fmt.Println(" ", entry.Name(), entry.IsDir())
// 	}
// 	err = os.Chdir("../../..")
// 	check(err)

// 	fmt.Println("Visiting subdir")
// 	err = filepath.Walk("subdir", visit)
// }

// func visit(p string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(" ", p, info.IsDir())
// 	return nil
// }
