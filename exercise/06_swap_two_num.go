package main

import "fmt"

func swap(x,y string) (string,string) {
	fmt.Println("Before Swap:",x,y)
	return y,x
}
func main(){
	a,b := swap("Hello","World")
	fmt.Println("After Swap:",a,b)

}
