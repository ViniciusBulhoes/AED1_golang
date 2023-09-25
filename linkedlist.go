package main

import "fmt"

type LinkedList struct {
	head *Node1P
	size int
}

type Node1P struct {
	value int
	next  *Node1P
}

func (linkedlist *LinkedList) Add(value int) {
	newNode := Node1P{value, nil}

	aux := linkedlist.head
	prev := aux
	//iterate until aux points to nil, and prev points to the last node
	for aux != nil {
		prev = aux
		aux = aux.next
	}

	if prev == nil {
		//head was pointing to nil
		linkedlist.head = &newNode
	} else {
		//head points to a node but prev.next needs to be updated
		prev.next = &newNode
	}
	linkedlist.size++
}

func (linkedlist *LinkedList) AddPos(pos int, value int) {
	newNode := Node1P{value, nil}
	if pos >= 0 && pos <= linkedlist.size {
		aux := linkedlist.head
		prev := aux
		//iterate until aux points to nil, and prev points to the last node
		for i := 0; i < pos; i++ {
			prev = aux
			aux = aux.next
		}

		if prev == nil {
			//head was pointing to nil
			linkedlist.head = &newNode
		} else if aux.next == nil {
			//head points to a node but prev.next needs to be updated
			prev.next = &newNode
		} else {
			prev.next = &newNode
			newNode.next = aux
		}
		linkedlist.size++
	}
}

func (linkedlist *LinkedList) Reverse() {
	aux1 := linkedlist.head
	for i := 0; i < linkedlist.size/2; i++ {
		aux2 := aux1
		aux3 := aux1.value
		for j := i; j < linkedlist.size-i-1; j++ {
			aux2 = aux2.next
		}
		aux1.value = aux2.value
		aux2.value = aux3
		aux1 = aux1.next
	}
}

func main() {
	newList := LinkedList{nil, 0}
	for i := 0; i < 10; i++ {
		newList.Add(i)
	}
	newList.Reverse()
	printedNode := newList.head
	for i := 0; i < 10; i++ {
		fmt.Println(printedNode.value)
		printedNode = printedNode.next
	}
}
