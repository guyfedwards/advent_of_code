package main

import (
	"testing"
)

type test struct {
	h        Hand
	expected int
}

func TestHandGet(t *testing.T) {
	// h := Hand{
	// }
	ts := []test{
		{
			h: Hand{
				counts: map[string]int{"K": 1, "8": 1, "6": 2, "J": 1},
				cards:  []string{"K", "8", "6", "J", "6"},
			},
			expected: 4,
		},
		{
			h: Hand{
				counts: map[string]int{"A": 3, "K": 2},
				cards:  []string{"A", "K", "K", "A", "A"},
			},
			expected: 5,
		},
		{
			h: Hand{
				counts: map[string]int{"7": 2, "A": 2, "J": 1},
				cards:  []string{"7", "A", "A", "7", "J"},
			},
			expected: 5,
		},
		{
			h: Hand{
				counts: map[string]int{"Q": 2, "K": 3},
				cards:  []string{"K", "K", "Q", "K", "Q"},
			},
			expected: 5,
		},
		{
			h: Hand{
				counts: map[string]int{"A": 1, "4": 2, "3": 1, "J": 1},
				cards:  []string{"A", "4", "3", "4", "J"},
			},
			expected: 4,
		},
	}

	for _, tst := range ts {
		expected := tst.expected
		got := tst.h.getHandWeightJokers()
		if got != expected {
			t.Logf("for %s, expected %d, got %d", tst.h.cards, expected, got)
			t.Fail()
		}
	}
}
