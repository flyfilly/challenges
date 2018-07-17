package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := "acbde"
	actual := reverseParentheses("a(bc)de")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := "apmnolkjihgfedcbq"
	actual := reverseParentheses("a(bcdefghijkl(mno)p)q")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := "cosfighted"
	actual := reverseParentheses("co(de(fight)s)")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := "CodeegnlleahC"
	actual := reverseParentheses("Code(Cha(lle)nge)")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := "Where are the parentheses?"
	actual := reverseParentheses("Where are the parentheses?")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := "abcabcabcabc"
	actual := reverseParentheses("abc(cba)ab(bac)c")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := "The god quick nworb xof jumps over the lazy"
	actual := reverseParentheses("The ((quick (brown) (fox) jumps over the lazy) dog)")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%s\nActual:%s", expected, actual)
	}
}
