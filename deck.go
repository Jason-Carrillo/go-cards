package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Create a new type of "deck"
//A slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Replace values like indexes with an "_" to tell Go that we aren't going to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Typically calling the reference a shortened version of the receiver d = deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck{
	return bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log error and return a call to newDeck()
		// option #2 - Log error and entirely quite the program
		fmt.Println("Error:", err)

		//This will Close the program it must be ANYTHING except 0
		os.Exit(1)
	}

	s := strings.Split( string(bs), "," )
}