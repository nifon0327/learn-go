package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 浮点类型 float float64
	// 复数：
	// complex64 32实数+32虚数组成 i虚数的单位
	// complex126 64实数+64虚数组成 虚数单位i
	/*
		float32 4
		float64 8
	*/
	// 1.以十进制的形式来显示
	var floatVariables1 float32 = 3.1415926
	floatVariables2 := .1416926
	fmt.Printf("floatVariables1的类型=%T,占用的字节大小=%d\n", floatVariables1, unsafe.Sizeof(floatVariables1))
	fmt.Printf("floatVariables2的类型=%T,占用的字节大小=%d\n", floatVariables2, unsafe.Sizeof(floatVariables2))

	// 2.以科学计数法来展示
	floatVariables3 := 3.1415926e2
	floatVariables4 := 3.1415926e-2
	fmt.Println(floatVariables3, floatVariables4)

	var floatVariables5 float32 = 3.14
	var floatVariables6 = 3.14
	floatVariables6 = float64(floatVariables5)
	floatVariables6 = floatVariables6
	// 复数 实数"+"虚数i, 虚数的单位
	// complex64 32实数+32虚数组成 i虚数的单位
	// complex126 64实数+64虚数组成 虚数单位i
	var complexVariables1 = 3.14 + 12i
	var complexVariables complex64
	var complexVariables3 complex128
	complexVariables2 := complex(3.14, 32)

	fmt.Printf("complexVariables的类型=%T,complexVariables的值=%v\n", complexVariables, complexVariables)
	fmt.Printf("complexVariables1的类型=%T,complexVariables1的值=%v\n", complexVariables1, complexVariables1)
	fmt.Printf("complexVariables2的类型=%T,complexVariables2的值=%v\n", complexVariables2, complexVariables2)
	fmt.Printf("complexVariables3的类型=%T,complexVariables3的值=%v\n", complexVariables3, complexVariables3)

	fmt.Println(real(complexVariables1), imag(complexVariables1))

}
