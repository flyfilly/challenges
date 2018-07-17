package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{
		"aba",
		"vcd",
		"aba",
	}
	actual := allLongestStrings([]string{
		"aba",
		"aa",
		"ad",
		"vcd",
		"aba",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []string{
		"aa",
	}
	actual := allLongestStrings([]string{
		"aa",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := []string{
		"eeee",
		"abcd",
	}
	actual := allLongestStrings([]string{
		"abc",
		"eeee",
		"abcd",
		"dcd",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := []string{
		"zzzzzz",
		"abcdef",
		"aaaaaa",
	}
	actual := allLongestStrings([]string{
		"a",
		"abc",
		"cbd",
		"zzzzzz",
		"a",
		"abcdef",
		"asasa",
		"aaaaaa",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}
