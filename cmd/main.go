package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to GO DICE!")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(6) + 1)
}
