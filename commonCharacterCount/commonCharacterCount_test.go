package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 3
	actual := commonCharacterCount("aabcc", "adcaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 4
	actual := commonCharacterCount("zzzz", "zzzzzzz")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 3
	actual := commonCharacterCount("abca", "xyzbac")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 0
	actual := commonCharacterCount("a", "b")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase5(t *testing.T) {
	expected := 1
	actual := commonCharacterCount("a", "aaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
