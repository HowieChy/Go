package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//创建一个新的type deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"红心", "方块", "黑桃", "梅花"}
	cardValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+value)
		}
	}
	// fmt.Println(cards)
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//发牌
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//从文件中读取数据
func newDeckFromFile(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//洗牌
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
