package main

import (
	"fmt"
	"math"
)

// 在导入了一个包之后，就可以用其导出的名称来调用它。
// 在 Go 中，首字母大写的名称是被导出的。
func main() {
	fmt.Println(math.Pi)
}
