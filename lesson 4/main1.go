package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var floatVariables1 float32 = 3.1415926535
	floatVariables2 := .1416926
	fmt.Printf("floatVariables1的类型=%T,占用字节大小=%d\n", floatVariables1, unsafe.Sizeof(floatVariables1))
	fmt.Printf("floatVariables2的类型=%T,占用字节大小=%d\n", floatVariables2, unsafe.Sizeof(floatVariables2))

	// 2. 科学计数法
	floatVariables3 := 3.1415926e2
	floatVariables4 := 3.1415926e-2
	fmt.Println(floatVariables3, floatVariables4)

	var floatVariables5 float32 = 3.14
	var floatVariables6 float64 = 3.14
	// Cannot use 'floatVariables5' (type float32) as type float64
	floatVariables6 = float64(floatVariables5)
	fmt.Println(floatVariables6)

	// 复数
	var complexVariables1 complex64
	complexVariables1 = 3.14 + 12i
	complexVariables2 := complex(3.14, 12)

	fmt.Printf("complexVariables1的类型=%T,值=%v\n", complexVariables1, complexVariables1)
	fmt.Printf("complexVariables2的类型=%T,值=%v\n", complexVariables2, complexVariables2)

}
