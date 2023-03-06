package main

import "golang-problems/data-structures/linked_list"

func main() {

	//create singlyLinkedList
	linkedList := linked_list.CreateSinglyLinkedList(10)
	linkedList.InsertToHeadNode(20)
	linkedList.PrintLinkedListToScreen()

	linkedList.InsertToHeadNode(30)
	linkedList.PrintLinkedListToScreen()

}
