package main

import "fmt"

func main() {
	const constVariables1 float64 = 3.141592653589793238462643383279502816939937
	const constVariables2, constVariables3 = 200, "波哥"

	const (
		iotaVariables4 = iota
		iotaVariables5 = iota
		iotaVariables6 = iota
	)

	const (
		iotaVariables7 = iota
		iotaVariables8
		iotaVariables9
	)

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	const (
		iotaVariables10, iotaVariables11, iotaVariables12 = iota, iota, iota
	)

	const (
		iotaVariables13 = iota
		iotaVariables14 = "boob"
		iotaVariables15 = iota
	)

	fmt.Println(constVariables1) // 3.141592653589793
	fmt.Println(constVariables2, constVariables3) // 200 波哥
	fmt.Println(iotaVariables4, iotaVariables5, iotaVariables6) // 0 1 2
	fmt.Println(iotaVariables7, iotaVariables8, iotaVariables9) // 0 1 2
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) // 0 1 2 3 4 5 6
	fmt.Println(iotaVariables10, iotaVariables11, iotaVariables12) // 0 0 0
	fmt.Println(iotaVariables13, iotaVariables14, iotaVariables15) // 0 boob 2

}
