package main

import (
	"fmt"
	"math/rand"
)

// 每个 Go 程序都是由包组成的。
// 程序运行的入口是包 `main`。
func main() {
	fmt.Println("Today's lucky number is:", rand.Intn(10))
}
