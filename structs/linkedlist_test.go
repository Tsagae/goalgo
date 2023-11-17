package structs

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	if list.Size() != 0 {
		t.Errorf("Incorrect starting size")
	}

	for i := 5; i < 10; i++ {
		list.Add(i)
	}

	if list.Size() != 5 {
		t.Errorf("Incorrect size after Add")
	}

	fmt.Println(list.ToString())
	list.Insert(2, 99)

	if list.Get(2) != 99 {
		t.Errorf("Incorrect Get or Insert")
	}

	if list.Size() != 6 {
		t.Errorf("Incorrect size after Insert")
	}

	list.Remove(2)

	if list.Get(2) != 7 {
		t.Errorf("Incorrect Get or Remove")
	}

	if list.Size() != 5 {
		t.Errorf("Incorrect size after Remove")
	}

}
