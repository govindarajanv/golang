package main

import (
	"fmt"
)

func add(x int,y int) int {
	return x+y
}

func subtract(x,y int) int {
	return x-y
}

func main(){
	fmt.Println(add(4,14))
	fmt.Println(subtract(24,14))
}
