package implements

import (
	"fmt"
	"mxudong.github.io/SkipLink"
	"testing"
)

func Test_main(t *testing.T) {
	s := SkipLink.InitSkipLink(SkipLink.DefaultMaxLevel, NumberSortableImplementPacking, SkipLink.GeneratorDefaultHasNextLevelFunc())
	s.AddValue(1)
	s.AddValue(123)
	fmt.Println(s.ToArray())
}
