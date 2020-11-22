package main

import "fmt"

func main () {
	// 常量 const
	const constVariables1 float64 = 3.14159265358979323846264338327950288419716939937
	//const intVariabels1 int64 = 314159265358979323846264338327950288419716939937 溢出
	fmt.Println(constVariables1)
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
	const constVariables2,constVariable3 = 100,"波哥"

	const (
		iotaVariables1 = iota
		iotaVariables2 = iota
		iotaVariables3 = iota
	)

	const iotaVariables4 = iota
	fmt.Println(iotaVariables1,iotaVariables2,iotaVariables3,iotaVariables4)

	const (
		iotaVariables5 = iota
		iotaVariables6
		iotaVariables7
	)
	fmt.Println(iotaVariables5,iotaVariables6,iotaVariables7)

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday)

	const (
		iotaVariables8,iotaVariables9,iotaVariables10 = iota,iota,iota
	)

	fmt.Println(iotaVariables8,iotaVariables9,iotaVariables10)

	const (
		iotaVariables11 = iota
		iotaVariables12 = "BoBo"
		iotaVariables13 = iota
	)
	fmt.Println(iotaVariables11,iotaVariables12,iotaVariables13)





}
