package day02

import (
	_ "embed"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Game struct {
	id   int
	sets []GameSet
}

type GameSet map[string]int

func part1() int {
	lines := strings.Split(data, "\n")
	games := parseGames(lines)
	sum := 0
	maxPossible := GameSet{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range *games {
		isPossible := lo.EveryBy(game.sets, func(set GameSet) bool {
			for color, num := range set {
				if num > maxPossible[color] {
					return false
				}
			}

			return true
		})

		if isPossible {
			sum += game.id
		}
	}

	return sum
}

func parseGames(lines []string) *[]Game {
	games := make([]Game, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ": ")
		games[i].id, _ = strconv.Atoi(parts[0][5:len(parts[0])])

		sets := strings.Split(parts[1], ";")
		games[i].sets = make([]GameSet, len(sets))

		for j, set := range sets {
			colorCubes := strings.Split(set, ", ")
			games[i].sets[j] = make(GameSet, len(colorCubes))

			for _, colorCube := range colorCubes {
				info := strings.Fields(colorCube)
				num, _ := strconv.Atoi(info[0])
				color := info[1]
				games[i].sets[j][color] = num
			}
		}
	}

	return &games
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 8
	if data == Input {
		expect = 2541
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	games := parseGames(lines)
	sum := 0

	for _, game := range *games {
		maxima := make(GameSet)

		for _, set := range game.sets {
			for color, num := range set {
				currentMaximum, ok := maxima[color]
				if !ok {
					maxima[color] = num
				} else {
					maxima[color] = max(currentMaximum, num)
				}
			}
		}

		power := multiply(lo.Values(maxima))
		sum += power
	}

	return sum
}

func multiply(v []int) int {
	result := v[0]
	for i := 1; i < len(v); i++ {
		result *= v[i]
	}
	return result
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 2286
	if data == Input {
		expect = 66016
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
