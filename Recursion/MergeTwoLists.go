package main

/**
 * Definition for singly-linked list.
 * https://leetcode.com/explore/learn/card/recursion-i/253/conclusion/2382/
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 不影響原來的兩個隊列
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return copyLists(l2)
	}
	if l2 == nil {
		return copyLists(l1)
	}

	var _head *ListNode
	var _pos *ListNode
	_l1Pos := l1
	_l2Pos := l2
	for {
		if _l1Pos.Val < _l2Pos.Val {
			if _head == nil {
				_head = &ListNode{Val: _l1Pos.Val}
				_pos = _head
			} else {
				_pos.Next = &ListNode{Val: _l1Pos.Val}
				_pos = _pos.Next
			}
			_l1Pos = _l1Pos.Next
		} else if _l1Pos.Val > _l2Pos.Val {
			if _head == nil {
				_head = &ListNode{Val: _l2Pos.Val}
				_pos = _head
			} else {
				_pos.Next = &ListNode{Val: _l2Pos.Val}
				_pos = _pos.Next
			}
			_l2Pos = _l2Pos.Next
		} else {
			if _head == nil {
				_head = &ListNode{Val: _l1Pos.Val, Next: &ListNode{Val: _l2Pos.Val}}
				_pos = _head.Next
			} else {
				_pos.Next = &ListNode{Val: _l1Pos.Val, Next: &ListNode{Val: _l2Pos.Val}}
				_pos = _pos.Next.Next
			}
			_l1Pos = _l1Pos.Next
			_l2Pos = _l2Pos.Next
		}
		if _l1Pos == nil {
			_pos.Next = copyLists(_l2Pos)
			break
		}
		if _l2Pos == nil {
			_pos.Next = copyLists(_l1Pos)
			break
		}
	}

	return _head
}

func copyLists(source *ListNode) *ListNode {
	var _ret *ListNode
	var _retPos *ListNode
	_pos := source
	for ; _pos != nil; _pos = _pos.Next {
		_tmp := &ListNode{Val: _pos.Val}
		if _ret == nil {
			_ret = _tmp
			_retPos = _ret
		} else {
			_retPos.Next = _tmp
			_retPos = _tmp
		}
	}
	return _ret
}

func _mergeRecursiveLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	_head := new(ListNode)
	if l1.Val < l2.Val {
		_head.Val = l1.Val
		l1 = l1.Next
	} else {
		_head.Val = l2.Val
		l2 = l2.Next
	}
	_head.Next = _mergeRecursiveLists(l1, l2)
	return _head
}
