package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	x := generateValue()
	y := generateValue()

	sum := x + y

	fmt.Println(sum)
}

func generateValue() int {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	return randN.Intn(10)
}
