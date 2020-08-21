package structures

import "fmt"

// 单向链表 Node节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//  List2Ints: covert list to []int
func List2Ints(head *ListNode) []int {
	limit := 100
	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// Ints2List: convert []int to list
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}

	return l.Next
}

func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if val == res.Next.Val {
			break
		}
		res = res.Next
	}
	return res
}

func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos --
	}

	tail := c
	if tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
