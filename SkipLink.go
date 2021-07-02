package SkipLink

import (
	"fmt"
	"math/rand"
	"time"
)

const ()

// In the node, all value must packing to this struct.
// Sortable is one data of the SkipLink.
//
// Different implementations can form different data structures, such as stacks and queues.
// If the implementations strictly follow the requirements to achieve the sorting ability.
//
// The head is right of all data.
type Sortable interface {

	// If the source data should appear on the right side of the target data
	// Such like 3 < 4, so in the list is 'head ... 3 ... 4 ... '
	// Example:
	// header33 - > node31 -----------> node33
	//    |           |                   |
	// header22 - > node21 -> node22 -> node23
	//    |           |         |         |
	// header11 - > node11 -> node12 -> node13 -> node14
	// the node11 should less than node 12
	//
	// If has error(such like type not equals), the value will ignore
	IsLessThan(Sortable) (isLess bool, err error)

	// If the target data can't append to the SkipLink, return true.
	// The inner implement, if datas are equal, only one will save.
	//
	// If has error(such like type not equals), the value will ignore
	IsEquals(Sortable) (isEquals bool, err error)

	// Return the target value.
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
// Such like link: header1 -> node1 -> node3, invoke : node1.appendNext(node2)
//           res : header1 -> node1 -> node2 -> node3
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
// Such like link : header1 -> node2 -> node3, invoke : node2.appendPre(node1)
//           res  : header1 -> node1 -> node2 -> node3
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

// findMinLevel will return the now node's min level node.
// Example:
// header33 - > node31 -----------> node33
//    |           |                   |
// header22 - > node21 -> node22 -> node23
//    |           |         |         |
// header11 - > node11 -> node12 -> node13 -> node14
//
// the header33's min level node is header11, the node33's min level node is node13
func (e *elementNode) findMinLevel() *elementNode {
	if e.parentNode == nil {
		return e
	}
	return e.parentNode.findMinLevel()
}

// elementNode provide vertical access, the link node base.
// The SkipLink's data will packing into elementNode.value.
type elementNode struct {
	levelHeaderNode *elementNode // now level header
	childNode       *elementNode // the vertical next node
	parentNode      *elementNode // the vertical pre node
	head            *elementNode // head node
	pre             *elementNode // pre node, head node has no pre node
	next            *elementNode // next node
	value           *Sortable    // the value of this node, header node has no value
	isUsed          bool         // to quick delete, but not delete value
	level           uint64
}

// createChildNode will create empty node and append source node.
// Such like:
// header2
//    |
// header1
//
// invoke the header2's createChildNode
//
// header3
//    |
// header2
//    |
// header1
//
// Note that, the createChildNode will cover origin node, it is dangerous.
func (h *elementNode) createChildNode() *elementNode {
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
	var res = header
	for now != nil {
		isLessThan, err := (*now.value).IsLessThan(*sortable)
		if err != nil {
			return nil, false
		}
		isEquals, err := (*now.value).IsEquals(*sortable)
		if err != nil {
			return nil, false
		}
		if isLessThan {
			res = now
			now = now.next
			continue
		} else if isEquals {
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

// SkipLink, the skip link can use as Stack, queue, Single link by different Sortable-Implementations.
// The maxLevel limit the high of SkipLink.
// All of the SkipLink func is not sync, it mean that dangerous in a multi-threaded environment.
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

	// This function determines whether the current data should continue to grow，but the grow high be limit by
	// maxLevel, if there is no clear requirement in the implementation to formulate the return data, please try
	// to use the default function: GeneratorDefaultHasNextLevelFunc() -> func() bool
	hasNextLevel func() bool
}

// GeneratorDefaultHasNextLevelFunc return a func to return random result in bool (true and false)
func GeneratorDefaultHasNextLevelFunc() func() bool {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return func() bool {
		// random result
		return r.Int()%2 == 0
	}
}

func DefaultSkipLink() SkipLink {
	return InitSkipLink(DefaultMaxLevel, nil, GeneratorDefaultHasNextLevelFunc())
}

func InitSkipLink(maxLevel uint64, valuePackingFunc func(interface{}) (*Sortable, error), hasNextLevel func() bool) SkipLink {
	return SkipLink{
		maxLevel:         maxLevel,
		valuePackingFunc: valuePackingFunc,
		hasNextLevel:     hasNextLevel,
		header: &elementNode{
			level: 0,
		},
	}
}

// Length return the number of valid data
func (s *SkipLink) Length() uint64 {
	return s.elementCount
}

// AllDataCount return the number of all node, it is great than Length()
func (s *SkipLink) AllDataCount() uint64 {
	return s.allElementCount
}

// MaxLevel return the SkipLink's max level
func (s *SkipLink) MaxLevel() uint64 {
	return s.maxLevel
}

// Add the value, if not inert into the skip link, return false
func (s *SkipLink) Add(sortable *Sortable) bool {
	if s.header == nil {
		s.header = &elementNode{
			level: 0,
		}
		s.allElementCount++
	}
	minHead := s.header.findMinLevel()

	node, isEquals := s.header.findLessNode(sortable)
	// equals || type not equals
	if isEquals || node == nil {
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
	s.allElementCount++
	s.elementCount++

	// graw add
	nowHead := minHead
	var nowLevel uint64 = 1
	for s.hasNextLevel() {
		nowLevel++
		if nowHead.childNode == nil {
			if nowLevel >= s.maxLevel {
				break
			}
			nowHead = nowHead.createChildNode()
			s.allElementCount++

		} else {
			nowHead = nowHead.childNode
		}

		innerNowNode := nowHead
		innerNextNode := nowHead.next
		for innerNextNode != nil {
			if innerNextNode.value != nil {
				isLessThan, err := (*innerNextNode.value).IsLessThan(*sortable)
				// the type not equals, return false
				if err != nil {
					return false
				}

				if !isLessThan {
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
		s.allElementCount++
	}
	return true
}

// AddValue like add func, it will packing the input value to sortable by init param: valuePackingFunc
func (s *SkipLink) AddValue(value interface{}) bool {
	v, err := s.valuePackingFunc(value)
	if err != nil {
		return false
	}

	return s.Add(v)
}

// Delete the value, if not find into the skip link, return false
func (s *SkipLink) Delete(sortable *Sortable) bool {
	node, ok := s.header.findLessNode(sortable)
	if !ok {
		return false
	}
	s.elementCount--
	nowNode := node
	for nowNode != nil {
		tmp := nowNode.childNode
		nowNode.del()
		nowNode = tmp
		s.allElementCount--
	}

	return false
}

// DeleteValue will delete target value, it will packing the input value to sortable by init param: valuePackingFunc
func (s *SkipLink) DeleteValue(value interface{}) bool {
	v, err := s.valuePackingFunc(value)
	if err != nil {
		return false
	}

	return s.Delete(v)
}

// Get return the value which index is equals
func (s *SkipLink) Get(index uint64) Sortable {
	if index >= s.elementCount {
		panic(fmt.Sprintf("Range out of index for : %d", index))
	}

	nowNode := s.header.findMinLevel().next
	for index > 0 {
		index--

		nowNode = nowNode.next
	}

	return *nowNode.value
}

// Remove will remove the target value
func (s *SkipLink) Remove(index uint64) Sortable {
	if index >= s.elementCount {
		panic(fmt.Sprintf("Range out of index for : %d", index))
	}

	nowNode := s.header.findMinLevel().next
	for index > 0 {
		index--
		nowNode = nowNode.next
	}
	value := nowNode.value
	for nowNode != nil {
		s.allElementCount--
		nowNode.del()
		nowNode = nowNode.childNode
	}
	s.elementCount--
	return *value
}

// ToSortableArray return the array of Sortable
func (s *SkipLink) ToSortableArray() []Sortable {
	value := []Sortable{}

	minHead := s.header.findMinLevel()
	for minHead != nil {
		if minHead.value != nil {
			value = append(value, *minHead.value)
		}
		minHead = minHead.next
	}
	return value
}

// ToArray return all the value
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

// Get all sortable, return all node as array.
func (s *SkipLink) GetAllSortable() [][]Sortable {
	var reses [][]Sortable

	nowHead := s.header.findMinLevel()
	for nowHead != nil {
		value := []Sortable{}
		minHead := nowHead
		for minHead != nil {
			if minHead.value != nil {
				value = append(value, *minHead.value)
			}
			minHead = minHead.next
		}
		reses = append(reses, value)
		nowHead = nowHead.childNode
	}

	return reses
}
