package main

import (
	"fmt"
)

// func accept two float arguments and return a float as sum
func add(x float64, y float64) float64 {
	return x + y
}

func addv1(x, y float64) float64 {
	return x + y
}

func multiple(x, y string) (string, string, string) {
	return x, y, x + y
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 11.5
	var num3, num4 float64 = 7.2, 8.7

	fmt.Println(add(num1, num2))
	fmt.Println(addv1(num3, num4))

	str1 := "go"
	str2 := "lang"

	fmt.Println(multiple(str1, str2))
}
