package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Error("expected Ace of Spades as first card. got: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Error("expected Ace of Spades as first card. got: ", cards[0])
	}
}