package main

import "fmt"

var (
	variablesx int
	slicex []int
	interfacex interface{}
)
// variables := 100

func main() {

	// 零值
	// int int32 int64
	// float32 float64
	// bool
	// string

	// 定义变量注意事项:
	var variables1 int
	// 简短声明
	variables2 := "我是波哥"

	// _,x := 22,100
	// var index = 1

	var variables3 = 100
	var i = 100
	var j = 200
	fmt.Println(i,j)
	i,j = j,i
	fmt.Println(i,j)
	var variables4,variables5,variables6,variables7 = 1,"我是波哥",3.14,true
	fmt.Println(variables1,variables2,variables3,variables4,variables5,variables6,variables7)
}
