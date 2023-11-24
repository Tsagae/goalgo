package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet(t *testing.T) {
	mapSet := NewMapSet[int]()
	assert.Equal(t, 0, mapSet.Size())

	mapSet.Put(10)
	assert.Equal(t, 1, mapSet.Size())
	assert.Equal(t, true, mapSet.Find(10))

	mapSet.Put(10)
	mapSet.Remove(10)
	assert.Equal(t, 0, mapSet.Size())
	assert.Equal(t, false, mapSet.Find(10))

	mapSet.Remove(10)
	assert.Equal(t, 0, mapSet.Size())
	assert.Equal(t, false, mapSet.Find(10))

}

func TestMapSet_Union(t *testing.T) {
	mapSetA := NewMapSet[int]()
	mapSetB := NewMapSet[int]()
	mapSetA.Put(1)
	mapSetA.Put(2)
	mapSetA.Put(3)
	mapSetA.Put(4)
	mapSetA.Put(4)
	mapSetB.Put(5)
	mapSetB.Remove(5)

	mapSetB.Put(3)
	mapSetB.Put(4)
	mapSetB.Put(6)

	mapSetC := Union(mapSetA, mapSetB)
	assert.Equal(t, 5, mapSetC.Size())

	assert.Equal(t, true, mapSetC.Find(1))
	assert.Equal(t, true, mapSetC.Find(2))
	assert.Equal(t, true, mapSetC.Find(3))
	assert.Equal(t, true, mapSetC.Find(4))
	assert.Equal(t, true, mapSetC.Find(6))

	assert.Equal(t, false, mapSetC.Find(5))
	assert.Equal(t, false, mapSetC.Find(10))

}

func TestMapSet_Items(t *testing.T) {
	mapSet := NewMapSet[int]()
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mapSet.PutAll(items...)
	assert.ElementsMatch(t, items, mapSet.ToSlice())
}
