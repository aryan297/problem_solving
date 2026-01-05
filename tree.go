

package main

import "fmt"


type Node struct{
	Val string 
	Left *Node
	Right *Node
}


func Inorder(root *Node) {
	if root==nil{
		return
	}
	Inorder(root.Left)
	fmt.Println(root.Val)
	Inorder(root.Right)
}


func LevelOrder(root *Node) [][]string{

	if root == nil{
		return [][]string{}
	}

	var result [][]string
	queue := []*Node{root} 



}

func main(){
   root:=&Node{Val:"1"}
   root.Left= &Node{Val:"2"}
   root.Right=&Node{Val:"3"}
   fmt.Println(root.Val)
    fmt.Println(root.Left.Val)
    fmt.Println(root.Right.Val)

   Inorder(root)

}