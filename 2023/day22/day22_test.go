package day22

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"runtime"
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
	x, y, z int
}

type Brick struct {
	id     int
	startZ int
	endX   int
	voxels []Voxel
}

// Helper functions.
func parseBricks() []Brick {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	bricks := make([]Brick, 0, 1300)
	nextId := 0
	for _, line := range lines {
		match := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '~' })
		brick := Brick{
			nextId,
			Atoi(match[2]),
			Atoi(match[3]),
			Voxels(
				Atoi(match[0]), Atoi(match[1]), Atoi(match[2]),
				Atoi(match[3]), Atoi(match[4]), Atoi(match[5]),
			)}
		bricks = append(bricks, brick)
		nextId++
	}

	// Sort from bottom to top.
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].startZ < bricks[j].startZ
	})

	return bricks
}

// Voxels returns a list of voxels. The noinline directive is used for speed up.
func Voxels(startX, startY, startZ, endX, endY, endZ int) []Voxel {
	res := make([]Voxel, (endX-startX+1)*(endY-startY+1)*(endZ-startZ+1))
	i := 0

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			for z := startZ; z <= endZ; z++ {
				res[i] = Voxel{x, y, z}
				i++
			}
		}
	}

	return res
}

// calculateDroppedHeights returns the heights of the bricks after dropping.
func calculateDroppedHeights(bricks []Brick, collisionMap []int, width, skip int, checkHeights []int) []int {
	heights := make([]int, len(bricks))
	for i := range collisionMap {
		collisionMap[i] = 0
	}

	for i, b1 := range bricks {
		if i == skip {
			continue
		}

		dropZ := 100000000

		for _, voxel := range b1.voxels {
			currentZ := collisionMap[voxel.y*width+voxel.x]
			dropZ = min(dropZ, voxel.z-(currentZ+1))
		}

		blockHeight := b1.startZ - dropZ

		// Check if the dropped height of b1 is the same as the expected height.
		if len(checkHeights) > 0 {
			if h := checkHeights[b1.id]; h > 0 && h != blockHeight {
				// Height difference not expected.
				return nil
			}
		}

		// Stored dropped height of b1.
		heights[b1.id] = blockHeight

		// Update collision map.
		for _, voxel := range b1.voxels {
			collisionMap[voxel.y*width+voxel.x] = voxel.z - dropZ
		}
	}

	return heights
}

// makeHeightmap returns a 1D grid of voxels.
func makeHeightmap(bricks []Brick) ([]int, int) {
	maxX := 0
	for _, brick := range bricks {
		maxX = max(maxX, brick.endX)
	}

	width := maxX + 1
	return make([]int, width*width), width
}

// Part 1.

func part1() int {
	bricks := parseBricks()

	// Calculate the heights of the bricks after dropping.
	collisionMap, width := makeHeightmap(bricks)
	brickHeights := calculateDroppedHeights(bricks, collisionMap, width, -1, nil)
	sum := atomic.Int64{}
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())
	pageSize := len(bricks)/runtime.NumCPU() + 1

	for i := 0; i < runtime.NumCPU(); i++ {
		start := i * pageSize
		end := min(start+pageSize, len(bricks))

		eg.Go(func() error {
			collisionMap, width := makeHeightmap(bricks)
			for j := start; j < end; j++ {
				if calculateDroppedHeights(bricks, collisionMap, width, j, brickHeights) != nil {
					sum.Add(1)
				}
			}

			return nil
		})
	}

	_ = eg.Wait()
	return int(sum.Load())
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
	bricks := parseBricks()

	// Calculate the heights of the bricks after dropping.
	collisionMap, width := makeHeightmap(bricks)
	brickHeights := calculateDroppedHeights(bricks, collisionMap, width, -1, nil)
	sum := atomic.Int64{}
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())
	pageSize := len(bricks)/runtime.NumCPU() + 1

	for i := 0; i < runtime.NumCPU(); i++ {
		start := i * pageSize
		end := min(start+pageSize, len(bricks))

		eg.Go(func() error {
			collisionMap, width := makeHeightmap(bricks)
			for j := start; j < end; j++ {
				newHeights := calculateDroppedHeights(bricks, collisionMap, width, j, nil)

				changed := 0
				for k := 0; k < len(bricks); k++ {
					if newHeights[k] > 0 && newHeights[k] != brickHeights[k] {
						changed++
					}
				}

				sum.Add(int64(changed))
			}

			return nil
		})
	}

	_ = eg.Wait()
	return int(sum.Load())
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 7, result)
	} else {
		assert.Equal(t, 59266, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
