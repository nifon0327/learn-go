package main

import "fmt"

var (
	//variablesx int
	//slicex []int
	//interfacex interface{}
)

func main(){
	var variabels1 int
	variables2 := 1
	var variables3 = 2

	fmt.Println(variabels1,variables2,variables3)

	var i = 100
	var j = 200
	fmt.Println(i,j)
	i,j = j,i
	fmt.Println(i,j)

	var variables6 float32
	variables6 = 0.5

	fmt.Println(variables6)


	var variables4,variables5,variables8,variabels7 = 1,"aaa",50 / 100,true


	fmt.Println(variables4,variables5,variables6,variabels7,variables8)
}