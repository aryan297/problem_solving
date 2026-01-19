package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


func CreateNode() *TreeNode{
	tree := &TreeNode{Val:1}

	tree.Left=&TreeNode{Val:2}
	tree.Right=&TreeNode{Val:3}
	tree.Left.Right = &TreeNode{Val:5}
	tree.Left.Left = &TreeNode{Val:5}

	return tree
}


func printAllNode(root *TreeNode){

	if root==nil{
		return
	}


	//if root.Left==nil && root.Right==nil{
		//fmt.Println(root.Val)  // if need to print leaf node this logic will work
	//}
	fmt.Println(root.Val)

	printAllNode(root.Left)
	printAllNode(root.Right)

}


func CountNodes(root *TreeNode) int{
	if root==nil{
		return 0
	}

	return 1+CountNodes(root.Left)+CountNodes(root.Right)
}


func maxDepth(root *TreeNode) int{

	if root==nil{
		return 0
	}

	leftDepth:=maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth >rightDepth {
		return leftDepth +1
	}

	return rightDepth +1

}

func minDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}

	if root.Left==nil{
		return 1+minDepth(root.Right)
	}

    if root.Right==nil{

    	return 1+minDepth(root.Left)

    }

    leftDepth:=maxDepth(root.Left)
     rightDepth:=maxDepth(root.Right)

     if leftDepth >rightDepth{
     	return rightDepth+1
     
      }

      return leftDepth+1
}

func countLeaves(root  *TreeNode) int{ 

	if root == nil{
		return 0
	}

	if root.Left== nil && root.Right ==nil{
		return 1
	}

	return countLeaves(root.Left) +countLeaves(root.Right)

}


func MaxDiameter(tree *TreeNode) int{
	maxDiameter :=0

	var height func(*TreeNode) int

	height= func(node *TreeNode) int{

		if node == nil{
			return 0
		}
		left :=height(node.Left)
		right :=height(node.Right)

		if left+right > maxDiameter{
			maxDiameter=left+right
		}

		if left > right{
			return left+1
		}
		return right+1
	}
	height(tree)
  return maxDiameter

}

func levelOrderTraversal(tree *TreeNode) {

	queue :=[]*TreeNode{tree}

	res :=[][]int{}


	for len(queue)>0{

		size:=len(queue)

		level:=[]int{}


		for i:=0;i<size;i++{

			node :=queue[0]
			queue = queue[1:]

			level = append(level,node.Val)

			if node.Left!=nil{
				queue=append(queue,node.Left)
			}

			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
		}

		res =append(res,level)



	}

	fmt.Println(res,"result")


}




func main() {
 tree:= CreateNode()

 printAllNode(tree)

 value := CountNodes(tree)

 depth := maxDepth(tree)

 minDepths := minDepth(tree)

 leaf :=countLeaves(tree)

 fmt.Println(depth, "max depth")

 fmt.Println(value ,"count")

 fmt.Println(leaf,"leaf")
 	fmt.Println("Diameter:", MaxDiameter(tree))

 fmt.Println(minDepths , "min depth")

 levelOrderTraversal(tree)


}