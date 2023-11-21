package common

import (
	heap2 "github.com/zyedidia/generic/heap"
	"slices"
	"time"
)

type SearchResult[T any] struct {
	Best         *T
	Path         []*T
	Milliseconds int64
	Iterations   int
}

type heapItem[T any] struct {
	priority float64
	weight   float64
	branch   *T
}

func IterativeSearch[T any](
	start *T,
	// branches are possible iterations based on the current.
	branchFn func(current *T) []*T,
	// predicate terminates the search when true.
	predicateFn func(current *T) bool,
	// identity is a map key that represents the unique iteration.
	identityFn func(current *T) any,
	// weightFn is a additive/cumulative weight.
	weightFn func(current *T) float64,
	// heuristicFn is an absolute weight modifier.
	heuristicFn func(current *T) float64,
	// beam width limits search space on each iteration.
	beamWidth int,
) *SearchResult[T] {
	result := &SearchResult[T]{nil, nil, 0, 0}
	now := time.Now().UnixMilli()

	lessFn := func(a heapItem[T], b heapItem[T]) bool { return a.priority < b.priority }
	heap := heap2.New(lessFn)
	beam := heap2.New(lessFn)

	var trail = make(map[any]*T)
	var weights = make(map[any]float64)

	heap.Push(heapItem[T]{0, 0, start})
	for heap.Size() > 0 {
		result.Iterations++
		current, _ := heap.Pop()

		if predicateFn(current.branch) {
			result.Best = current.branch
			break
		}

		for _, branch := range branchFn(current.branch) {
			weight := current.weight
			if weightFn != nil {
				weight += weightFn(branch)
			}

			var id any
			if identityFn != nil {
				id = identityFn(branch)
			}

			if id != nil {
				knownWeight, known := weights[id]
				if known && weight >= knownWeight {
					continue
				}

				weights[id] = weight
				trail[id] = current.branch
			}

			priority := weight
			if heuristicFn != nil {
				priority += heuristicFn(branch)
			}

			if beamWidth == 0 {
				heap.Push(heapItem[T]{priority, weight, branch})
				continue
			}

			if beam.Size() >= beamWidth {
				beam.Pop()
			}

			beam.Push(heapItem[T]{priority, weight, branch})
		}

		for beam.Size() > 0 {
			item, _ := beam.Pop()
			heap.Push(item)
		}
	}

	if result.Best != nil {
		result.Path = append(result.Path, result.Best)
		nextStep := result.Best
		for nextStep != nil {
			nextStep = trail[identityFn(nextStep)]
			result.Path = append(result.Path, nextStep)
			if nextStep == start {
				break
			}
		}
		slices.Reverse(result.Path)
	}

	result.Milliseconds = time.Now().UnixMilli() - now
	return result
}
