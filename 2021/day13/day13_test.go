package day13

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Dot struct {
	y, x int
}

type Fold struct {
	axis   string
	offset int
}

func part1() int {
	sections := strings.Split(data, "\n\n")
	if sections == nil {
		panic("No data")
	}

	dots, folds := parsePaper(sections)

	foldPoints(dots, folds[:1])

	dots = lo.UniqBy(dots, func(item *Dot) int {
		return item.y*1000000 + item.x
	})

	return len(dots)
}

func parsePaper(sections []string) ([]*Dot, []Fold) {
	dots := lo.Map(
		strings.Split(sections[0], "\n"),
		func(item string, _ int) *Dot {
			s := strings.Split(item, ",")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			return &Dot{y, x}
		},
	)

	folds := lo.Map(
		strings.Split(sections[1], "\n"),
		func(item string, _ int) Fold {
			s := strings.Split(item, "=")
			axis := strings.Replace(s[0], "fold along ", "", 1)
			offset, _ := strconv.Atoi(s[1])
			return Fold{axis, offset}
		},
	)
	return dots, folds
}

func foldPoints(dots []*Dot, folds []Fold) {
	for _, fold := range folds {
		if fold.axis == "y" {
			for idx, dot := range dots[:] {
				mirrorY := dot.y - fold.offset
				if mirrorY > 0 {
					dots[idx] = &Dot{fold.offset - mirrorY, dot.x}
				}
			}
		} else if fold.axis == "x" {
			for idx, dot := range dots[:] {
				mirrorX := dot.x - fold.offset
				if mirrorX > 0 {
					dots[idx] = &Dot{dot.y, fold.offset - mirrorX}
				}
			}
		}
	}
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 17
	if data == Input {
		expect = 751
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	sections := strings.Split(data, "\n\n")
	if sections == nil {
		panic("No data")
	}

	dots, folds := parsePaper(sections)

	foldPoints(dots, folds)

	dots = lo.UniqBy(dots, func(item *Dot) int {
		return item.y*1000000 + item.x
	})

	for y := 0; y < 7; y++ {
		for x := 0; x < 8*5; x++ {
			_, exists := lo.Find(dots, func(item *Dot) bool {
				return item.x == x && item.y == y
			})

			if exists {
				fmt.Print("@")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	return len(dots)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 16
	if data == Input {
		expect = 95
		// Eight capital letters: PGHRKLKL
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
