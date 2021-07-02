package implements

import (
	"fmt"
	"mxudong.github.io/SkipLink"
	"testing"
)

func Test_main(t *testing.T) {
	s := SkipLink.InitSkipLink(SkipLink.DefaultMaxLevel, NumberSortableImplementPacking, SkipLink.GeneratorDefaultHasNextLevelFunc())
	s.Add(1)
	s.Add(123)
	fmt.Println(s.ToArray())
}
