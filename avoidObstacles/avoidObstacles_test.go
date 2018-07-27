package main

import (
	"strconv"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 4
	actual := avoidObstacles([]int{5, 3, 6, 7, 9})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase2(t *testing.T) {
	expected := 4
	actual := avoidObstacles([]int{2, 3})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase3(t *testing.T) {
	expected := 7
	actual := avoidObstacles([]int{1, 4, 10, 6, 2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase4(t *testing.T) {
	expected := 6
	actual := avoidObstacles([]int{1000, 999})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase5(t *testing.T) {
	expected := 6
	actual := avoidObstacles([]int{1000, 999})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}

func TestCase6(t *testing.T) {
	expected := 6
	actual := avoidObstacles([]int{1000, 999})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", strconv.Itoa(expected), strconv.Itoa(actual))
	}
}
