package main

func main() {
	// var card string = "Ace of Spades"
	//BOTH of these are the same
	// card := "Ace of Spades"
	//ONLY use ":=" when Initializing the value (card = "NEW VALUE")

	// card := newCard()

	// fmt.Println(card)

	cards := newDeck()

	// cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}

//To run MULTPIPLE Files you call them in the same GO RUN (go run main.go deck.go)

// func newCard() string {
// 	return "Five of Diamonds"
// }
