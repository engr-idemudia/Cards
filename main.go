package main

func main() {
	// listOfCards := newDeck()
	// //listOfCards.print()
	// //newDeck().print()   //this is easier

	// hand, remainingCards := deal(listOfCards, 5)

	// hand.print()
	// remainingCards.print()

	//greetings := "saga=1"

	// listOfCards := newDeck()
	// //fmt.Println(listOfCards.toString())
	// listOfCards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
