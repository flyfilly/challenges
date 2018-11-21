package main

import (
	"reflect"
	"testing"
)

func TestCase10(t *testing.T) {
	expected := []string{
		"AXAL",
		"RJ",
		"TALA",
		"TAXA",
		"VITTA",
	}
	actual := wordBoggle([][]string{
		{"A", "X", "V", "W"},
		{"A", "L", "T", "I"},
		{"T", "T", "J", "R"},
	}, []string{
		"AXOLOTL",
		"TAXA",
		"ABA",
		"VITA",
		"VITTA",
		"GO",
		"AXAL",
		"LATTE",
		"TALA",
		"RJ",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase11(t *testing.T) {
	expected := []string{}
	actual := wordBoggle([][]string{
		{"G", "T"},
		{"O", "A"},
	}, []string{})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase12(t *testing.T) {
	expected := []string{
		"GO",
		"GOAT",
	}

	actual := wordBoggle([][]string{
		{"G", "T"},
		{"O", "A"},
	}, []string{
		"TOGGLE",
		"GOAT",
		"TOO",
		"GO",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase13(t *testing.T) {
	expected := []string{
		"CODE",
		"RULES",
	}

	actual := wordBoggle([][]string{
		{"R", "L", "D"},
		{"U", "O", "E"},
		{"C", "S", "O"},
	}, []string{
		"CODE",
		"SOLO",
		"RULES",
		"COOL",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase14(t *testing.T) {
	expected := []string{
		"SOME",
		"WE",
	}

	actual := wordBoggle([][]string{
		{"S", "A"},
		{"M", "O"},
		{"W", "E"},
		{"H", "R"},
	}, []string{
		"SOME",
		"DRONE",
		"WHERE",
		"SOMEWHERE",
		"WORD",
		"WE",
		"MORE",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase15(t *testing.T) {
	expected := []string{
		"DOT", "LOIS", "SIT", "ZOTIS",
	}

	actual := wordBoggle([][]string{
		{"X", "Q", "A", "D", "E"},
		{"Z", "O", "T", "I", "S"},
		{"I", "N", "D", "O", "L"},
		{"Y", "R", "U", "N", "B"},
		{"F", "A", "E", "H", "K"},
		{"X", "Q", "A", "D", "E"},
		{"Z", "O", "T", "I", "S"},
	}, []string{
		"ZOTIS", "SIT", "LOIS",
		"DOT", "SLOB",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
