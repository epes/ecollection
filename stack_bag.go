package ecollection

import "sync"

type stackBag[A any] struct {
	stack   []A
	size    uint64
	pointer uint64
	mtx     sync.Mutex
}

func NewFullNumberStackBag[N Number](size uint64) Bag[N] {
	s := make([]N, size)
	for i := range s {
		s[i] = N(i)
	}

	return &stackBag[N]{
		stack:   s,
		size:    size,
		pointer: size,
	}
}

func NewStackBag[A any](size uint64) Bag[A] {
	s := make([]A, size)

	return &stackBag[A]{
		stack:   s,
		size:    size,
		pointer: size,
	}
}

func (b *stackBag[A]) Take() (A, error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.pointer <= 0 {
		var empty A
		return empty, ErrEmpty
	}

	b.pointer--

	return b.stack[b.pointer], nil
}

func (b *stackBag[A]) Put(item A) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.pointer >= b.size {
		return ErrFull
	}

	b.stack[b.pointer] = item
	b.pointer++

	return nil
}
