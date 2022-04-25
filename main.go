package main

func main() {
	cards := newDeck()
	cards.shufleDeck()

	cards.print()
	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	// fmt.Println(hand.toString())
	// hand.saveToFile("myfile.txt")

	// newCarsd := newDeckFromFile("myfile1.txt")
	// fmt.Println("Cards from file: ", newCarsd)
}
