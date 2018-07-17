package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{-1, 150, 160, 170, -1, -1, 180, 190}
	actual := sortByHeight([]int{-1, 150, 190, 170, -1, -1, 160, 180})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []int{-1, -1, -1, -1, -1}
	actual := sortByHeight([]int{-1, -1, -1, -1, -1})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := []int{-1}
	actual := sortByHeight([]int{-1})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := []int{2, 2, 4, 9, 11, 16}
	actual := sortByHeight([]int{4, 2, 9, 11, 2, 16})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := []int{1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2}
	actual := sortByHeight([]int{2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := []int{1, 3, -1, 23, 43, -1, -1, 54, -1, -1, -1, 77}
	actual := sortByHeight([]int{23, 54, -1, 43, 1, -1, -1, 77, -1, -1, -1, 3})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
