package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//BOTH of these are the same
	// card := "Ace of Spades"
	//ONLY use ":=" when Initializing the value (card = "NEW VALUE")

	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
