package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 3
	actual := arrayMaximalAdjacentDifference([]int{2, 4, 1, 0})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 0
	actual := arrayMaximalAdjacentDifference([]int{1, 1, 1, 1})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 7
	actual := arrayMaximalAdjacentDifference([]int{-1, 4, 10, 3, -2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 2
	actual := arrayMaximalAdjacentDifference([]int{10, 11, 13})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase5(t *testing.T) {
	expected := 0
	actual := arrayMaximalAdjacentDifference([]int{-2, -2, -2, -2, -2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase6(t *testing.T) {
	expected := 4
	actual := arrayMaximalAdjacentDifference([]int{-1, 1, -3, -4})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
