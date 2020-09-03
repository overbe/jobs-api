package linkedList

import (
	it "jobs/internal/platform/datastructure/iterator"
	"testing"
)

func TestAppend(t *testing.T) {
	li := New(nil, nil)
	li.Append(32, "", "")
	li.Append(50, "", "")
	li.Append(100, "", "")

	head := func() int {
		return li.Head().Item()
	}

	if head() != 32 {
		t.Errorf("got %v, expected %v  \n", head(), 32)
	}

	li.RemoveFront()

	if head() != 50 {
		t.Errorf("got %v, expected %v  \n", head(), 50)
	}

	last := li.Head().Item()
	var iter it.IIterator = li.Iterator()
	for iter.HasNext() {
		last = iter.Next().Item()
	}
	if last != 100 {
		t.Errorf("iterator error got %v , expected %v \n", last, 100)
	}
}

func TestRemoveFront(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	li2 := New(nil, nil)

	if li2.RemoveFront() != nil {
		t.Errorf("should be nil")
	}

	li := New(nil, nil)
	for _, number := range numbers {
		li.Append(number, "", "")
	}
	li.RemoveFront()

	// New Head
	if got := li.Head().Item(); got != 2 {
		t.Errorf("got %v , expected %v \n", got, 2)
	}
	last := li.Head().Item()
	var iter it.IIterator = li.Iterator()
	for iter.HasNext() {
		last = iter.Next().Item()
	}
	if last != 6 {
		t.Errorf("iterator error got %v , expected %v \n", last, 6)
	}
}
