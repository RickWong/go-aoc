package day22

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"maps"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Voxel struct {
	name    string
	x, y, z int
}

type Brick struct {
	name  string
	start Voxel
	end   Voxel
	cubes []Voxel
}

// Helper functions.

func Cubes(b *Brick) []Voxel {
	res := make([]Voxel, 0)

	if b.start.x > b.end.x ||
		b.start.y > b.end.y ||
		b.start.z > b.end.z {
		panic("inverse brick")
	}

	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			for z := b.start.z; z <= b.end.z; z++ {
				res = append(res, Voxel{b.name, x, y, z})
			}
		}
	}

	return res
}

// calculateDroppedHeights returns the heights of the bricks after dropping.
func calculateDroppedHeights(bricks []Brick, skip int) map[string]int {
	heights := make(map[string]int, len(bricks))
	collisionMap := makeHeightmap(bricks)

	// Drop the bricks one by one.
	for i := 0; i < len(bricks); i++ {
		if skip == i {
			continue
		}

		b1 := &bricks[i]
		distance := 100000000

		for _, cube := range b1.cubes {
			b2 := collisionMap[cube.x][cube.y]
			h := b2.z + 1
			distance = min(distance, cube.z-h)
		}

		// Stored dropped height of b1.
		heights[b1.name] = b1.start.z - distance

		for _, cube := range b1.cubes {
			b2 := collisionMap[cube.x][cube.y]
			droppedZ := cube.z - distance

			if b2.name == "" || droppedZ > b2.z {
				collisionMap[cube.x][cube.y] = Voxel{b1.name, cube.x, cube.y, droppedZ}
			}
		}
	}

	return heights
}

// makeHeightmap returns a 2D grid of voxels.
func makeHeightmap(bricks []Brick) map[int]map[int]Voxel {
	// Find max of X.
	maxX := 0
	for _, brick := range bricks {
		maxX = max(maxX, brick.end.x)
	}

	// Create the 2D grid.
	heightmap := make(map[int]map[int]Voxel)
	for x := 0; x <= maxX; x++ {
		heightmap[x] = make(map[int]Voxel)
	}

	return heightmap
}

// Part 1.

func part1() int {
	// Parse the bricks.
	bricksRe := regexp.MustCompile(`(?m)^(\d+).*(\d+),(\d+)~(\d+),(\d+),(\d+)`)
	bricks := make([]Brick, 0)
	for _, match := range bricksRe.FindAllStringSubmatch(strings.TrimSpace(data), -1) {
		id := strconv.Itoa(len(bricks) + 1)
		brick := Brick{
			id,
			Voxel{id, Atoi(match[1]), Atoi(match[2]), Atoi(match[3])},
			Voxel{id, Atoi(match[4]), Atoi(match[5]), Atoi(match[6])},
			nil,
		}
		brick.cubes = Cubes(&brick)
		bricks = append(bricks, brick)
	}

	// Sort from bottom to top.
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})

	// Calculate the heights of the bricks after dropping.
	brickHeights := calculateDroppedHeights(bricks, -1)
	sum := 0

	// Count the bricks that can be skipped, with the same remaining brick heights.
	for i := 0; i < len(bricks); i++ {
		expectedHeights := maps.Clone(brickHeights)
		delete(expectedHeights, bricks[i].name) // remove skipped

		actualHeights := calculateDroppedHeights(bricks, i) // recalculate

		if maps.Equal(expectedHeights, actualHeights) {
			sum++
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 5, result)
	} else {
		assert.Equal(t, 407, result)
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
