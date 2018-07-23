package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := areEquallyStrong(10, 15, 15, 10)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := true
	actual := areEquallyStrong(15, 10, 15, 10)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := false
	actual := areEquallyStrong(15, 10, 15, 9)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := true
	actual := areEquallyStrong(10, 5, 5, 10)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := false
	actual := areEquallyStrong(10, 15, 5, 20)

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
