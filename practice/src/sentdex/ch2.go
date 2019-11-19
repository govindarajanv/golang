package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func print() {
	fmt.Println("The square root for 4 is", math.Sqrt(4))
}

// Only main func is always called, other funcs are called only if they are called from main func
func main() {
	print()
	fmt.Println("A number from 1 to 100 is", rand.Intn(100))
	fmt.Println("Now we will be generating a true random number")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("New Random number is", r1.Intn(100))
}
