package main

import "fmt"


type Tree struct{
	Val int
	Left *Tree
	Right *Tree
}


func CreateTree() *Tree{
	tree := &Tree{Val:1}

	tree.Left = &Tree{Val:3}
	tree.Right = &Tree{Val:5}

	return tree
}


func PrintLeave(root *Tree){
	if root==nil{
		return
	}

	if root.Left==nil || root.Right==nil{
		fmt.Println(root.Val,"")
		return
	}
		fmt.Println(root.Val)
	PrintLeave(root.Left)
	PrintLeave(root.Right)

}

func main(){
	tree:=CreateTree()
	PrintLeave(tree)

}