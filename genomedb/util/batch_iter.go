package util

type BatchIter[T any] struct {
	curStep int
	stepNum int
	slice   []T
}

func NewBatchIter[T any](stepNum int, slice []T) BatchIter[T] {
	return BatchIter[T]{curStep: 0, stepNum: stepNum, slice: slice}
}

func (b *BatchIter[T]) Next() (next []T, finish bool) {
	if b.curStep > len(b.slice) {
		return next, true
	}

	if b.curStep+b.stepNum > len(b.slice) {
		next = b.slice[b.curStep:]
	} else {
		next = b.slice[b.curStep : b.curStep+b.stepNum]
	}

	b.curStep += b.stepNum
	return next, false
}
