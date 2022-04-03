package ecollection

type Bag[A any] interface {
	Take() (A, error)
	Put(A) error
}
