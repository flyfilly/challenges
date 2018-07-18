package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{
		"*****",
		"*abc*",
		"*ded*",
		"*****",
	}
	actual := addBorder([]string{
		"abc",
		"ded",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := []string{
		"***",
		"*a*",
		"***",
	}
	actual := addBorder([]string{
		"a",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := []string{
		"****",
		"*aa*",
		"****",
		"*zz*",
		"****",
	}
	actual := addBorder([]string{
		"aa",
		"**",
		"zz",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := []string{
		"*******",
		"*abcde*",
		"*fghij*",
		"*klmno*",
		"*pqrst*",
		"*uvwxy*",
		"*******",
	}
	actual := addBorder([]string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"uvwxy",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := []string{
		"*******",
		"*wzy***",
		"*******",
	}
	actual := addBorder([]string{
		"wzy**",
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
