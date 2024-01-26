package common

import (
	deque "github.com/edwingeng/deque/v2"
)

type Queue[V any] interface {
	Push(V)
	Pop() V
	Size() int
}

type BucketQueue[V any] struct {
	Buckets        []*deque.Deque[V]
	size           int
	bucketCapacity int
	lookFirst      int
}

func NewBucketQueue[V any](numBuckets int, bucketCapacity int) *BucketQueue[V] {
	bq := BucketQueue[V]{
		make([]*deque.Deque[V], numBuckets),
		0,
		bucketCapacity,
		numBuckets + 1,
	}
	for prio := 0; prio < numBuckets; prio++ {
		bq.Buckets[prio] = deque.NewDeque[V](deque.WithChunkSize(bucketCapacity))
	}
	return &bq
}

func (q *BucketQueue[V]) Push(prio int, v V) {
	if prio >= cap(q.Buckets) {
		numBuckets := (1 + (prio / cap(q.Buckets))) * cap(q.Buckets)
		bigger := make([]*deque.Deque[V], numBuckets)
		for p := 0; p < cap(q.Buckets); p++ {
			bigger[p] = q.Buckets[p]
		}
		for p := cap(q.Buckets); p < cap(bigger); p++ {
			bigger[p] = deque.NewDeque[V](deque.WithChunkSize(q.bucketCapacity))
		}
		q.Buckets = bigger
	}
	q.Buckets[prio].PushBack(v)
	q.size++
	if prio < q.lookFirst {
		q.lookFirst = prio
	}
}

func (q *BucketQueue[V]) Pop() (V, int) {
	var v V
	if q.size > 0 {
		for prio := q.lookFirst; prio < cap(q.Buckets); prio++ {
			bucket := q.Buckets[prio]
			if !bucket.IsEmpty() {
				v = bucket.PopBack()
				q.size--
				if bucket.IsEmpty() && q.lookFirst <= prio {
					q.lookFirst = prio + 1
				}
				return v, prio
			} else if q.lookFirst < prio {
				q.lookFirst = prio
			}
		}
	}
	return v, -1
}

func (q *BucketQueue[V]) Size() int {
	return q.size
}
