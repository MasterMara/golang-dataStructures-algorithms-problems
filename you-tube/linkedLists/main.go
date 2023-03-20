package main

import (
	singly_linked_list "linkedLists/linkedLists"
)

// Entrypoint of application
func main() {

	//Create linked List
	myLinkedList := singly_linked_list.CreateSinglyLinkedList(10)

	/* Test cases for PrintSinglyLinkedListToScreen

	//List is not empty --> Test passed
	myLinkedList.PrintSinglyLinkedListToScreen()

	//List is empty case --> Test passed
	mySecondLinkedList := singly_linked_list.SinglyLinkedList{}
	mySecondLinkedList.PrintSinglyLinkedListToScreen()
	*/

	/*  Test cases for InserToHeadNode

	//If list is empty --> Test Passed
	mySecondLinkedList := singly_linked_list.SinglyLinkedList{}
	mySecondLinkedList.InsertToHeadNode(100)
	mySecondLinkedList.PrintSinglyLinkedListToScreen()

	// Change The Head Node --> Test Passed
	myLinkedList.InsertToHeadNode(20)
	myLinkedList.InsertToHeadNode(30)
	myLinkedList.PrintSinglyLinkedListToScreen()
	*/

	/*  Test cases for InserToLastNode

	//If list is empty --> Test Passed
	mySecondLinkedList := singly_linked_list.SinglyLinkedList{}
	mySecondLinkedList.InserToLastNode(100)
	mySecondLinkedList.PrintSinglyLinkedListToScreen()

	// Change The Last Node --> Test Passed
	myLinkedList.InsertToLastNode(20)
	myLinkedList.InsertToLastNode(30)
	myLinkedList.PrintSinglyLinkedListToScreen()

	*/

	/* Test cases for InserToAfterTargetNode

	//Found and insert after Head Node case:

	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.InserToAfterTargetNode(10, 50)
	//Expected 10,50,20,30,40
	myLinkedList.PrintSinglyLinkedListToScreen()

	//Found and insert after LastNode case:
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.InserToAfterTargetNode(40, 50)
	//Expected 10,50,20,30,40
	myLinkedList.PrintSinglyLinkedListToScreen()

	// Found and insert middle case:
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.InserToAfterTargetNode(30, 50)
	//Expected 10,20,30,50,40
	myLinkedList.PrintSinglyLinkedListToScreen()

	// NotFound case:
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.InserToAfterTargetNode(100, 50)
	//Expected There is no target data
	myLinkedList.PrintSinglyLinkedListToScreen()
	*/

	/* Test cases for deleteFromHeadNode:

	//Delete Singly Head Node

	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.deleteFromHeadNode()
	//AfterDelete
	myLinkedList.PrintSinglyLinkedListToScreen()

	//Delete Head Node From Normal List case:

	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.deleteFromHeadNode()
	//AfterDelete --> expected: no expected should be ,20,30,40 deleteFromHeadNode() dude :) be careful
	myLinkedList.PrintSinglyLinkedListToScreen()

	// //If Empty List case:
	mySecondLinkedList := singly_linked_list.SinglyLinkedList{}
	mySecondLinkedList.deleteFromHeadNode() //Expected is : Linked List is empty.
	mySecondLinkedList.PrintSinglyLinkedListToScreen() // we saw 2 empty cause deleteFromHeadNode is also print it so it's okey.
	*/

	/* Test cases for deleteFromLastNode

	//Delete From LatNode

	//Before Delete
	myLinkedList.PrintSinglyLinkedListToScreen()

	myLinkedList.DeleteFromLastNode()
	myLinkedList.DeleteFromLastNode()

	//After Delete
	myLinkedList.PrintSinglyLinkedListToScreen()

	//Before Delete
	myLinkedList.PrintSinglyLinkedListToScreen()

	//As Expected if 1 element in the list and delete it so list is empty:) Good Job.
	myLinkedList.DeleteFromLastNode()

	//After Delete
	myLinkedList.PrintSinglyLinkedListToScreen()
	*/

	/* Test cases for DeleteFromTargetNode:

	// Exist from Header
	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.DeleteFromTargetNode(10)
	//After deletion: Expected: 20,30,40 great result as you see:)
	myLinkedList.PrintSinglyLinkedListToScreen()

	// Found and DeleteFromLastNode
	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.DeleteFromTargetNode(40)
	//After deletion: Expected: 10,20,30 great result as you see:) Perfection:)
	myLinkedList.PrintSinglyLinkedListToScreen()

	// Only 1 node Exist and found
	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.DeleteFromTargetNode(10)
	//After deletion: Expected: should print  List is empty. ooo clap clap clap :)
	myLinkedList.PrintSinglyLinkedListToScreen()

	// NotFound CASE:
	//BeforeDelete
	myLinkedList.PrintSinglyLinkedListToScreen()
	myLinkedList.DeleteFromTargetNode(100)
	//After deletion: Expected: Target node does not exist in this list. Perfectoo:)

	*/

	/* Test Cases for PrintReverseSinglyLinkedListToScreen function:

	//Print Normally
	myLinkedList.PrintSinglyLinkedListToScreen()

	//Print Reverse Result is as expected :)
	singly_linked_list.PrintReverseSinglyLinkedListToScreen(myLinkedList.Head)

	*/

	//Create List 10 --> 20 --> 30 --> 40
	myLinkedList.InsertToLastNode(20)
	myLinkedList.InsertToLastNode(30)
	myLinkedList.InsertToLastNode(40)

	//Print Normally
	myLinkedList.PrintSinglyLinkedListToScreen()

	//Print Reverse Result is as expected :)
	singly_linked_list.PrintReverseSinglyLinkedListToScreen(myLinkedList.Head)

	// 180 + 205 = 385 line. That's clear. I wrote Some SinglyLinkedList functions to better understanding of linked list using golang.
	// If you enjoy the video plase subsrice the golang mastery channel and support me to next videos.
	// If you have any question just ask me in comments or send email to me
	// Todo for you: Refactor the existing whole codebase :) good lack :)
	// See you gophers :) lOVE GOLANG AND LOVE CODİNG

}
