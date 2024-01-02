package day22

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
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
	if b.start.x > b.end.x ||
		b.start.y > b.end.y ||
		b.start.z > b.end.z {
		panic("inverse brick")
	}

	res := make([]Voxel, 0, (b.end.x-b.start.x+1)*(b.end.y-b.start.y+1)*(b.end.z-b.start.z+1))

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
func calculateDroppedHeights(bricks []Brick, skip int, checkHeights map[string]int) map[string]int {
	heights := make(map[string]int, len(bricks))
	collisionMap, width := makeHeightmap(bricks)

	// Drop the bricks one by one.
	for i := 0; i < len(bricks); i++ {
		if i == skip {
			continue
		}

		b1 := &bricks[i]
		distance := 100000000

		for _, cube := range b1.cubes {
			b2 := collisionMap[cube.y*width+cube.x]
			h := b2.z + 1
			distance = min(distance, cube.z-h)
		}

		blockHeight := b1.start.z - distance

		// Check if the dropped height of b1 is the same as the expected height.
		if len(checkHeights) > 0 {
			if h, ok := checkHeights[b1.name]; !ok || h != blockHeight {
				return nil
			}
		}

		// Stored dropped height of b1.
		heights[b1.name] = blockHeight

		for _, cube := range b1.cubes {
			b2 := collisionMap[cube.y*width+cube.x]
			droppedZ := cube.z - distance

			if b2.name == "" || droppedZ > b2.z {
				collisionMap[cube.y*width+cube.x] = Voxel{b1.name, cube.x, cube.y, droppedZ}
			}
		}
	}

	return heights
}

// makeHeightmap returns a 1D grid of voxels.
func makeHeightmap(bricks []Brick) (map[int]Voxel, int) {
	// Find max of X.
	maxX := 0
	for _, brick := range bricks {
		maxX = max(maxX, brick.end.x)
	}
	width := maxX + 1

	// Create the 1D grid.
	heightmap := make(map[int]Voxel, width*width) // Assume square.

	return heightmap, width
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
	brickHeights := calculateDroppedHeights(bricks, -1, nil)
	sum := int64(0)
	eg := errgroup.Group{}

	// Count the bricks that can be skipped, with the same remaining brick heights.
	for i := 0; i < len(bricks); i++ {
		i := i

		eg.Go(func() error {
			newHeights := calculateDroppedHeights(bricks, i, brickHeights) // recalculate

			if newHeights != nil {
				atomic.AddInt64(&sum, 1)
			}

			return nil
		})
	}

	_ = eg.Wait()
	return int(sum)
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
