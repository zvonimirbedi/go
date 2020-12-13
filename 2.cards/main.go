package main

import "fmt"

func main() {
	// string variable declaration
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println(card)

	// get data from another file of type deck
	cards := newDeck()
	cards.print()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println(hand)
	fmt.Println(remainingDeck)

	// slice to string
	fmt.Println("Slice converted to string")
	fmt.Println(cards.deckToString())

	// save to file
	// cards.saveToFile("my_cards.txt")

	cards = newDeckFromFile("my_cards.txt")
	cards.print()
	cards.shuffleDeck()
	cards.print()

}
