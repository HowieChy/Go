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
	go generateValue(c)
	go generateValue(c)

	sum := 0
	i := 0

	for num := range c {
		sum += num
		i++
		//读取到四个管道的值表明可以关闭管道了
		if i == 4 {
			close(c)
		}
	}

	fmt.Println(sum)
}

func generateValue(c chan int) int {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := randN.Intn(10)
	c <- result
	return result
}
