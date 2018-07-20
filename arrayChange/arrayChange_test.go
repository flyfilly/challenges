package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 3
	actual := arrayChange([]int{1, 1, 1})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 5
	actual := arrayChange([]int{-1000, 0, -2, 0})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 12
	actual := arrayChange([]int{2, 1, 10, 1})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 13
	actual := arrayChange([]int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
