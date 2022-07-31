package main

import (
	"fmt"

	"github.com/howie/firstapp/greeting"
)

func main() {

	fmt.Println(greeting.GreetingText)

	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)
}
