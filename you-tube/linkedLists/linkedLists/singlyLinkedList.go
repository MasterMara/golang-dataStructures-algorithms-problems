package linkedlists

import (
	"fmt"
)

// Define Node
type Node struct {
	Data int
	Next *Node
}

// Define SinglyLinkedList
type SinglyLinkedList struct {
	Head *Node
}

// Create Node
func CreateNode(nodeData int) *Node {

	myNode := Node{
		Data: nodeData,
		Next: nil,
	}

	return &myNode
}

// Create SinglyLinkedList
func CreateSinglyLinkedList(headNodeData int) *SinglyLinkedList {

	myLinkedList := SinglyLinkedList{
		Head: CreateNode(headNodeData),
	}

	return &myLinkedList
}

// Print SinglyLinkedList values.
func (l *SinglyLinkedList) PrintSinglyLinkedListToScreen() {

	// if List is empty
	if l.Head == nil {
		fmt.Println("Singly Linked List is empty.")
	}

	tempNode := l.Head

	for tempNode != nil {
		fmt.Printf("%d\t", tempNode.Data)
		tempNode = tempNode.Next
	}

	//For good formatting.
	fmt.Printf("\n")
}

// Bonus function for you :)
func PrintReverseSinglyLinkedListToScreen(node *Node) {

	if node != nil {
		PrintReverseSinglyLinkedListToScreen(node.Next)
		fmt.Printf("%d\t", node.Data)
	}

}

// Insert node to Head Node
func (l *SinglyLinkedList) InsertToHeadNode(nodeData int) {

	//if list is empty
	if l.Head == nil {
		l.Head = CreateNode(nodeData)
		return
	}

	tempNode := CreateNode(nodeData)

	tempNode.Next = l.Head
	l.Head = tempNode

}

// Insert node to Last Node
func (l *SinglyLinkedList) InsertToLastNode(nodeData int) {

	// if list is empty
	if l.Head == nil {
		l.Head = CreateNode(nodeData)
		return
	}

	tempNodeForIteration := l.Head
	tempNode := CreateNode(nodeData)

	for tempNodeForIteration.Next != nil {
		tempNodeForIteration = tempNodeForIteration.Next
	}

	tempNodeForIteration.Next = tempNode
}

// InsertNodeToAfterTargetNode Assume that LinkedList has unique values.
func (l *SinglyLinkedList) InserToAfterTargetNode(specificTargetData int, nodeData int) {

	//If List is empty
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}

	tempNodeForIteration := l.Head
	tempNode := CreateNode(nodeData)

	//Change Middle positions
	for tempNodeForIteration.Next != nil {

		//Find data in the list for insert after
		if tempNodeForIteration.Data == specificTargetData {
			tempNode.Next = tempNodeForIteration.Next
			tempNodeForIteration.Next = tempNode
			return
		}

		tempNodeForIteration = tempNodeForIteration.Next
	}

	// Change Last
	if tempNodeForIteration.Data == specificTargetData {
		l.InsertToLastNode(nodeData)
		return
	}

	fmt.Printf("There is no target data: %d in this linkedList\n", specificTargetData)
}

// Delete Head Node
func (l *SinglyLinkedList) DeleteFromHeadNode() {

	//if list is empty
	if l.Head == nil {
		fmt.Printf("Linked List is empty")
		return
	}

	//Only Head node exist
	if l.Head.Next == nil {
		l.Head = nil
		return
	}

	l.Head = l.Head.Next
}

// Delete  Last Node
func (l *SinglyLinkedList) DeleteFromLastNode() {

	//If only 1 element is exist in list
	if l.Head.Next == nil {
		l.DeleteFromHeadNode()
		return
	}

	tempNodeForIteration := l.Head

	for tempNodeForIteration.Next.Next != nil {
		tempNodeForIteration = tempNodeForIteration.Next
	}
	//Cut The Connection Link whatever you call :)
	tempNodeForIteration.Next = nil
}

// Delete Target Node if exist :)
func (l *SinglyLinkedList) DeleteFromTargetNode(targetData int) {

	var (
		flag bool
	)

	// Delete From Head
	if l.Head.Data == targetData {
		l.DeleteFromHeadNode()
		return
	}

	tempNodeForIteration := l.Head

	for tempNodeForIteration.Next != nil {

		//Found Case
		if tempNodeForIteration.Next.Data == targetData {
			flag = true
			break
		}

		tempNodeForIteration = tempNodeForIteration.Next
	}

	if flag {
		tempNodeForIteration.Next = tempNodeForIteration.Next.Next
	} else {
		fmt.Println("Target node does not exist in this list.")
	}

}
