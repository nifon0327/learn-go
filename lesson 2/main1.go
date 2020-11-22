package main

import "fmt"

func main() {
	// 常量 const
	const constVariables1 float64 = 3.141592653589793238462643383279502816939937

	const constVariables2, constVariables3 = 100, "波哥"

	fmt.Println(constVariables1)                  // 3.14159265358973
	fmt.Println(constVariables2, constVariables3) // 100 波哥

	const (
		iotaVariables1 = iota // 0
		iotaVariables2 = iota // 1
		iotaVariables3 = iota // 2
	)
	const iotaVariables4 = iota // 0
	fmt.Println(iotaVariables1, iotaVariables2, iotaVariables3, iotaVariables4)

	const (
		iotaVariables5 = iota
		iotaVariables6
		iotaVariables7
	)

	fmt.Println(iotaVariables5, iotaVariables6, iotaVariables7) // 0 1 2

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		iotaVariables8, iotaVariables9, iotaVariables10 = iota, iota, iota
	)

	fmt.Println(iotaVariables8,iotaVariables9,iotaVariables10)
	const (
		iotaVariables11 = iota
		iotaVariables12 = "123"
		iotaVariables13 = iota
	)

	fmt.Println(iotaVariables11,iotaVariables12,iotaVariables13)
}
