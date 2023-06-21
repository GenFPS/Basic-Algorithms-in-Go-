package main

import "fmt"

// Определим структуру односвязного списка
type LinkNode struct {
	value int
	next  *LinkNode
}

// Определим структуру двусвязного спписка
type DoubLinkNode struct {
	value int
	next  *DoubLinkNode
	prev  *DoubLinkNode
}

// Создадим интерфейс FuncLinkNode для наших структур
type FuncLinkNode interface {
	Print()
}

// Реализация метода Print() для структуры LinkNode
func (n *LinkNode) Print() {
	cur := n
	splitter := ""
	for cur != nil {
		if cur != n {
			splitter = " -> "
		}
		fmt.Printf("%s%d", splitter, cur.value)
		cur = cur.next
	}
	fmt.Println()
}

// Реализация метода Print() для структуры DoubLinkNode
func (n *DoubLinkNode) Print() {
	cur := n
	splitter := ""
	if cur.prev == nil {
		for cur.next != nil {
			if cur != n {
				splitter = " -> "
			}
			fmt.Printf("%s%d", splitter, cur.value)
			cur = cur.next
		}
		if cur.next == nil {
			for cur.prev != nil {
				if cur != n {
					splitter = " -> "
				}
				fmt.Printf("%s%d", splitter, cur.value)
				cur = cur.prev
			}
		}
		if cur.prev == nil {
			fmt.Printf("%s%d", splitter, cur.value)
		}
	}
}

func main() {
	// приведем пример инициализации односвязного списка
	headSingleN1 := &LinkNode{value: 1}

	sN2 := &LinkNode{value: 2}
	headSingleN1.next = sN2

	sN3 := &LinkNode{value: 3}
	sN2.next = sN3

	headSingleN1.Print() // Вывод: 1 -> 2 -> 3

	// приведем пример инициализации двусвязного списка
	var headDoubleN1, dN2, dN3 DoubLinkNode
	headDoubleN1 = DoubLinkNode{value: 1, next: &dN2, prev: nil}

	dN2 = DoubLinkNode{value: 2, next: &dN3, prev: &headDoubleN1}

	dN3 = DoubLinkNode{value: 3, next: nil, prev: &dN2}

	headDoubleN1.Print() // Вывод: 1 -> 2 -> 3 -> 2 -> 1
}
