package main

import "fmt"

func main() {
	// 十进制
	var floatVariables1 float32 = 3.14
	floatVariables2 := 3.14

	fmt.Printf("floatVariables1的类型%T,值%v\n",floatVariables1,floatVariables1)
	fmt.Printf("floatVariables2的类型%T,值%v\n",floatVariables2,floatVariables2)

	// 科学计数法
	var floatVariables3 float32 = 3.1415926e2
	var floatVariables4 float64 = 3.1415926e-2
	fmt.Printf("floatVariables3的类型%T,值%v\n",floatVariables3,floatVariables3)
	fmt.Printf("floatVariables4的类型%T,值%v\n",floatVariables4,floatVariables4)

	floatVariables3 = float32(floatVariables4)

	// 复数
	var complexVariables complex64
	complexVariables = 3.14+12i
	complexVariables1 := complex(3.14 ,12)

	fmt.Printf("complexVariables的类型%T,值=%v\n",complexVariables,complexVariables)
	fmt.Printf("complexVariables1的类型%T,值=%v\n",complexVariables1,complexVariables1)

	fmt.Println(real(complexVariables),imag(complexVariables))
}