package main

import (
	"fmt"
	"unsafe"
)

// 珠穆朗玛峰 8848
// 乔戈里峰 8611
// 千城章嘉峰 8586
// 洛子峰 8516
// 马卡鲁峰 8463
// 卓奥友峰 8201
// 道拉吉里峰 8172
// 马纳斯鲁峰 8156
// 南伽峰 8125
// 安纳布尔纳峰 8091

func main() {
	// 整型
	/*
		int8	1		-2^7~2^7-1
		int16	2		-2^15~2^15-1
		int32	4		-2^31~2^31-1
		int64	8		-2^63~2^63-1

		uint8	1		2^8
		uint16	2		2^16
		uint32  4		2^32
		uint64  8		2^64
	*/

	var intVariables1 = 100 // int
	intVariables2 := 200    // int
	var intVariables3 int32 // int32
	intVariables := 126     // int

	intVariables3 = int32(intVariables)

	var intVariables4 int64 = 123456789
	fmt.Printf("intVariables1=%T,intVariables2=%T,intVariables3=%T\n", intVariables1, intVariables2, intVariables3)
	fmt.Println(unsafe.Sizeof(intVariables4))
}
