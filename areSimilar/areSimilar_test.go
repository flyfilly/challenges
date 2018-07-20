package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCase1(t *testing.T) {
	start := time.Now()
	expected := true
	actual := areSimilar([]int{1, 2, 3}, []int{1, 2, 3})
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	if expected != actual && elapsed < 4 {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	expected := true
	actual := areSimilar([]int{1, 2, 3}, []int{2, 1, 3})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	expected := false
	actual := areSimilar([]int{1, 2, 2}, []int{2, 1, 1})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	expected := false
	actual := areSimilar([]int{1, 1, 4}, []int{1, 2, 3})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	expected := false
	actual := areSimilar([]int{1, 2, 3}, []int{1, 10, 2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase6(t *testing.T) {
	expected := true
	actual := areSimilar([]int{2, 3, 1}, []int{1, 3, 2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase7(t *testing.T) {
	expected := false
	actual := areSimilar([]int{2, 3, 9}, []int{10, 3, 2})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase8(t *testing.T) {
	expected := false
	actual := areSimilar([]int{4, 6, 3}, []int{3, 4, 6})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase9(t *testing.T) {
	expected := true
	actual := areSimilar([]int{
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
	}, []int{
		832, 998, 148, 570, 533, 561, 455, 147, 894, 279,
	})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase10(t *testing.T) {
	expected := false
	actual := areSimilar([]int{
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
	}, []int{
		832, 570, 148, 998, 533, 561, 455, 147, 894, 279,
	})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}

func TestCase11(t *testing.T) {
	expected := true
	actual := areSimilar([]int{
		832, 570, 570, 998, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 148, 533, 561, 894, 147, 455, 279,
	}, []int{
		832, 570, 148, 998, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 455, 147, 894, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
		832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
	})

	if expected != actual {
		t.Errorf("\nTests failed!\nExpected:\t%t\nActual:\t%t", expected, actual)
	}
}
