package main

import "fmt"

type LinkedList struct{

	Val int
	Next *LinkedList
}


func reverseList(h *LinkedList) *LinkedList{

	var prev *LinkedList
	cur:=h

	for cur!=nil{
		next:=cur.Next
		cur.Next=prev
		prev=cur
		cur=next

	}

	return prev
}



func Printlist(h *LinkedList){

	for h!=nil{
		fmt.Print(h.Val," ")
		h=h.Next
	}

	fmt.Println()

}





func main(){

	node := &LinkedList{0,nil}
	node.Next= &LinkedList{1 , nil}
	node.Next.Next= &LinkedList{2 , nil}
	node.Next.Next.Next= &LinkedList{3 , nil}

	Printlist(node)

	p := reverseList(node)
		Printlist(p)
}