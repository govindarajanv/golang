package main

import "fmt"

//package level
var x,y bool
var a,b int = 1,2

func main(){

	//function level
	var i int
	j := 5
	fmt.Println(i,j,x,y,a,b)
}
