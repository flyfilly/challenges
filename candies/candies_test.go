package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 9
	actual := candies(3, 10)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 2
	actual := candies(1, 2)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 0
	actual := candies(10, 5)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
