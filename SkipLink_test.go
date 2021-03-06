package SkipLink

import (
	"fmt"
	"testing"
)

type IntegerSortable struct {
	value int64
}

func (i *IntegerSortable) IsLessThan(valueI Sortable) (isLess bool, err error) {
	value, ok := (valueI.Value()).(int64)
	if !ok {
		return false, nil
	}
	return i.value < value, nil
}
func (i *IntegerSortable) IsEquals(valueI Sortable) (isEquals bool, err error) {
	value, ok := (valueI.Value()).(int64)
	if !ok {
		return false, nil
	}
	return i.value == value, nil
}

func (i *IntegerSortable) Value() interface{} {
	return i.value
}

func Test_main(t *testing.T) {

	var is int64 = 1

	f := func(x interface{}) {
		fmt.Println(float64(x.(int64)))
	}
	f(is)

	// if value packing func is nil, is dangerous
	s := InitSkipLink(8, nil, GeneratorDefaultHasNextLevelFunc())

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

	s.AddSortable(&v1s)
	s.AddSortable(&v2s)
	s.AddSortable(&v0s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v1s)
	s.AddSortable(&v8s)
	s.AddSortable(&v9s)
	s.AddSortable(&v7s)
	s.AddSortable(&v6s)

	allNode := s.GetAllSortable()
	for _, item := range allNode {
		for _, value := range item {
			fmt.Print("", value.Value(), " ")
		}
		fmt.Println()
	}

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())
	s.DeleteSortable(&v0s)

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())
	s.DeleteSortable(&v1s)

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())
	s.DeleteSortable(&v9s)

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())
	s.DeleteSortable(&v0s)

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())

	fmt.Println((s.Get(0)).Value())
	fmt.Println((s.Get(3)).Value())
	fmt.Println((s.Remove(3)).Value())
	fmt.Println(s.Remove(0).Value())

	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())

	s.DeleteSortable(&v2s)
	s.DeleteSortable(&v6s)
	s.DeleteSortable(&v7s)
	s.DeleteSortable(&v8s)
	fmt.Println(s.ToArray())
	fmt.Println(s.AllDataCount(), s.Length())
}
