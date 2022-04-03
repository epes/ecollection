package ecollection

type Number interface {
	~int | ~int16 | ~int32 | ~uint8 | ~uint16 | ~uint32
}
