package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	expected := "2a3bc"
	actual := lineEncoding("aabbbc")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := "a2bca2b"
	actual := lineEncoding("abbcabb")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := "abcd"
	actual := lineEncoding("abcd")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := "4z"
	actual := lineEncoding("zzzz")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := "7wa7w"
	actual := lineEncoding("wwwwwwwawwwwwww")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := "15c"
	actual := lineEncoding("ccccccccccccccc")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := "qwertyuioplkjhg"
	actual := lineEncoding("qwertyuioplkjhg")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := "2s2i2gk3o"
	actual := lineEncoding("ssiiggkooo")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	expected := "adf3a"
	actual := lineEncoding("adfaaa")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	expected := "2bj2adlkjdl"
	actual := lineEncoding("bbjaadlkjdl")

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:%v\nActual:%v", expected, actual)
	}
}
