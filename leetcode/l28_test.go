package l28

import "testing"

func TestCase1(t *testing.T) {
	sum := strStr("sadbutsad", "sad")
	expected := 0

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCase2(t *testing.T) {
	sum := strStr("leetcode", "leeto")
	expected := -1

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCase3(t *testing.T) {
	sum := strStr("asadbutsad", "sad")
	expected := 1

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
