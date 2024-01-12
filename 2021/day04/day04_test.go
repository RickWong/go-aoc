package day04

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var input string

var data = input

type MarkableNumber struct {
	string
	marked bool
}
type Board struct {
	id   int
	grid [][]MarkableNumber
}

func sumUnmarkedNumbers(board *Board, drawnNumber string) int {
	sumUnmarked := 0
	for _, rows := range board.grid {
		for _, item := range rows {
			if !item.marked {
				unmarkedNumber, _ := strconv.ParseInt(item.string, 10, 0)
				sumUnmarked += int(unmarkedNumber)
			}
		}
	}
	lastNumber, _ := strconv.ParseInt(drawnNumber, 10, 0)
	return sumUnmarked * int(lastNumber)
}

func part1() int {
	lines := strings.Split(data, "\n")
	drawnNumbers, boards := parseData(lines)
	score := 0

out:
	for _, drawnNumber := range drawnNumbers {
		for _, board := range boards {
			score = checkBoard(board, drawnNumber)
			if score > 0 {
				break out
			}
		}
	}

	return score
}

func parseData(lines []string) ([]string, []Board) {
	if lines == nil {
		panic("No data")
	}

	drawnNumbers := strings.Split(lines[0], ",")
	boards := make([]Board, 0)
	board := Board{}

	for i := 2; i < len(lines); i++ {
		numbers := strings.Fields(lines[i])

		if len(numbers) == 0 {
			continue
		}

		board.grid = append(board.grid,
			common.Map(numbers,
				func(s string) MarkableNumber {
					return MarkableNumber{s, false}
				},
			),
		)

		if len(board.grid) == 5 {
			boards = append(boards, board)
			board = Board{id: len(boards)}
		}
	}

	return drawnNumbers, boards
}

func checkBoard(board Board, drawnNumber string) int {
	for y := 0; y < len(board.grid); y++ {
		for x := 0; x < len(board.grid[y]); x++ {
			if board.grid[y][x].string == drawnNumber {
				board.grid[y][x].marked = true

				win := true
				for cx := 0; cx < len(board.grid[y]); cx++ {
					if !board.grid[y][cx].marked {
						win = false
						break
					}
				}

				if !win {
					win = true
					for cy := 0; cy < len(board.grid); cy++ {
						if !board.grid[cy][x].marked {
							win = false
							break
						}
					}
				}

				if win {
					score := sumUnmarkedNumbers(&board, drawnNumber)
					return score
				}
			}
		}
	}

	return 0
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 4512
	if data == input {
		expect = 21607
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	drawnNumbers, boards := parseData(lines)
	score := 0

	for _, drawnNumber := range drawnNumbers {
		for idx, board := range boards[:] {
			score := checkBoard(board, drawnNumber)
			if score > 0 {
				if len(boards) == 1 {
					return score
				}

				if idx < len(boards)-1 {
					boards[idx] = boards[len(boards)-1]
				}
				boards = boards[:len(boards)-1]
			}
		}
	}

	return score
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 1924
	if data == input {
		expect = 19012
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}

}
