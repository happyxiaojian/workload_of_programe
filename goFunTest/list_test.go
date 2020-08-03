package goFunTest

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Len int
}

func PrintList(list *List){
	if nil == list || list.Head == nil {
		fmt.Printf("list if empty...\n")
	}
	for v := list.Head.Next; v != nil; v = v.Next {
		fmt.Printf("%d->", v.Val)
	}
	fmt.Printf("Nil\n")
}

func (l *List)AddHead(node *ListNode)bool{
	if l == nil || l.Head == nil{
		return false
	}
	cur := l.Head.Next
	node.Next = cur
	l.Head.Next = node
	l.Len++
	return true
}


func reverseList(head *ListNode) *ListNode {
	pre, cur := head, head.Next

	nxt := cur.Next
	pre.Next = nil

	for nxt != nil {
		cur.Next = pre
		pre = cur
		cur = nxt
		nxt = nxt.Next
	}

	return cur
}

func TestList(t *testing.T){
	list := &List{
		Head: new(ListNode),
		Len:  0,
	}

	node1 := &ListNode{Val:  1}
	node2 := &ListNode{Val:  2}
	node3 := &ListNode{Val:  3}
	node4 := &ListNode{Val:  4}
	node5 := &ListNode{Val:  5}
	node6 := &ListNode{Val:  6}

	list.Head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	PrintList(list)

	reverseList(list.Head)

	PrintList(list)

}
