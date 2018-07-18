package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{180, 105}
	actual := alternatingSums([]int{
		50, 60, 60, 45, 70,
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []int{100, 50}
	actual := alternatingSums([]int{
		100, 50,
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := []int{80, 0}
	actual := alternatingSums([]int{
		80,
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := []int{150, 150}
	actual := alternatingSums([]int{
		100, 50, 50, 100,
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
