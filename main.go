package main

func main() {
	/*cards := newDeck()

	hand, remainigDeck := deal(cards, 5)

	hand.print()
	remainigDeck.print()*/

	//cards := newDeck()
	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
