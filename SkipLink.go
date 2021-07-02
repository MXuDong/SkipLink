package SkipLink

type Sortable interface {
}

// NodeValue packed value in the LinkNode
type NodeValue interface {
	Key() Sortable      // to get key in the NodeValue, the key should implement the interface: Sortable
	Value() interface{} // to get value in the node value, can be every thing
}

// LinkNode is the unit that makes up SkipLink
type LinkNode struct {
	vPre  *LinkNode // vertical pre node
	hPre  *LinkNode // horizontal pre node
	vNext *LinkNode // vertical next node
	hNext *LinkNode // horizontal next node

	value NodeValue // the node value
	level uint64    // the node level in vertical
}

// ================ add node methods
// Add target node to l vPre, and set level -1
// if the level already in l is zero, will append fail
func (l *LinkNode) AddNodeToVPre(ln *LinkNode) bool {
	// if ln is nil, do nothing
	if ln == nil {
		// add fail
		return false
	}
	if l.level == 0 {
		// level limit
		return false
	}

	ln.vPre = l.vPre
	if l.vPre != nil {
		l.vPre.vNext = ln
	}
	ln.vNext = l
	l.vPre = ln
	ln.level = l.level - 1
	return true
}

// Add target node to l hPre, and set level
func (l *LinkNode) AddNodeToHPre(ln *LinkNode) bool {
	if ln == nil {
		return false
	}

	ln.hPre = l.hPre
	if l.hPre != nil {
		l.hPre.hNext = ln
	}
	ln.hNext = l
	l.hPre = ln
	ln.level = l.level
	return true
}

// Add target node to l vNext, and set level + 1
func (l *LinkNode) AddNodeToVNext(ln *LinkNode) bool {
	if ln == nil {
		return false
	}

	ln.vNext = l.vNext
	if l.vNext != nil {
		l.vNext.vPre = ln
	}
	ln.vPre = l
	l.vNext = ln
	ln.level = l.level + 1
	return true
}

// Add target node to l hNext, and set level
func (l *LinkNode) AddNodeToHNext(ln *LinkNode) bool {
	if ln == nil {
		return false
	}

	ln.hNext = l.hNext
	if l.hNext != nil {
		l.hNext.hPre = ln
	}
	ln.hPre = l
	l.hNext = ln
	ln.level = l.level
	return true
}

// invoke AddNodeMethod to add value
func createNodeAndTryToAdd(value NodeValue, f func(node *LinkNode) bool) (*LinkNode, bool) {
	ln := &LinkNode{
		value: value,
	}
	if f(ln) {
		return ln, true
	}
	return nil, false
}

// Add value to vPre, and if success, return be added node
func (l *LinkNode) AddValueToVPre(value NodeValue) (*LinkNode, bool) {
	return createNodeAndTryToAdd(value, l.AddNodeToVPre)
}

// Add value to hPre, and if success, return be added node
func (l *LinkNode) AddValueToHPre(value NodeValue) (*LinkNode, bool) {
	return createNodeAndTryToAdd(value, l.AddNodeToHPre)
}

// Add value to vNext, and if success, return be added node
func (l *LinkNode) AddValueToVNext(value NodeValue) (*LinkNode, bool) {
	return createNodeAndTryToAdd(value, l.AddNodeToVNext)
}

// Add value to hNext, and if success, return be added node
func (l *LinkNode) AddValueToHNext(value NodeValue) (*LinkNode, bool) {
	return createNodeAndTryToAdd(value, l.AddNodeToHNext)
}

// create a empty node and append to vPre, return itself
func (l *LinkNode) GeneratorEmptyToVPre() (*LinkNode, bool) {
	return l.AddValueToVPre(nil)
}

// create a empty node and append to hPre, return itself
func (l *LinkNode) GeneratorEmptyToHPre() (*LinkNode, bool) {
	return l.AddValueToHPre(nil)
}

// create a empty node and append to vNext, return itself
func (l *LinkNode) GeneratorEmptyToVNext() (*LinkNode, bool) {
	return l.AddValueToVNext(nil)
}

// create a empty node and append to hNext, return itself
func (l *LinkNode) GeneratorEmptyToHNext() (*LinkNode, bool) {
	return l.AddValueToHNext(nil)
}

// ================ delete node methods (delete itSelf)
func (l *LinkNode) Delete() {
	if l.vNext != nil {
		l.vNext.vPre = l.vPre
	}
	if l.vPre != nil {
		l.vPre.vNext = l.vNext
	}
	if l.hNext != nil {
		l.hNext.hPre = l.hPre
	}
	if l.hPre != nil {
		l.hPre.hNext = l.hNext
	}
}

// =============== search
func (l *LinkNode) VHead() *LinkNode {
	if l.vPre != nil {
		return l.vPre.VHead()
	}
	return l
}

func (l *LinkNode) VEnd() *LinkNode {
	if l.vNext != nil {
		return l.vNext.VEnd()
	}
	return l
}

func (l *LinkNode) HHead() *LinkNode {
	if l.hPre != nil {
		return l.hPre.HHead()
	}
	return l
}

func (l *LinkNode) HEnd() *LinkNode {
	if l.hNext != nil {
		return l.hNext.HEnd()
	}
	return l
}

// ================== value method
// GetNodeValue will return value of LinkNode
func (l *LinkNode) GetNodeValue() interface{} {
	if l.value == nil {
		return nil
	}
	return l.value.Value()
}

// GetNodeKey will return key of LinkNode, the Key should implement Sortable
func (l *LinkNode) GetNodeKey() Sortable {
	if l.value == nil {
		return nil
	}
	return l.value.Key()
}

func (l *LinkNode) Level() uint64 {
	return l.level
}
