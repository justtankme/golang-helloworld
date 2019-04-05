package main

import (
	"fmt"
	"math"
)

// 表达式 T(v) 将值 v 转换为类型 `T`。
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// 与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换
	// 此处会报错：cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	// var f float64 = math.Sqrt((x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}
