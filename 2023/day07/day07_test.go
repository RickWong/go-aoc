package day07

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Player struct {
	hand  string
	score int
	bid   int
}

func part1() int {
	lines := strings.Split(data, "\n")
	players := make([]Player, 0, len(lines))
	marks := "_23456789TJQKA"

	for _, line := range lines {
		info := strings.Fields(line)
		score, hand, bid := 0, info[0], info[1]
		counts := make(map[string]int, len(hand))

		for i := range hand {
			card := hand[i : i+1]
			counts[card]++
			mark := strings.Index(marks, card)
			score += (mark << (16 - i*4)) & 0x000FFFFF
		}

		for _, count := range counts {
			score += (1 << (20 + count*2)) & 0xFFF00000
		}

		players = append(players, Player{hand, score, common.Atoi(bid)})
	}

	slices.SortFunc(players, func(a, b Player) int { return a.score - b.score })

	sum := 0
	for i, player := range players {
		sum += (i + 1) * player.bid
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 6440
	if data == Input {
		expect = 250453939
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	players := make([]Player, 0, len(lines))
	marks := "_J23456789TQKA"

	for _, line := range lines {
		info := strings.Fields(line)
		score, hand, bid := 0, info[0], info[1]
		counts := make(map[string]int)

		for i := range hand {
			card := hand[i : i+1]
			counts[card]++
			mark := strings.Index(marks, card)
			score += (mark << (16 - i*4)) & 0x000FFFFF
		}

		mostCommonNonJoker := 0
		for card, count := range counts {
			if card != "J" && count > mostCommonNonJoker {
				mostCommonNonJoker = count
			}
		}

		for card, count := range counts {
			if card == "J" {
				if count == 5 { // "JJJJJ" edge case
					score += (1 << (20 + count*2)) & 0xFFF00000
					break
				}

				continue
			}

			if counts["J"] > 0 && count == mostCommonNonJoker {
				count += counts["J"]
				mostCommonNonJoker = -1
			}

			score += (1 << (20 + count*2)) & 0xFFF00000
		}

		players = append(players, Player{hand, score, common.Atoi(bid)})
	}

	slices.SortFunc(players, func(a, b Player) int { return a.score - b.score })

	sum := 0
	for i, player := range players {
		sum += (i + 1) * player.bid
	}

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 5905
	if data == Input {
		expect = 248652697
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
