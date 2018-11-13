package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := isMAC48Address("00-1B-63-84-45-E6")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := false
	actual := isMAC48Address("Z1-1B-63-84-45-E6")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := false
	actual := isMAC48Address("not a MAC-48 address")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := true
	actual := isMAC48Address("FF-FF-FF-FF-FF-FF")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := true
	actual := isMAC48Address("00-00-00-00-00-00")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := false
	actual := isMAC48Address("G0-00-00-00-00-00")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := false
	actual := isMAC48Address("02-03-04-05-06-07-")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := true
	actual := isMAC48Address("12-34-56-78-9A-BC")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	expected := true
	actual := isMAC48Address("FF-FF-AB-CD-EA-BC")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	expected := false
	actual := isMAC48Address("12-34-56-78-9A-BG")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
