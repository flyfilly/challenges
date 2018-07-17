package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := isLucky(1230)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := false
	actual := isLucky(239017)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := true
	actual := isLucky(134008)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := false
	actual := isLucky(10)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := true
	actual := isLucky(11)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := true
	actual := isLucky(1010)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := false
	actual := isLucky(261534)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := false
	actual := isLucky(100000)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	expected := true
	actual := isLucky(999999)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	expected := true
	actual := isLucky(123321)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%t\nActual:%t", expected, actual)
	}
}
