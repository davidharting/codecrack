package strings

import "testing"

func TestHasAllUniqueCharacters(t *testing.T) {
	type pair struct {
		input    string
		expected bool
	}

	pairs := [2]pair{
		pair{"abc", true},
		pair{"aabc", false},
	}

	for _, p := range pairs {
		got := HasAllUniqueCharacters(p.input)
		if got != p.expected {
			t.Errorf("For input \"%v\" expected %v, got %v", p.input, p.expected, got)
		}
	}
}

func TestReverse(t *testing.T) {
	type pair struct {
		input    string
		expected string
	}

	pairs := [2]pair{
		pair{"david harting", "gnitrah divad"},
		pair{"Today is a great day!", "!yad taerg a si yadoT"},
	}

	for _, p := range pairs {
		got := Reverse(p.input)
		if got != p.expected {
			t.Errorf("For input \"%v\" expected \"%v\", got \"%v\"", p.input, p.expected, got)
		}
	}
}
