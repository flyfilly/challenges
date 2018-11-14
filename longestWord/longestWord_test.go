package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := "steady"
	actual := longestWord("Ready, steady, go!")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := "steady"
	actual := longestWord("Ready[[[, steady, go!")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := "ABCd"
	actual := longestWord("ABCd")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := "not"
	actual := longestWord("To be or not to be")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := "CodeFighter"
	actual := longestWord("You are the best!!!!!!!!!!!! CodeFighter ever!")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
