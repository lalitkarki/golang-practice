package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	cards.print()
}
