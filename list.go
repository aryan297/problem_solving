package main


import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}


func reverse(head *ListNode) *ListNode{

	var prev *ListNode

	cur:=head

	for cur!=nil{

		next := cur.Next
		cur.Next=prev
		prev=cur
		cur=next

	}

return prev
}

func printNode(head *ListNode){
	for head!=nil{
		fmt.Println(head.Val,"value")
		head=head.Next
	}
}


func main(){

head := &ListNode{Val:2}
head.Next = &ListNode{Val:4}
head.Next.Next = &ListNode{Val:3}


printNode(head)

r := reverse(head)

printNode(r)


}