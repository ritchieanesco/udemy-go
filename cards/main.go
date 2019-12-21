package main

func main() {

	// splitting slices
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	// save to file example
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	// read file from hard drive
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// shuffling string slice
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
