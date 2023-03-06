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

func (l *SinglyLinkedList) InsertToLastNode(data int) {

	tempNodeForIteration := l.Head
	tempNode := CreateNode(data)

	for tempNodeForIteration.Next != nil {
		tempNodeForIteration = tempNodeForIteration.Next
	}

	tempNodeForIteration.Next = tempNode
}

// InsertToAfterTargetNode Assume LinkedList Has unique Datum.
func (l *SinglyLinkedList) InsertToAfterTargetNode(specificTargetData int, data int) {

	tempNodeForIteration := l.Head
	tempNode := CreateNode(data)

	//Change Mid
	for tempNodeForIteration.Next != nil {

		if tempNodeForIteration.Data == specificTargetData {
			tempNode.Next = tempNodeForIteration.Next
			tempNodeForIteration.Next = tempNode
			return
		}
		tempNodeForIteration = tempNodeForIteration.Next
	}

	//Change Last
	if tempNodeForIteration.Data == specificTargetData {
		l.InsertToLastNode(data)
		return
	}

	fmt.Printf("There is no target data: %d in this linkedList:", specificTargetData)
}
