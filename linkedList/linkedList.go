package main

import "fmt"

// Определим структуру односвязного списка
type ListNode struct {
	Next *ListNode
	Val  int
}

// Вывод односвязного списка
func (head *ListNode) Print() {
	curr := head
	for curr != nil {
		splitt := ""
		if curr != head {
			splitt = " -> "
		}
		fmt.Printf("%s%d", splitt, curr.Val)
		curr = curr.Next
	}
}

// Функция перевернутого списка
func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}
	var rev *ListNode
	curr := head

	for curr != nil {
		rev = &ListNode{
			Next: rev,
			Val:  curr.Val,
		}
		curr = curr.Next
	}
	return rev
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}

	n1.Next = n2
	n3 := &ListNode{Val: 3}
	n2.Next = n3

	n1.Print()
	fmt.Println()
	n1.Reverse().Print()
}
