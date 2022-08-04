package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	x := make(chan int)
	y := make(chan int)

	go generateValue(x)
	go generateValue(y)

	var a int
	var b int

	//select channel中的类似switch语句
	select {
	case a = <-x:
		fmt.Printf("X finished first, value is %v", a)
	case b = <-y:
		fmt.Printf("Y finished first, value is %v", b)
	}
}

func generateValue(c chan int) int {

	fmt.Println("Generating value...")
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := randN.Intn(10)
	c <- result
	return result
}
