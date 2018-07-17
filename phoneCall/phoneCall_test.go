package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 14
	actual := phoneCall(3, 1, 2, 20)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 1
	actual := phoneCall(2, 2, 1, 2)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 11
	actual := phoneCall(10, 1, 2, 22)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 14
	actual := phoneCall(2, 2, 1, 24)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase5(t *testing.T) {
	expected := 3
	actual := phoneCall(1, 2, 1, 6)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase6(t *testing.T) {
	expected := 0
	actual := phoneCall(10, 10, 10, 8)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
