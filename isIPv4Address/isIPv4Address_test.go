package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := isIPv4Address("172.16.254.1")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := false
	actual := isIPv4Address("172.316.254.1")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := false
	actual := isIPv4Address(".254.255.0")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := false
	actual := isIPv4Address("1.1.1.1a")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := false
	actual := isIPv4Address("1")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := true
	actual := isIPv4Address("0.254.255.0")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := false
	actual := isIPv4Address("1.23.256.255.")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := false
	actual := isIPv4Address("1.23.256..")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
