package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 整形，数值类型
	/*
		int8	1		-2^7~2^7-1
		int16	2		-2^15~2^15-1
		int32	4		-2^31~2^31-1
		int63	8		-2^7~2^63-1

		uint8	1		2^8
		uint16	2		2^16
		uint32 	4 		2^32
		uint64	8		2^64
	*/

	var intVariables1 = 100
	intVariables2 := 200
	var intVariables3 int32
	intVariables := 126

	intVariables3 = int32(intVariables)

	var intVariables4 int64 = 123456789
	fmt.Printf("intVariables1=%T,intVariables2=%T,intVariables=%T\n", intVariables1, intVariables2, intVariables3)
	fmt.Printf("%T,%T,%T\n", intVariables, int32(intVariables), intVariables3)
	fmt.Println(unsafe.Sizeof(intVariables4))

}
