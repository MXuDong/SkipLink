package SkipLink

import (
	"fmt"
	"math/rand"
	"testing"
)

type IntegerSortable struct {
	value int64
}

func (i *IntegerSortable) IsLessThan(valueI Sortable) (isLess bool) {
	value, ok :=(valueI.Value()).(int64)
	if !ok {
		return false
	}
	return i.value < value
}
func (i *IntegerSortable) IsEquals(valueI Sortable) (isEquals bool) {
	value, ok :=(valueI.Value()).(int64)
	if !ok {
		return false
	}
	return i.value == value
}

func (i *IntegerSortable) Value() interface{} {
	return i.value
}

func Test_main(t *testing.T) {
	s := SkipLink{
		maxLevel: 1,
		hasNextLevel: func() bool {
			return rand.Int()%2 == 0
		},
	}

	v0 := IntegerSortable{value: 0}
	v1 := IntegerSortable{value: 1}
	v2 := IntegerSortable{value: 2}
	v8 := IntegerSortable{value: 8}
	v9 := IntegerSortable{value: 9}
	v7 := IntegerSortable{value: 7}
	v6 := IntegerSortable{value: 6}

	v0s := Sortable(&v0)
	v1s := Sortable(&v1)
	v2s := Sortable(&v2)
	v8s := Sortable(&v8)
	v9s := Sortable(&v9)
	v7s := Sortable(&v7)
	v6s := Sortable(&v6)

	s.Add(&v1s)
	s.Add(&v2s)
	s.Add(&v0s)
	s.Add(&v1s)
	s.Add(&v8s)
	s.Add(&v9s)
	s.Add(&v7s)
	s.Add(&v6s)

	fmt.Println(s.ToArray())
	s.Delete(&v0s)


	fmt.Println(s.ToArray())
	s.Delete(&v1s)

	fmt.Println(s.ToArray())
	s.Delete(&v9s)

	fmt.Println(s.ToArray())
	s.Delete(&v0s)

	fmt.Println(s.ToArray())
}
