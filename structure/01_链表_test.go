package structure

import (
	"testing"
)

func Test11(t *testing.T) {
	var link *OneLinkNode
	for i := 0; i < 5; i++ {
		if link == nil {
			link = &OneLinkNode{Val: i}
		} else {
			link = &OneLinkNode{Val: i, Next: link}
		}
	}
	for link != nil {
		t.Log(link.Val)
		link = link.Next
	}
}

// linkNode
type OneLinkNode struct {
	Val  int
	Next *OneLinkNode
}

func Test12(t *testing.T) {
	var link *TwoLinkNode
	for i := 0; i < 5; i++ {
		if link == nil {
			link = &TwoLinkNode{Val: i}
		} else {
			link = &TwoLinkNode{Val: i, Next: link}
		}
	}
	for link != nil {
		t.Log(link.Val)
		link = link.Next
	}
}

type TwoLinkNode struct {
	Val  int
	Pre  *TwoLinkNode
	Next *TwoLinkNode
}
