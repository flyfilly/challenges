package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	coord := Coord{0, 0}
	expected := true
	actual := coord.isNeighbor(Coord{0, 1})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	coord := Coord{0, 0}
	expected := true
	actual := coord.isNeighbor(Coord{1, 1})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	coord := Coord{0, 0}
	expected := false
	actual := coord.isNeighbor(Coord{1, 2})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	coord := Coord{1, 0}
	expected := false
	actual := coord.isNeighbor(Coord{1, 2})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	coord := Coord{0, 0}
	expected := false
	actual := coord.isNeighbor(Coord{1, 2})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	coord := Coord{2, 3}
	expected := true
	actual := coord.isNeighbor(Coord{3, 2})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	coord := Coord{2, 3}
	expected := false
	actual := coord.isNeighbor(Coord{2, 3})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	coord := Coord{2, 3}
	expected := true
	actual := coord.isNeighbor(Coord{1, 3})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	coord := Coord{1, 1}
	expected := true
	actual := coord.isNeighbor(Coord{2, 0})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
