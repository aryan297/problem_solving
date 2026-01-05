package main

import "fmt"



type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}


func ( l *LinkedList) InsertAtEnd(value int){
	  newNode := &Node{data : value}

	  if l.head ==nil{
	  	l.head=newNode
	  	return
	  }

	  current := l.head;

	  for current.next !=nil{
	  	current = current.next
	  }

	  current.next=newNode
}

func (l *LinkedList) DeleteNode(value int){
	if l.head==nil{
		return
	}

	if l.head.data==value{
		l.head=l.head.next
		return
	}

	current:=l.head
	if current.next!=nil && current.next.data!=value{
		current=current.next
	}

if current.next!=nil{
	current.next=current.next.next
}

}


func (l *LinkedList) InsertAtBegin(value int){
   newNode := &Node{data:value,next:l.head}

   l.head=newNode
}


func (l *LinkedList) print(){
	current:=l.head

	for current !=nil{
		fmt.Println(current.data)
		current=current.next
	}

	fmt.Println("nil")

}

func main(){

	list:=&LinkedList{}

	list.InsertAtEnd(30)
	list.InsertAtBegin(10)

    list.print()
	

}