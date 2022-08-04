package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	c := make(chan int)

	go generateValue(c)
	go generateValue(c)

	x := <-c
	y := <-c

	sum := x + y

	fmt.Println(sum)
}

func generateValue(c chan int) int {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := randN.Intn(10)
	c <- result
	return result
}
