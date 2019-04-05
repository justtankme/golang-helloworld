package main

import (
	"fmt"
	"math"
)

// 在定义一个变量但不指定其类型时（使用没有类型的 var 或 := 语句）， 变量的类型由右值推导得出。
// 当右值定义了类型时，新变量的类型与其相同
// 但是当右边包含了未指名类型的数字常量时，新的变量就可能是 int 、 float64 或 `complex128`。 这取决于常量的精度
func main() {
	v := 42
	v2 := math.Sqrt(3.12)
	fmt.Printf("v,v2 is type of %T %T\n", v, v2)
}
