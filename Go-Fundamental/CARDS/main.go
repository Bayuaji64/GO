package main

var deckSize int

func main() {
	// var card string = "Ace of Spades"

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {

	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// fmt.Println(cards)

	// deckSize = 20
	// fmt.Println(deckSize)

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_c")

	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
