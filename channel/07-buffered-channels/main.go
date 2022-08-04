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
	limiter := make(chan int, 3)

	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)
	go generateValue(c, limiter)

	sum := 0
	i := 0

	for num := range c {
		sum += num
		i++
		if i == 4 {
			close(c)
		}
	}

	fmt.Println(sum)
}

func generateValue(c chan int, limit chan int) int {
	limit <- 1
	fmt.Println("Generating value...")
	// sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(3) * time.Second)

	result := randN.Intn(10)
	c <- result
	<-limit
	return result
}
