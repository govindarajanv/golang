package main

import "fmt"

func main() {
	x := 25
	b := &x
	fmt.Println("x=", x)
	fmt.Println("b=", b)
	fmt.Println("*b=", *b)
	*b = 20
	fmt.Println("x=", x)
	*b = *b * *b
	fmt.Println("x=", x)
}
