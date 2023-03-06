package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
}

func CreateNode(data int) *Node {

	node := Node{
		Data: data,
		Next: nil,
	}

	return &node
}

func CreateSinglyLinkedList(headData int) *SinglyLinkedList {

	linkedList := SinglyLinkedList{
		Head: CreateNode(headData),
	}

	return &linkedList
}

func (l *SinglyLinkedList) PrintLinkedListToScreen() {

	tempNode := l.Head

	for tempNode != nil {
		fmt.Printf("%d\t", tempNode.Data)
		tempNode = tempNode.Next
	}
	fmt.Printf("\n")
}

func (l *SinglyLinkedList) InsertToHeadNode(data int) {

	tempNode := CreateNode(data)
	tempNode.Next = l.Head
	l.Head = tempNode
}
