package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	actual := knapsackLight(10, 5, 6, 4, 8)
	expected := 10

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	actual := knapsackLight(10, 5, 6, 4, 9)
	expected := 16

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	actual := knapsackLight(5, 3, 7, 4, 6)
	expected := 7

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	actual := knapsackLight(10, 2, 11, 3, 1)
	expected := 0

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase5(t *testing.T) {
	actual := knapsackLight(15, 2, 20, 3, 2)
	expected := 15

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase6(t *testing.T) {
	actual := knapsackLight(2, 5, 3, 4, 5)
	expected := 3

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase7(t *testing.T) {
	actual := knapsackLight(4, 3, 3, 4, 4)
	expected := 4

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase8(t *testing.T) {
	actual := knapsackLight(3, 5, 3, 8, 10)
	expected := 3

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
