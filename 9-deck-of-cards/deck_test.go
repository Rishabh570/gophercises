package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	deckLength = 52
)

func TestNew(t *testing.T) {
	deck := New()

	expectedDeckLen := 13 * 4

	if len(deck) != expectedDeckLen {
		t.Error("Deck length is not as expected")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	fmt.Println("cards:", cards)
	expectedFirstCard := Card{Suit: Spade, Rank: One}
	if cards[0] != expectedFirstCard {
		t.Error("Cards are not sorted as expected")
	}
}

func TestJokers(t *testing.T) {
	n := 2
	cards := New(Jokers(n))
	jokersCount := 0

	for _, card := range cards {
		if card.Suit == Joker {
			jokersCount++
		}
	}

	if jokersCount != n {
		t.Error("Wrong number of jokers in the deck")
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(sampleFilter))
	fmt.Println("cards:", len(cards))

	for _, card := range cards {
		fmt.Println("card:", card.Suit, "rank:", card.Rank)
		if card.Suit != Spade {
			t.Error("Non spade card is present in the resulting deck")
		}
	}

	if len(cards) != 13 {
		t.Error("Wrong number of cards in the resulting set")
	}
}

func TestDeck(t *testing.T) {
	n := 3
	cards := New(Deck(n))

	if len(cards) != deckLength*n {
		t.Error("Invalid deck generated, length does not match")
	}
}

func sampleFilter(card Card) bool {
	return card.Suit == Spade
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 ... ]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s.", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s.", second, cards[1])
	}
}
