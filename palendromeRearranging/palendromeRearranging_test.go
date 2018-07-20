package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	actual := palindromeRearranging("aabb")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := false
	actual := palindromeRearranging("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := true
	actual := palindromeRearranging("abbcabb")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := true
	actual := palindromeRearranging("zyyzzzzz")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := true
	actual := palindromeRearranging("z")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCas6(t *testing.T) {
	expected := true
	actual := palindromeRearranging("zaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCas7(t *testing.T) {
	expected := false
	actual := palindromeRearranging("abca")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCas8(t *testing.T) {
	expected := false
	actual := palindromeRearranging("abcad")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCas9(t *testing.T) {
	expected := false
	actual := palindromeRearranging("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbccccaaaaaaaaaaaaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCas10(t *testing.T) {
	expected := false
	actual := palindromeRearranging("abdhuierf")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
