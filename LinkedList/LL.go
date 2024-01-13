package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(data int) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	currentNode := ll.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
}

func (ll *LinkedList) insertAtFirst(data int) {
	newnode := &Node{data, nil}

	if ll.head != nil {
		newnode.next = ll.head
		ll.head = newnode
	}
}

func (ll *LinkedList) insertAfterSomeElement(data int, afterElemen int) {
	currentNode := ll.head
	for currentNode.data != 20 {
		currentNode = currentNode.next
	}
	fmt.Println("kk", currentNode.data)
	newnode := &Node{data, nil}

	newnode.next = currentNode.next
	currentNode.next = newnode
}

func (ll *LinkedList) Display() {
	currentNode := ll.head
	for currentNode.next != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println(currentNode.data)
}

func (ll *LinkedList) DeleteNodeWhichContainsGivenData(data int) {
	currentNode := ll.head

	for currentNode.next.data != data {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
}


func main() {
	list := LinkedList{}
	list.insert(10)
	list.insert(20)
	list.insert(30)
	list.insert(40)
	list.insertAtFirst(5)
	list.insertAfterSomeElement(25, 20)

	list.Display()

	fmt.Println("**********************")
	list.DeleteNodeWhichContainsGivenData(25)
	list.Display()

}
