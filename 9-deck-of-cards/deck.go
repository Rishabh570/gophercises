//go:generate go run golang.org/x/tools/cmd/stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank int

const (
	_ Rank = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var ranks = [...]Rank{One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Suit Suit
	Rank Rank
}

func New(options ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	// @user: Apply functional options
	// - We usually have an entity that we want to process/mutate
	// - The fn takes a variadic arg which can have n no. of parameters, just like we pass an `options` object as a last arg to a JS fn
	// - The entity is passed to all of those optional args sequentially to process/mutate it
	// - We get back the updated entity value at the end
	// - Pros:
	//   1. Clients don't need to know or pass null/empty values for functional options
	//   2. No changes required in New fn (in this case) to support a new functional option; only requirement is that a functional option need to use the accepted inteface
	for _, option := range options {
		cards = option(cards)
	}

	return cards
}

func (c Card) String() string {
	if c.Suit == Suit(Joker) {
		return c.Suit.String()
	}
	return fmt.Sprintf("Suit: %s, Value: %s", c.Suit, c.Rank)
}

// @user: One of the possible comparators
func DefaultComparator(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank < cards[j].Rank
	}
}

// @dev: The return type of the comparator is enforced by the sort package
// @user: You can pass any comparator to Sort fn below to carry out custom sorting
func Sort(comparator func(cards []Card) func(i, j int) bool, cards []Card) []Card {
	sort.Slice(cards, comparator(cards))
	return cards
}

// @user: Default functional option
func DefaultSort(cards []Card) []Card {
	return Sort(DefaultComparator, cards)
}

// My haphazard implementation
// @user: Shuffle functional option
// func ShuffleCards(cards []Card) []Card {
// 	rand.Shuffle(len(cards), func(i, j int) {
// 		cards[i], cards[j] = cards[j], cards[i]
// 	})
// 	return cards
// }

// Shuffle implementation by Jon
// @user: Shuffle functional option
var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// @user: Jokers functional option
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 1; i <= n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

// @user: Filter functional option
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		results := make([]Card, 0)
		for _, card := range cards {
			if f(card) {
				results = append(results, card)
			}
		}
		return results
	}
}

func Deck(n int) func([]Card) []Card {
	return func([]Card) []Card {
		var results []Card
		for i := 1; i <= n; i++ {
			cards := New()
			results = append(results, cards...)
		}
		return results
	}
}
