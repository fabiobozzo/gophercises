package deck

import (
	"fmt"
	"math/rand"
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
		t.Errorf("expected the first card to be %s, got %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("expected the second card to be %s, got %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 4 {
		t.Error("expected 4 jokers. got: ", count)
	}
}

func TestFilterOut(t *testing.T) {
	twoAndThrees := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(FilterOut(twoAndThrees))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("expected all of two and threes are filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards, got: %d", 13*4*3, len(cards))
	}
}
