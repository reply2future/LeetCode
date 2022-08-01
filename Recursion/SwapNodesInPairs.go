package main
/**
 * Given a linked list, swap every two adjacent nodes and return its head.
 * You may not modify the values in the list's nodes, only nodes itself may be changed.
 * eg.
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func _swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
    var _beforeNode *ListNode
    var _nextNode *ListNode
    _midNode := head
    for {
        if _beforeNode != nil {
            _beforeNode.Next = _midNode.Next
        } else {
            head = head.Next
        }
        
        _nextNode = _midNode.Next
        _midNode.Next = _nextNode.Next
        _nextNode.Next = _midNode
        
        _beforeNode = _midNode
        _midNode = _beforeNode.Next
        
        if _midNode == nil || _midNode.Next == nil {
            break
        }
    }
    
    return head
}

// 递归做法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	_subHead := swapPairs(head.Next.Next)
	_nextNode := head.Next
	head.Next = _subHead
	_nextNode.Next = head
	return _nextNode	
}
