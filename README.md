# Skip link implement by golang

跳跃表 golang 实现版本

# Usage:

```go
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
```