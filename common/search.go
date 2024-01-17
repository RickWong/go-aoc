package common

import (
	"github.com/zyedidia/generic/heap"
	"slices"
	"time"
)

type SearchResult[T any, W Number] struct {
	Best         *T
	BestWeight   W
	BestPath     []*T
	Paths        int
	Milliseconds int64
	Iterations   int
}

type heapItem[T any, W Number] struct {
	priority W
	weight   W
	branch   *T
}

func IterativeSearch[T any, H Hashable, W Number](
	root *T,
// branches are possible iterations based on the current branch.
	branchFn func(branch *T) []*T,
// predicate terminates the search when true.
	predicateFn func(branch *T) bool,
// identity is a map key that represents the unique iteration.
	identityFn func(branch *T) H,
// weight is the absolute weight of the branch.
	weightFn func(branch *T, parentWeight W) W,
// heuristic is a relative priority modifier.
	heuristicFn func(branch *T) W,
// beam width limits search space on each iteration.
	beamWidth int,
// returnFirst terminates the search after the first result.
	returnFirst bool,
// maximize true will search for the highest weight.
	maximize bool,
) *SearchResult[T, W] {
	result := &SearchResult[T, W]{nil, 0, nil, 0, 0, 0}
	now := time.Now().UnixMilli()

	lessFn := func(a *heapItem[T, W], b *heapItem[T, W]) bool { return a.priority < b.priority }
	queue := heap.New(lessFn)
	beam := heap.New(lessFn)

	trail := make(map[H]*T, 32)
	weights := make(map[H]W, 1024)

	queue.Push(&heapItem[T, W]{0, 0, root})
	for queue.Size() > 0 {
		result.Iterations++
		current, _ := queue.Pop()

		if predicateFn != nil && predicateFn(current.branch) {
			result.Paths++

			if result.Best == nil ||
				(maximize && current.weight > result.BestWeight) ||
				(!maximize && current.weight < result.BestWeight) {
				result.BestWeight = current.weight
				result.Best = current.branch

				if returnFirst {
					break
				}
			}

			continue
		}

		branches := branchFn(current.branch)
		for _, branch := range branches {
			if branch == nil {
				continue
			}

			weight := current.weight
			if weightFn != nil {
				weight = weightFn(branch, current.weight)
			}

			if identityFn != nil {
				id := identityFn(branch)
				knownWeight, known := weights[id]
				if known &&
					((maximize && knownWeight >= weight) ||
						(!maximize && knownWeight <= weight)) {
					continue
				}

				weights[id] = weight
				trail[id] = current.branch
			}

			priority := weight
			if heuristicFn != nil {
				priority += heuristicFn(branch)
			}

			if maximize {
				priority = -priority
			}

			if beamWidth > 0 {
				beam.Push(&heapItem[T, W]{priority, weight, branch})
			} else {
				queue.Push(&heapItem[T, W]{priority, weight, branch})
			}
		}

		for i := 0; i < beamWidth && beam.Size() > 0; i++ {
			item, _ := beam.Pop()
			queue.Push(item)
		}

		// Clear the rest of the beam.
		for beam.Size() > 0 {
			beam.Pop()
		}
	}

	if result.Best != nil && identityFn != nil {
		result.BestPath = append(result.BestPath, result.Best)
		nextStep := result.Best
		for nextStep != nil {
			nextStep = trail[identityFn(nextStep)]
			result.BestPath = append(result.BestPath, nextStep)
			if nextStep == root {
				break
			}
		}
		slices.Reverse(result.BestPath)
	}

	result.Milliseconds = time.Now().UnixMilli() - now
	return result
}
