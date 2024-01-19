package day23

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
	"runtime"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Point struct {
	x, y int
	text string
}

type Trail struct {
	*Point
	steps int
	last  *Trail
}

type Trail2 struct {
	*GraphPoint
	steps   int
	visited bitmap.Bitmap
}

type GraphPoint struct {
	Id    uint32
	X, Y  int
	Text  byte
	Edges map[*GraphPoint]int // V = distance
}

// Helper functions.

func id(y, x int) uint32 {
	return uint32(y<<16 | x)
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := make([][]Point, len(lines))
	for y, line := range lines {
		grid[y] = make([]Point, len(line))
		for x, text := range strings.Split(line, "") {
			grid[y][x] = Point{x, y, text}
		}
	}

	start := &grid[0][strings.Index(lines[0], ".")]
	end := &grid[len(grid)-1][strings.Index(lines[len(lines)-1], ".")]

	result := common.IterativeSearch(
		&Trail{start, 0, nil},
		func(t *Trail) []*Trail {
			branches := make([]*Trail, 0, 3)

			for {
				leftAllowed := (t.text == "<" || t.text == ".") && (t.last == nil || t.last.x != t.x-1)
				upAllowed := (t.text == "^" || t.text == ".") && (t.last == nil || t.last.y != t.y-1)
				downAllowed := (t.text == "v" || t.text == ".") && (t.last == nil || t.last.y != t.y+1)
				rightAllowed := (t.text == ">" || t.text == ".") && (t.last == nil || t.last.x != t.x+1)

				if upAllowed && t.y > 0 && grid[t.y-1][t.x].text != "#" {
					branches = append(branches, &Trail{&grid[t.y-1][t.x], t.steps + 1, t})
				}
				if downAllowed && t.y < len(grid)-1 && grid[t.y+1][t.x].text != "#" {
					branches = append(branches, &Trail{&grid[t.y+1][t.x], t.steps + 1, t})
				}
				if leftAllowed && t.x > 0 && grid[t.y][t.x-1].text != "#" {
					branches = append(branches, &Trail{&grid[t.y][t.x-1], t.steps + 1, t})
				}
				if rightAllowed && t.x < len(grid[0])-1 && grid[t.y][t.x+1].text != "#" {
					branches = append(branches, &Trail{&grid[t.y][t.x+1], t.steps + 1, t})
				}

				if len(branches) == 1 && branches[0].Point != end {
					t = branches[0]
					branches = branches[:0]
					continue
				}

				return branches
			}
		},
		func(t *Trail) bool {
			return t.Point == end
		},
		func(t *Trail) uint32 {
			return uint32(t.y<<16 | t.x)
		},
		func(t *Trail, _ int) int {
			return t.steps
		},
		nil,
		0,
		false,
		true,
	)

	return result.BestWeight
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 94, result)
	} else {
		assert.Equal(t, 2170, result)
	}
}

// Part 2.

func part2() int {
	// Parse graph nodes.
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := make([][]Point, len(lines))
	graph := make([][]*GraphPoint, len(lines))
	for y, line := range lines {
		grid[y] = make([]Point, len(line))
		graph[y] = make([]*GraphPoint, len(line))
		for x, text := range strings.Split(line, "") {
			grid[y][x] = Point{x, y, text}

			if text[0] != '#' {
				graph[y][x] = &GraphPoint{id(y, x), x, y, text[0], make(map[*GraphPoint]int, 4)}
			}
		}
	}

	// Build graph edges.
	for y, row := range graph {
		for x, point := range row {
			if graph[y][x] == nil {
				continue
			}

			if y > 0 && graph[y-1][x] != nil && graph[y-1][x].Text != '#' {
				point.Edges[graph[y-1][x]] = 1
			}
			if y < len(graph)-1 && graph[y+1][x] != nil && graph[y+1][x].Text != '#' {
				point.Edges[graph[y+1][x]] = 1
			}
			if x > 0 && graph[y][x-1] != nil && graph[y][x-1].Text != '#' {
				point.Edges[graph[y][x-1]] = 1
			}
			if x < len(graph[0])-1 && graph[y][x+1] != nil && graph[y][x+1].Text != '#' {
				point.Edges[graph[y][x+1]] = 1
			}
		}
	}

	// Compact graph.
	for y, row := range graph {
		for x, point := range row {
			if point == nil {
				continue
			}

			if len(point.Edges) == 2 {
				// Connect neighbors to each other.
				neighbors := maps.Keys(point.Edges)
				distance := point.Edges[neighbors[0]] + point.Edges[neighbors[1]]
				neighbors[0].Edges[neighbors[1]] = distance
				neighbors[1].Edges[neighbors[0]] = distance

				// Remove point.
				delete(neighbors[0].Edges, point)
				delete(neighbors[1].Edges, point)
				point = nil
				graph[y][x] = nil
			}
		}
	}

	// Prepare DFS.
	start := graph[0][strings.Index(lines[0], ".")]
	end := graph[len(grid)-1][strings.Index(lines[len(lines)-1], ".")]

	// Move starting point to a node with multiple edges, for parallelism.
	offset := 0
	for len(start.Edges) == 1 {
		next := maps.Keys(start.Edges)[0]
		offset += start.Edges[next]
		start = next
	}

	results := make(map[*GraphPoint]int, len(start.Edges))
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())

	for edge, distance := range start.Edges {
		edge := edge
		distance := distance

		eg.Go(func() error {
			println("Starting goroutine")

			visited := make(bitmap.Bitmap, 8)
			visited.Set(start.Id) // Can't go back to start.

			result := common.IterativeSearch[Trail2, int, int](
				&Trail2{edge, distance, visited},
				func(t *Trail2) []*Trail2 {
					branches := make([]*Trail2, 0, 3)
					t.visited = t.visited.Clone(nil)

					for {
						t.visited.Set(t.Id)

						for edge, steps := range t.Edges {
							// Always take the last step, otherwise it cannot be taken later.
							if edge.Id == end.Id {
								return []*Trail2{{edge, t.steps + steps, t.visited}}
							}

							if !t.visited.Contains(edge.Id) {
								branches = append(branches, &Trail2{edge, t.steps + steps, t.visited})
							}
						}

						if len(branches) == 1 && branches[0].Id != end.Id {
							t = branches[0]
							branches = branches[:0]
							continue
						}

						return branches
					}
				},
				func(t *Trail2) bool {
					return t.Id == end.Id
				},
				// Don't use identity func to prune, the longest path could start off with a short path.
				// Only use it if it also includes the edges left?
				nil,
				func(t *Trail2, cw int) int {
					return t.steps
				},
				nil,
				0,
				false,
				true,
			)

			results[edge] = result.BestWeight
			println("result", result.BestWeight)
			return nil
		})
	}

	_ = eg.Wait()
	return slices.Max(maps.Values(results))
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 154, result)
	} else {
		assert.Equal(t, 6502, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
