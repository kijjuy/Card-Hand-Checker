package main

import (
	"cards/models"
	"fmt"
	"math/rand/v2"
)

var (
	trialsCount  = 0
	trialsPassed = 0
)

const handSize = 4

func drawCard(deck *[]models.Card) models.Card {
	i := rand.IntN(len(*deck))
	drawnCard := (*deck)[i]
	*deck = append((*deck)[:i], (*deck)[i+1:]...)
	return drawnCard
}

func drawHand(deck *[]models.Card) []models.Card {
	hand := make([]models.Card, handSize)
	for i := 0; i < handSize; i++ {
		drawnCard := drawCard(deck)
		hand[i] = drawnCard
	}
	return hand
}

func checkHand(hand *[]models.Card) {
	if (*hand)[0].Suit != "Spades" && (*hand)[1].Suit != "Spades" && (*hand)[2].Suit == "Spades" && (*hand)[3].Suit != "Spades" {
		trialsPassed++
	}
	trialsCount++
}

func drawManyCards() {
	cap := 10000000
	for i := 0; i < cap; i++ {
		deck := models.MakeDeck()
		hand := drawHand(&deck)

		checkHand(&hand)

		if i%1000 == 0 {
			numRemaining := cap - i
			fmt.Printf("Done %d. %d more to go.\n", i, numRemaining)
		}
	}
}

func main() {
	drawManyCards()
	fmt.Printf("Total trials: %d\n", trialsCount)
	fmt.Printf("Total valid hands: %d\n", trialsPassed)

}
