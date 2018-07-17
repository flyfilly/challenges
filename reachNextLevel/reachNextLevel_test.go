package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := reachNextLevel(10, 15, 5)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	actual := reachNextLevel(10, 15, 4)
	expected := false

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	actual := reachNextLevel(3, 6, 4)
	expected := true

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	actual := reachNextLevel(83, 135, 36)
	expected := false

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	actual := reachNextLevel(74, 170, 58)
	expected := false

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	actual := reachNextLevel(80, 199, 15)
	expected := false

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	actual := reachNextLevel(97, 57, 7)
	expected := true

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	actual := reachNextLevel(235, 293, 4)
	expected := false

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	actual := reachNextLevel(222, 137, 58)
	expected := true

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	actual := reachNextLevel(16, 23, 16)
	expected := true

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}
