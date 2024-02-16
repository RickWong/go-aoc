package day13

import (
	_ "embed"
	"slices"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type ScanResult struct {
	horizontal bool
	vertical   bool
	offset     int
	errors     int
	score      int
}

// Helper functions.

func TransposeStringSlice(slice []string) []string {
	transpose := make([]string, len(slice[0]))
	buf := make([]byte, len(slice))
	for i := range slice[0] {
		for j := range slice {
			buf[j] = slice[j][i]
		}
		transpose[i] = string(buf)
	}
	return transpose
}

func unsmudgeRows(rows []string) [][]string {
	numVariants := len(rows) * len(rows[0])
	unsmudged := make([][]string, numVariants)
	width := len(rows[0])
	buf := make([]byte, width)

	for i := 0; i < numVariants; i++ {
		unsmudged[i] = slices.Clone(rows)

		y := i / width
		x := i % width
		old := unsmudged[i][y]
		for j := 0; j < len(old); j++ {
			v := old[j]
			if j == x {
				if v == '.' {
					v = '#'
				} else {
					v = '.'
				}
			}
			buf[j] = v
		}
		unsmudged[i][y] = string(buf)

		if old[x] == unsmudged[i][y][x] || len(old) != len(unsmudged[i][y]) {
			panic("unsmudgeRows failed")
		}
	}

	return unsmudged
}

// FindMirror scans the grid to find a scored mirror.
func FindMirror(rows []string) []ScanResult {
	results := make([]ScanResult, 0, len(rows)/2)
	columns := TransposeStringSlice(rows)
	views := [2]*[]string{&rows, &columns}

	for iter := range views {
		view := *views[iter]

		for i := 0; i < len(view)-1; i++ {
			if view[i] == view[i+1] {
				mirrorErrors := 0
				for j := 1; j < min(i+1, len(view)-i-1); j++ {
					for k := 0; k < len(view[i-j]); k++ {
						if view[i-j][k] != view[i+1+j][k] {
							mirrorErrors++
							break
						}
					}
				}

				if mirrorErrors == 0 {
					factor := 1
					if iter == 0 {
						factor = 100
					}

					results = append(results, ScanResult{
						horizontal: factor == 100,
						vertical:   factor == 1,
						offset:     i + 1,
						errors:     mirrorErrors,
						score:      factor * (i + 1),
					})
				}
			}
		}
	}

	return results
}

// Part 1.

func part1() int {
	puzzles := strings.Split(strings.TrimSpace(data), "\n\n")
	sum := 0

	for _, puzzle := range puzzles {
		rows := strings.Split(puzzle, "\n")
		mirror := FindMirror(rows)
		if len(mirror) > 0 {
			sum += mirror[0].score
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 405, result)
	} else {
		assert.Equal(t, 37975, result)
	}
}

// Part 2.

func part2() int {
	puzzles := strings.Split(strings.TrimSpace(data), "\n\n")
	sum := int64(0)
	eg := errgroup.Group{}

	for _, puzzle := range puzzles {
		puzzle := puzzle
		eg.Go(func() error {
			rows := strings.Split(puzzle, "\n")
			originalMirror := FindMirror(rows)
			if len(originalMirror) == 0 {
				panic("no dirty mirror found ")
			}

			puzzleSum := 0
			for _, rows := range unsmudgeRows(rows) {
				smudgelessMirrors := FindMirror(rows)
				if len(smudgelessMirrors) > 0 {
					for k := range smudgelessMirrors {
						if smudgelessMirrors[k] == originalMirror[0] {
							smudgelessMirrors = slices.Delete(smudgelessMirrors, k, k+1)
							break
						}
					}
				}
				if len(smudgelessMirrors) > 0 {
					puzzleSum += smudgelessMirrors[0].score
					break
				}
			}

			atomic.AddInt64(&sum, int64(puzzleSum))

			return nil
		})
	}

	_ = eg.Wait()

	return int(sum)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 400, result)
	} else {
		// 18262 too low
		assert.Equal(t, 32497, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
