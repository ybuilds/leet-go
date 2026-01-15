package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) add(val int) *ListNode {
	if node == nil {
		return &ListNode{val, nil}
	}

	curr := node
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &ListNode{val, nil}
	return node
}

func (node *ListNode) display() {
	if node == nil {
		fmt.Println("List is empty")
	}

	curr := node
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
}
