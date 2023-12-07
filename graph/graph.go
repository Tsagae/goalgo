package graph

type Weight interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 |
		~uint64 | ~float32 | ~float64
}

type Node[T comparable, W Weight] interface {
	GetEdges() []Edge[T, W]
	GetLabel() T
}

type Edge[T comparable, W Weight] interface {
	GetNodeFrom() Node[T, W]
	GetNodeTo() Node[T, W]
	GetWeight() W
}
