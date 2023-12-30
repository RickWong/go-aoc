package day25

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	graph2 "github.com/twmb/algoimpl/go/graph"
	"regexp"
	"strings"
	"testing"
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
	nodes := make(map[string]graph2.Node, 1000)
	graph := graph2.New(graph2.Undirected)

	re := regexp.MustCompile(`(?m)(\w+): (.+)$`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		if _, ok := nodes[match[1]]; !ok {
			nodes[match[1]] = graph.MakeNode()
		}

		for _, connection := range strings.Fields(match[2]) {
			if _, ok := nodes[connection]; !ok {
				nodes[connection] = graph.MakeNode()
			}

			_ = graph.MakeEdgeWeight(nodes[match[1]], nodes[connection], 1)
			//_ = graph.MakeEdge(nodes[connection], nodes[match[1]])
		}
	}

	for key, node := range nodes {
		*node.Value = key
	}

	cuts := graph.RandMinimumCut(1000, 10)
	for _, cut := range cuts {
		graph.RemoveEdge(cut.Start, cut.End)
	}

	cycles := graph.StronglyConnectedComponents()
	ret := 1
	for _, cycle := range cycles {
		ret *= len(cycle)
	}

	return ret
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
	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 357134560737, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
