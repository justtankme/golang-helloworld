package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
}

// 在这个例子中，`add` 接受两个 int 类型的参数。
// 注意类型在变量名 _之后_。
// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add(x, y int) int {
	return x + y
}
