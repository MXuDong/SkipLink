package implements

import (
	"fmt"
	"mxudong.github.io/SkipLink"
	"testing"
)

func TestInitSingleSortableImplement(t *testing.T) {
	s := SkipLink.InitSkipLink(SkipLink.DefaultMaxLevel,
		SingleSortableImplementPacking(),
		SkipLink.GeneratorDefaultHasNextLevelFunc())

	s.Add("Te2st")
	s.Add(123)
	s.Add(123.2)
	s.Add(nil)

	fmt.Println(s.ToArray())
}
