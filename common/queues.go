package common

type Queue[V any] interface {
	Push(V)
	Pop() V
	Size() int
}

type BucketQueue[V any] struct {
	Buckets   [][]V
	size      int
	lookFirst int
}

func NewBucketQueue[V any](numBuckets int, bucketCapacity int) *BucketQueue[V] {
	bq := BucketQueue[V]{
		make([][]V, numBuckets),
		0,
		numBuckets + 1,
	}
	for prio := 0; prio < cap(bq.Buckets); prio++ {
		bq.Buckets[prio] = make([]V, 0, bucketCapacity)
	}
	return &bq
}

func (q *BucketQueue[V]) Push(prio int, v V) {
	if prio >= cap(q.Buckets) {
		numBuckets := (1 + (prio / cap(q.Buckets))) * cap(q.Buckets)
		bigger := make([][]V, numBuckets)
		bucketCapacity := cap(q.Buckets[0])
		for p := 0; p < cap(q.Buckets); p++ {
			bigger[p] = q.Buckets[p]
		}
		for p := cap(q.Buckets); p < cap(bigger); p++ {
			bigger[p] = make([]V, 0, bucketCapacity)
		}
		q.Buckets = bigger
	}
	q.Buckets[prio] = append(q.Buckets[prio], v)
	q.size++
	if prio < q.lookFirst {
		q.lookFirst = prio
	}
}

func (q *BucketQueue[V]) Pop() (V, int) {
	var v V
	if q.size > 0 {
		for prio, bucket := range q.Buckets[q.lookFirst:] {
			if len(bucket) > 0 {
				v = bucket[0]
				bucket[0] = bucket[len(bucket)-1]
				q.Buckets[prio] = bucket[:len(bucket)-1]
				q.size--
				if len(bucket) == 0 && q.lookFirst <= prio {
					q.lookFirst = prio + 1
				}
				return v, prio
			}
		}
	}
	return v, -1
}

func (q *BucketQueue[V]) Size() int {
	return q.size
}
