package SkipLink

import "testing"

func TestLinkNode_AddNodeToVPre(t *testing.T) {
	// pre nil test
	o11 := &LinkNode{
		level: 1,
	}
	o12 := &LinkNode{}
	o11.AddNodeToVPre(o12)
	if o11.vPre != o12 || o12.vNext != o11 {
		t.Errorf("Error for test")
	}
	// pre has value
	o21 := &LinkNode{
		level: 1,
	}
	o22 := &LinkNode{}
	o23 := &LinkNode{
		level: 0,
	}
	o21.vPre = o23
	o23.vNext = o21
	o21.AddNodeToVPre(o22)
	if o21.vPre != o22 || o22.vNext != o21 {
		t.Errorf("Error for test")
	}
	if o23.vNext != o22 || o22.vPre != o23 {
		t.Errorf("Error for test")
	}
}
