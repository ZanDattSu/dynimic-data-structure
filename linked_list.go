package main

import (
	"dynimic-data-structure/figures"
	"fmt"
)

type Node struct {
	value figures.Shape
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddFirst(s figures.Shape) *Node {
	node := &Node{value: s}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	return node
}

func (l *LinkedList) AddLast(s figures.Shape) *Node {
	node := &Node{value: s}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	return node
}

func (l *LinkedList) Remove(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}

	n.prev = nil
	n.next = nil
}

func (l *LinkedList) Print() {
	for node := l.head; node != nil; node = node.next {
		node.value.Draw()
	}
}

func (l *LinkedList) PrintAreas() {
	for node := l.head; node != nil; node = node.next {
		fmt.Printf("%s area = %.2f\n", node.value.GetName(), node.value.Area())
	}
}
