package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//BOTH of these are the same
	// card := "Ace of Spades"
	//ONLY use ":=" when Initializing the value (card = "NEW VALUE")

	// card := newCard()

	// fmt.Println(card)

	cards := []string{newCard(), newCard(), "Ace of Diamonds"}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
