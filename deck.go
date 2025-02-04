package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Heart", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, Suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+Suit)
		}
	}
	return cards
}
func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
