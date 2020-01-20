package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(100)
	fmt.Println("Generated random number is",rand.Intn(10))
}

