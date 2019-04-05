package main

import (
	"fmt"
)

func main() {
	fmt.Println(swap("src", "target"))
}

//函数可以返回任意数量的返回值。
// swap 函数返回了两个字符串。
func swap(src, target string) (string, string) {
	return target, src
}
