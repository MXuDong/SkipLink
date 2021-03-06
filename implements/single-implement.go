package implements

import (
	"errors"
	"fmt"
	"mxudong.github.io/SkipLink"
)

type SingleSortableImplement struct {
	index uint64
	value interface{}
}

// InitSingleSortableImplement will create an SingleSortableImplement, and set base value.
// And this implement is default implement.
func InitSingleSortableImplement(index uint64, value interface{}) *SkipLink.Sortable {
	ssi := SkipLink.Sortable(&SingleSortableImplement{
		index: index,
		value: value,
	})

	return &ssi
}

func (s SingleSortableImplement) IsLessThan(sortable SkipLink.Sortable) (bool, error) {
	ssi, ok := sortable.(*SingleSortableImplement)
	if !ok {
		return false, errors.New(fmt.Sprintf(SkipLink.SortableTypeError, "SingleSortableImplement"))
	}

	return s.index < ssi.index, nil
}

func (s *SingleSortableImplement) IsEquals(sortable SkipLink.Sortable) (bool, error) {
	ssi, ok := sortable.(*SingleSortableImplement)
	if !ok {
		return false, errors.New(fmt.Sprintf(SkipLink.SortableTypeError, "SingleSortableImplement"))
	}

	return s.index == ssi.index, nil
}

func (s SingleSortableImplement) Value() interface{} {
	return s.value
}

func SingleSortableImplementPacking() func(interface{}) (*SkipLink.Sortable, error) {
	var counter uint64 = 0

	return func(i interface{}) (*SkipLink.Sortable, error) {
		ssi := SkipLink.Sortable(&SingleSortableImplement{
			index: counter,
			value: i,
		})
		counter++
		return &ssi, nil
	}
}
