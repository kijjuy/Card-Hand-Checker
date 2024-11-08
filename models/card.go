package models

type suit string

type rank string

//const (
//	Hearts suit = iota
//	Clubs
//	Diamonds
//	Spades
//)
//
//const (
//	Ace rank = iota
//	Two
//	Three
//	Four
//	Five
//	Six
//	Seven
//	Eight
//	Nine
//	Ten
//	Jack
//	Queen
//	King
//)

const NumSuits = 4
const NumRanks = 13

var Suits = makeSuits()
var Ranks = makeRanks()

func makeSuits() []suit {
	return []suit{"Hearts", "Clubs", "Diamonds", "Spades"}
}

func makeRanks() []rank {
	return []rank{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
}

type Card struct {
	Suit suit
	Rank rank
}

func MakeDeck() []Card {
	deck := make([]Card, 52)
	count := 0
	for i := 0; i < len(Suits); i++ {
		for j := 0; j < len(Ranks); j++ {
			card := Card{
				Suit: Suits[i],
				Rank: Ranks[j],
			}
			deck[count] = card
			count++
		}
	}
	return deck
}
