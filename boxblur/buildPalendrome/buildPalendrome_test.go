package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := "abcdcba"
	actual := buildPalindrome("abcdc")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := "abababa"
	actual := buildPalindrome("ababab")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := "abba"
	actual := buildPalindrome("abba")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := "abaaba"
	actual := buildPalindrome("abaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
