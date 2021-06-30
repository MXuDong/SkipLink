package SkipLink

const (
	DefaultMaxLevel = 8
)

// Iterator return the iterator for SkipLink, it can be sync and not sync.
// For sync, it maybe use more memory.
type Iterator interface {
	hasNext() bool
	Next() (interface{}, error)
	Del() bool
}

type Sortable interface {
	IsLessThan(Sortable) (isLess bool)
	IsEquals(Sortable) (isEquals bool)
	Value() interface{}
}

// del will remove this single node, link next to pre node.
func (i *elementNode) del() {
	if i.pre != nil {
		i.pre.next = i.next
	}
	if i.next != nil {
		i.next.pre = i.pre
	}
}

// appendNext will append node to header of next-list
func (i *elementNode) appendNext(node *elementNode) *elementNode {
	node.pre = i
	node.next = i.next
	if i.next != nil {
		i.next.pre = node
	}
	i.next = node
	return node
}

// appendPre will append node to end of pre-list, if is header, it will append fail
func (i *elementNode) appendPre(node *elementNode) bool {
	if i.pre == nil {
		return false
	}
	node.pre = i.pre
	node.next = i
	if i.pre != nil {
		i.pre.next = node
	}
	i.pre = node
	return true
}

func (e *elementNode) findMinLevel() *elementNode {
	if e.parentNode == nil {
		return e
	}
	return e.parentNode.findMinLevel()
}

// elementNode provide vertical access
type elementNode struct {
	levelHeaderNode *elementNode
	childNode       *elementNode
	parentNode      *elementNode
	head            *elementNode // head node
	pre             *elementNode // pre node, head node has no pre node
	next            *elementNode // next node
	value           *Sortable    // the value of this node, header node has no value
	isUsed          bool         // to quick delete, but not delete value
	level           uint64
}

func (h *elementNode) addChildNode() *elementNode {
	hItem := elementNode{
		level:           h.level + 1,
		parentNode:      h,
		childNode:       nil,
		isUsed:          false,
		pre:             nil,
		next:            nil,
		levelHeaderNode: nil,
	}
	h.childNode = &hItem
	return &hItem
}

// find less node, and return min level node, if is equals, return now node and true, else return less node and false
func (h *elementNode) findLessNode(sortable *Sortable) (*elementNode, bool) {
	header := h
	for header.next == nil {
		if header.parentNode == nil {
			return h, false
		}
		header = header.parentNode
	}

	now := header.next
	var res *elementNode = header
	for now != nil {
		if (*now.value).IsLessThan(*sortable) {
			res = now
			now = now.next
			continue
		} else if (*now.value).IsEquals(*sortable) {
			res = now
			return res.findMinLevel(), true
		} else {
			if now.pre != nil {
				now = now.pre
				now = now.parentNode
			} else {
				return nil, false
			}
		}
	}
	return res.findMinLevel(), false
}

// SkipLink
type SkipLink struct {
	// maxLevel limit the max level of skip link, if maxLevel is zero, it is seem as 1
	// When maxLevel is 1, the skip will downgrade to the doubly link list.
	// this value can't change after init.
	// Default is DefaultMaxLevel
	maxLevel uint64
	// header is all link list header.
	header *elementNode

	// all valid data count
	elementCount uint64
	// all data count, more than elementCount
	allElementCount uint64

	// value packing func to packing the input value, and return the a sortable value
	valuePackingFunc func(interface{}) (*Sortable, error)
	hasNextLevel     func() bool
}

func (s *SkipLink) Length() uint64 {
	return s.elementCount
}
func (s *SkipLink) AllDataCount() uint64 {
	return s.allElementCount
}
func (s *SkipLink) MaxLevel() uint64 {
	return s.maxLevel
}
func (s *SkipLink) Add(sortable *Sortable) bool {
	if s.header == nil {
		s.header = &elementNode{
			level: 0,
		}
		s.allElementCount++
	}
	minHead := s.header.findMinLevel()

	node, isEquals := s.header.findLessNode(sortable)
	if isEquals {
		return false
	}

	beAppendNode := &elementNode{
		levelHeaderNode: minHead,
		childNode:       nil,
		parentNode:      nil,
		value:           sortable,
		isUsed:          true,
		head:            nil,
	}

	valueNode := node.findMinLevel().appendNext(beAppendNode)

	// graw add
	nowHead := minHead
	var nowLevel uint64 = 1
	for s.hasNextLevel() {
		if nowHead.childNode == nil {
			if nowLevel > s.maxLevel {
				break
			}
			nowHead = nowHead.addChildNode()
		} else {
			nowHead = nowHead.childNode
		}

		innerNowNode := nowHead
		innerNextNode := nowHead.next
		for innerNextNode != nil {
			if innerNowNode.value != nil {
				if !(*innerNowNode.value).IsLessThan(*sortable) {
					break
				}
			}
			innerNowNode = innerNextNode
			innerNextNode = innerNowNode.next
		}
		beAppendNode := &elementNode{
			levelHeaderNode: nowHead,
			childNode:       nil,
			parentNode:      valueNode,
			value:           sortable,
			isUsed:          true,
			head:            nil,
		}
		valueNode.childNode = beAppendNode
		innerNowNode.appendNext(beAppendNode)
		valueNode = valueNode.childNode
	}
	return true
}
func (s *SkipLink) Delete() bool {
	return false
}

func (s *SkipLink) ToArray() []interface{} {
	value := []interface{}{}

	minHead := s.header.findMinLevel()
	for minHead != nil {
		if minHead.value != nil {
			value = append(value, (*minHead.value).Value())
		}
		minHead = minHead.next
	}
	return value
}

//func (s *SkipLink) Iterator() Iterator {
//
//}
//func (s *SkipLink) SyncIterator() Iterator {
//
//}
