package day22

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"sort"
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
	id      int
	x, y, z int
}

type Brick struct {
	id     int
	startZ int
	endX   int
	cubes  []Voxel
}

// Helper functions.

func Cubes(id, startX, startY, startZ, endX, endY, endZ int) []Voxel {
	if startX > endX ||
		startY > endY ||
		startZ > endZ {
		panic("inverse brick")
	}

	res := make([]Voxel, 0, (endX-startX+1)*(endY-startY+1)*(endZ-startZ+1))

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			for z := startZ; z <= endZ; z++ {
				res = append(res, Voxel{id, x, y, z})
			}
		}
	}

	return res
}

// calculateDroppedHeights returns the heights of the bricks after dropping.
func calculateDroppedHeights(bricks []Brick, skip int, checkHeights []int) []int {
	heights := make([]int, len(bricks))
	collisionMap, width := makeHeightmap(bricks)

	// Drop the bricks one by one.
	for i := 0; i < len(bricks); i++ {
		if i == skip {
			continue
		}

		b1 := &bricks[i]
		distance := 100000000

		for _, cube := range b1.cubes {
			z := collisionMap[cube.y*width+cube.x]
			h := z + 1
			distance = min(distance, cube.z-h)
		}

		blockHeight := b1.startZ - distance

		// Check if the dropped height of b1 is the same as the expected height.
		if len(checkHeights) > 0 {
			if h := checkHeights[b1.id]; h > 0 && h != blockHeight {
				return nil
			}
		}

		// Stored dropped height of b1.
		heights[b1.id] = blockHeight

		for _, cube := range b1.cubes {
			z := collisionMap[cube.y*width+cube.x]
			droppedZ := cube.z - distance

			if droppedZ > z {
				collisionMap[cube.y*width+cube.x] = droppedZ
			}
		}
	}

	return heights
}

// makeHeightmap returns a 1D grid of voxels.
func makeHeightmap(bricks []Brick) ([]int, int) {
	// Find max of X.
	maxX := 0
	for _, brick := range bricks {
		maxX = max(maxX, brick.endX)
	}
	width := maxX + 1

	// Create the 1D grid.
	heightmap := make([]int, width*width) // Assume square.
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			heightmap[y*width+x] = 0
		}
	}

	return heightmap, width
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	bricks := make([]Brick, 0, 1300)
	nextId := 0
	for _, line := range lines {
		match := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '~' })
		brick := Brick{
			nextId,
			Atoi(match[2]),
			Atoi(match[3]),
			Cubes(
				nextId,
				Atoi(match[0]), Atoi(match[1]), Atoi(match[2]),
				Atoi(match[3]), Atoi(match[4]), Atoi(match[5]),
			)}

		nextId++
		bricks = append(bricks, brick)
	}

	// Sort from bottom to top.
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].startZ < bricks[j].startZ
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
		//part2()
	}
}
