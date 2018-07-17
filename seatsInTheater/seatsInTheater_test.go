package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 96
	actual := seatsInTheater(16, 11, 5, 3)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 0
	actual := seatsInTheater(1, 1, 1, 1)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 18
	actual := seatsInTheater(13, 6, 8, 3)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase6(t *testing.T) {
	expected := 99
	actual := seatsInTheater(60, 100, 60, 1)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
