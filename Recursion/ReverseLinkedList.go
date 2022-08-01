package main

/**
 * Given a linked list, swap every two adjacent nodes and return its head.
 * eg.
 * Given 1->2->3->4->NULL, you should return the list as 4->3->2->1->NULL.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func _swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	hd := head
	md := head.Next
	rd := md.Next

	head.Next = nil

	md.Next = head

	for rd != nil {
		hd = md
		md = rd
		rd = rd.Next

		md.Next = hd
	}
	head = md

	return head
}

// 官方遍歷解法
func _iterateSwap(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

// 递归做法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	_tmp := swapPairs(head.Next)
	head.Next.Next = head
	head.Next = nil
	return _tmp
}
