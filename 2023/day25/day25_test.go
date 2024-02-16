package day25

import (
	_ "embed"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"github.com/twmb/algoimpl/go/graph"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

// Helper functions.

// Part 1.

func part1() int {
	components := make(map[string]graph.Node, 2000)
	group := graph.New(graph.Undirected)

	re := regexp.MustCompile(`(?m)(\w+): (.+)$`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		if _, ok := components[match[1]]; !ok {
			components[match[1]] = group.MakeNode()
		}

		for _, connection := range strings.Fields(match[2]) {
			if _, ok := components[connection]; !ok {
				components[connection] = group.MakeNode()
			}

			_ = group.MakeEdgeWeight(components[match[1]], components[connection], 1)
		}
	}

	cuts := group.RandMinimumCut(len(components)/4, runtime.NumCPU())
	for _, cut := range cuts {
		group.RemoveEdge(cut.Start, cut.End)
	}

	return common.Product(common.Map(
		group.StronglyConnectedComponents(),
		func(components []graph.Node) int { return len(components) },
	))
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 54, result)
	} else {
		assert.Equal(t, 596376, result)
	}
}

// Part 2.

func part2() int {
	return -1
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, -1, result)
	} else {
		assert.Equal(t, -1, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
