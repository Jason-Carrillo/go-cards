package main

import "fmt"

//Create a new type of "deck"
//A slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hears", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

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
