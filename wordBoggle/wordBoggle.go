package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"
)

func main() {
	fmt.Println(wordBoggle([][]string{
		{"X", "Q", "A", "D", "E"},
		{"Z", "O", "T", "I", "S"},
		{"I", "N", "D", "O", "L"},
		{"Y", "R", "U", "N", "B"},
		{"F", "A", "E", "H", "K"},
	}, []string{
		"ZOTIS", "SIT", "LOIS",
		"DOT", "SLOBS",
	}))
}

type Coord struct {
	X int
	Y int
}

func (coord Coord) isNeighbor(potentialNeighbor Coord) bool {
	if (coord.X == potentialNeighbor.X) && (coord.Y == potentialNeighbor.Y) {
		return false
	}

	diffA := math.Abs(float64(coord.X - potentialNeighbor.X))
	diffB := math.Abs(float64(coord.Y - potentialNeighbor.Y))

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
	start := time.Now()
	found := make([]string, 0)
	wordMap := make(map[string][]Coord, 0)
	alreadyFound := func(str string) bool {
		for _, word := range found {
			if word == str {
				return true
			}
		}

		return false
	}

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
					builtWord = ""
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

		if builtWord == word && !alreadyFound(word) {
			found = append(found, builtWord)
		}
	}

	sort.Strings(found)

	elapsed := time.Since(start)
	log.Printf("WordBoggle took %s", elapsed)
	return found
}
