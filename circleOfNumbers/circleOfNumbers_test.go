package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 7
	actual := circleOfNumbers(10, 2)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 2
	actual := circleOfNumbers(10, 7)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 3
	actual := circleOfNumbers(4, 1)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 0
	actual := circleOfNumbers(6, 3)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
