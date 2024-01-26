package day23

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/edwingeng/deque/v2"
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

type GraphTrail struct {
	*GraphPoint
	steps   int
	visited uint64
}

type GraphPoint struct {
	Id    int
	X, Y  int
	Edges map[*GraphPoint]int // V = distance
}

// Helper functions.

func parseGraph(lines []string) [][]*GraphPoint {
	graph := make([][]*GraphPoint, len(lines))
	for y, line := range lines {
		graph[y] = make([]*GraphPoint, len(line))
		for x, text := range strings.Split(line, "") {
			if text[0] != '#' {
				graph[y][x] = &GraphPoint{0, x, y, make(map[*GraphPoint]int, 4)}
			}
		}
	}
	// Build graph edges.
	for y, row := range graph {
		for x, point := range row {
			if graph[y][x] == nil {
				continue
			}

			if y > 0 && graph[y-1][x] != nil {
				point.Edges[graph[y-1][x]] = 1
			}
			if y < len(graph)-1 && graph[y+1][x] != nil {
				point.Edges[graph[y+1][x]] = 1
			}
			if x > 0 && graph[y][x-1] != nil {
				point.Edges[graph[y][x-1]] = 1
			}
			if x < len(graph[0])-1 && graph[y][x+1] != nil {
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
	return graph
}

func assignGraphIDs(start *GraphPoint) {
	nextId := 1
	queue := deque.NewDeque[*GraphPoint](deque.WithChunkSize(16))
	queue.PushBack(start)
	for !queue.IsEmpty() {
		current := queue.PopFront()
		if current.Id != 0 {
			continue
		}

		current.Id = nextId
		nextId++

		for edge := range current.Edges {
			if edge.Id == 0 {
				queue.PushBack(edge)
			}
		}
	}
}

func findEdgesLeft(point *GraphPoint, visited uint64) uint64 {
	left := uint64(0)
	for edge := range point.Edges {
		if visited&(1<<edge.Id) == 0 {
			left |= 1 << edge.Id
			left |= findEdgesLeft(edge, left|visited)
		}
	}
	return left
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

	result := common.BucketSearch(
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
					branches[0] = nil
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
		-1000,
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
	graph := parseGraph(lines)

	// Prepare for DFS.
	start := graph[0][strings.Index(lines[0], ".")]
	end := graph[len(lines)-1][strings.Index(lines[len(lines)-1], ".")]

	// Re-assign IDs with BFS.
	assignGraphIDs(start)

	// Move starting point to a node with multiple edges, for parallelism.
	startOffset := 0
	visited := uint64(1 << start.Id)
	for len(start.Edges) == 1 {
		next := maps.Keys(start.Edges)[0]
		startOffset += start.Edges[next]
		start = next
		visited |= 1 << start.Id
	}
	// Move end point to a node with multiple edges, to finish early.
	endOffset := 0
	for len(end.Edges) == 1 {
		prev := maps.Keys(end.Edges)[0]
		endOffset += end.Edges[prev]
		end = prev
	}

	starts := maps.Keys(start.Edges)
	results := make([]int, len(starts))
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())

	for i := 0; i < len(starts); i++ {
		i := i
		edge := starts[i]
		distance := start.Edges[edge]

		eg.Go(func() error {
			result := common.BucketSearch[GraphTrail, uint64, int](
				&GraphTrail{edge, distance, visited},
				func(t *GraphTrail) []*GraphTrail {
					branches := make([]*GraphTrail, 0, 3)
					nextVisited := t.visited + 0 // Copy.

					for {
						nextVisited |= 1 << t.Id

						for edge, steps := range t.Edges {
							if t.visited&(1<<edge.Id) == 0 {
								// Source: https://www.reddit.com/r/adventofcode/comments/18oy4pc/2023_day_23_solutions/kfyvp2g/
								// Skip perimiter edge paths that "turn backwards" as they inevitably lead to a snake-ish dead end.
								perimeter := len(t.Edges) <= 3 && len(edge.Edges) <= 3
								behind := edge.Id+1 < t.Id
								if perimeter && behind {
									continue
								}

								branches = append(branches, &GraphTrail{edge, t.steps + steps, nextVisited})
							}
						}

						if len(branches) == 1 && branches[0].Id != end.Id {
							t = branches[0]
							branches[0] = nil
							branches = branches[:0]
							continue
						}

						return branches
					}
				},
				func(t *GraphTrail) bool {
					return t.Id == end.Id
				},
				// Normally can't prune based on position, since the longest path could start off with a short path.
				// But here we account for edges left, so only branches on a position with shared remaining paths are pruned.
				func(t *GraphTrail) uint64 {
					return uint64(t.Id<<54) | findEdgesLeft(t.GraphPoint, t.visited)
				},
				func(t *GraphTrail, cw int) int {
					return t.steps
				},
				nil,
				-1000,
				false,
				true,
			)

			results[i] = result.BestWeight
			println("result", result.BestWeight)
			return nil
		})
	}

	_ = eg.Wait()
	return startOffset + slices.Max(results) + endOffset
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
