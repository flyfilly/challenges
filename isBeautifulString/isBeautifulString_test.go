package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := isBeautifulString("bbbaacdafe")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := false
	actual := isBeautifulString("aabbb")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := false
	actual := isBeautifulString("bbc")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := false
	actual := isBeautifulString("bbbaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := false
	actual := isBeautifulString("abcdefghijklmnopqrstuvwxyzz")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := true
	actual := isBeautifulString("abcdefghijklmnopqrstuvwxyz")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := true
	actual := isBeautifulString("abcdefghijklmnopqrstuvwxyzqwertuiopasdfghjklxcvbnm")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := false
	actual := isBeautifulString("fyudhrygiuhdfeis")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	expected := false
	actual := isBeautifulString("zaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	expected := false
	actual := isBeautifulString("zyy")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
