package implements

import (
	"errors"
	"fmt"
	"mxudong.github.io/SkipLink"
)

// implement: valuePackingFunc func(interface{}) (*Sortable, error)
func NumberSortableImplementPacking(value interface{}) (*SkipLink.Sortable, error) {
	var number float64
	switch value.(type) {
	case uint64:
		number = float64(value.(uint64))
	case uint32:
		number = float64(value.(uint32))
	case uint16:
		number = float64(value.(uint16))
	case uint8:
		number = float64(value.(uint8))
	case uint:
		number = float64(value.(uint))
	case int64:
		number = float64(value.(int64))
	case int32:
		number = float64(value.(int32))
	case int16:
		number = float64(value.(int16))
	case int8:
		number = float64(value.(int8))
	case int:
		number = float64(value.(int))
	case float64:
		number = float64(value.(float64))
	case float32:
		number = float64(value.(float32))
	default:
		return nil, errors.New("Input value not a number ")
	}

	n := SkipLink.Sortable(&NumberSortableImplement{
		value: number,
	})

	return &n, nil
}

// For any number, tru to convert to float64, and return some res.
type NumberSortableImplement struct {
	value interface{}
}

func (n *NumberSortableImplement) IsLessThan(sortable SkipLink.Sortable) (bool, error) {
	nsi, ok := sortable.(*NumberSortableImplement)
	if !ok {
		return false, errors.New(fmt.Sprintf(SkipLink.SortableTypeError, "NumberSortableImplement"))
	}
	f1 := n.value.(float64)
	f2 := nsi.value.(float64)

	return f1 < f2, nil
}

func (n *NumberSortableImplement) IsEquals(sortable SkipLink.Sortable) (bool, error) {
	nsi, ok := sortable.(*NumberSortableImplement)
	if !ok {
		return false, errors.New(fmt.Sprintf(SkipLink.SortableTypeError, "NumberSortableImplement"))
	}
	return nsi.value == sortable.Value(), nil
}

func (n *NumberSortableImplement) Value() interface{} {
	return n.value
}
