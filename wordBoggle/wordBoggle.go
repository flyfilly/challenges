package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(wordBoggle([][]string{
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
	}))
}

type Coord struct {
	X int
	Y int
}

func (coord Coord) isNeighbor(check Coord) bool {
	if (coord.X == check.X) && (coord.Y == check.Y) {
		return false
	}

	diffA := math.Abs(float64(coord.X - check.X))
	diffB := math.Abs(float64(coord.Y - check.Y))

	return diffA <= 1 && diffB <= 1
}

func (coord Coord) inExempt(exempt []Coord) bool {
	for _, check := range exempt {
		if (coord.X == check.X) && (coord.Y == check.Y) {
			return true
		}
	}
	return false
}

func wordBoggle(board [][]string, words []string) []string {
	found := make([]string, 0)
	wordMap := make(map[string][]Coord, 0)

	for y, row := range board {
		for x, cell := range row {
			wordMap[cell] = append(wordMap[cell], Coord{x, y})
		}
	}

	for _, word := range words {
		builtWord := ""
		var prevCoord Coord
		var exempts []Coord

		for idx, letter := range word {
			char := string(letter)
			coords := wordMap[char]

			for _, coord := range coords {
				if idx == 0 {
					prevCoord = coord
					builtWord += char
				} else {
					if good := prevCoord.isNeighbor(coord); good && !coord.inExempt(exempts) {
						prevCoord = coord
						builtWord += char
						exempts = append(exempts, prevCoord)
						break
					}
				}
			}
		}
		fmt.Println(builtWord, word)
		if builtWord == word {
			found = append(found, builtWord)
		}
	}

	sort.Strings(found)
	return found
}
