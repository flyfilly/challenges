package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 2
	actual := electionsWinners([]int{2, 3, 5, 2}, 3)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := 0
	actual := electionsWinners([]int{1, 3, 3, 1, 1}, 0)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := 1
	actual := electionsWinners([]int{5, 1, 3, 4, 1}, 0)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := 4
	actual := electionsWinners([]int{1, 1, 1, 1}, 1)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := 0
	actual := electionsWinners([]int{1, 1, 1, 1}, 0)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := 2
	actual := electionsWinners([]int{3, 1, 1, 3, 1}, 2)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
