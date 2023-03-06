package main

import "golang-problems/data-structures/linked_list"

func main() {

	//create singlyLinkedList
	linkedList := linked_list.CreateSinglyLinkedList(10)

	/*Add Operations

	//InsertToLastNode
	linkedList.InsertToLastNode(20)
	linkedList.InsertToLastNode(30)
	linkedList.InsertToLastNode(40)
	linkedList.PrintLinkedListToScreen()

	//Add After Head
	linkedList.InsertToAfterTargetNode(100, 100)
	linkedList.PrintLinkedListToScreen()

	*/

	linkedList.InsertToLastNode(20)
	linkedList.InsertToLastNode(30)
	linkedList.InsertToLastNode(40)

	linkedList.PrintLinkedListToScreen()
	linkedList.DeleteTargetNode(400)
	linkedList.PrintLinkedListToScreen()

}
