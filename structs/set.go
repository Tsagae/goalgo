package structs

type Set[T comparable] interface {
	Find(T) bool
	Put(T)
	Remove(T)
	Size() int
	Items() []T
}

type MapSet[T comparable] struct {
	innerMap map[T]bool
}

//TODO: quickunion set implementation
//TODO: iterator

func NewMapSet[T comparable]() MapSet[T] {
	return MapSet[T]{make(map[T]bool)}
}

func (m *MapSet[T]) Find(item T) bool {
	_, ok := m.innerMap[item]
	return ok
}

func (m *MapSet[T]) Put(item T) {
	m.innerMap[item] = true
}

func (m *MapSet[T]) Remove(item T) {
	delete(m.innerMap, item)
}

func (m *MapSet[T]) Size() int {
	return len(m.innerMap)
}

func (m *MapSet[T]) Items() []T {
	keys := make([]T, len(m.innerMap))

	i := 0
	for k := range m.innerMap {
		keys[i] = k
		i++
	}
	return keys
}

// Merges two sets, returns the new set. The two original sets are modified in the process of merging them
func Union[T comparable](a MapSet[T], b MapSet[T]) MapSet[T] {
	bigger := a
	smaller := b
	if len(a.innerMap) > len(b.innerMap) {
		bigger, smaller = smaller, bigger
	}
	for k := range smaller.innerMap {
		bigger.innerMap[k] = true
	}
	return bigger
}
