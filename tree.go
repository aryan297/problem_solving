package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


func CreateNode() *TreeNode{
	root := &TreeNode{Val:1}
	root.Left=&TreeNode{Val:1}
	root.Right=&TreeNode{Val:2}
	root.Left.Left=&TreeNode{Val:3}
	root.Left.Right=&TreeNode{Val:4}

	return root 
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

func maxHeight(root *TreeNode) int{

	if root ==nil{
		return -1
	}

	left :=maxHeight(root.Left)
	right := maxHeight(root.Right)

	if left > right{
		return left+1
	}

	return right+1

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

	if tree==nil{
		return 
	}

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



func sumOfNode(root *TreeNode) int{

	if root == nil{
		return 0
	}


	left :=sumOfNode(root.Left)
	right:=sumOfNode(root.Right)

	return left+right+root.Val
}


func sumOfLeaf(root *TreeNode) int{

	if root==nil{
		return 0
	}

	if root.Left==nil && root.Right==nil{
		return root.Val
	}

	return sumOfLeaf(root.Left) + sumOfLeaf(root.Right)
}




func isCompleteBinaryTree(root *TreeNode) bool{

	if root ==nil{
		return true
	}

	queue :=[]*TreeNode{root}
	seeNull:=false

	for len(queue)>0{
		node :=queue[0]
		queue =queue[1:]


		if node == nil{
			seeNull=true
		} else{

			if seeNull{
				return false
			}

			queue=append(queue,node.Left)
			queue=append(queue,node.Right)
		}

	}


return true

}


func maxSum(root *TreeNode) int{
	maxSum := -1<<31

	var sum func(*TreeNode) int

	sum = func(node *TreeNode) int{


		if node==nil{
			return 0
		}

		left :=max(0,sum(node.Left))
		right := max(0,sum(node.Right))

		maxSum=max(maxSum , node.Val+left+right)

	    return node.Val + max(left, right)

	}

	sum(root)

	return maxSum
}




func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func invertTree( root *TreeNode)*TreeNode{
	if root ==nil{
		return nil
	}

	root.Left , root.Right = root.Right, root.Left
   
    invertTree(root.Left)
    invertTree(root.Right)
	return root

}


func longestUnivaluePath(root *TreeNode) int{
	longest :=0
	var dfs func(*TreeNode) int

	dfs=func(node *TreeNode) int{
		if node ==nil{
			return 0
		}

		leftPath :=dfs(node.Left)
		rightPath :=dfs(node.Right)

		left,right:=0,0

		if node.Left!=nil && node.Val==node.Left.Val{
			left=leftPath+1
		}

		if node.Right!=nil && node.Val==node.Right.Val{
			right=rightPath+1
		}

      longest = max(longest , left+right)

     return max(left, right)

	}

	dfs(root)

	return longest;
}


func main() {
 tree:= CreateNode()

 printAllNode(tree)

 value := CountNodes(tree)

 depth := maxDepth(tree)

 minDepths := minDepth(tree)

 leaf :=countLeaves(tree)

 sum := sumOfNode(tree)

 univalue := longestUnivaluePath(tree)

 isCompleteBinaryTrees := isCompleteBinaryTree(tree)

 invertTrees := invertTree(tree)


 sumLeaf := sumOfLeaf(tree)

 height :=maxHeight(tree)
 fmt.Println(height, "max height")

 fmt.Println(depth, "max depth")

 fmt.Println(value ,"count")

 fmt.Println(leaf,"leaf")
 	fmt.Println("Diameter:", MaxDiameter(tree))

 fmt.Println(minDepths , "min depth")

 fmt.Println(sum , "sum of node")
  fmt.Println(sumLeaf , "sum of leaf")
  fmt.Println(isCompleteBinaryTrees , " is complete binary tree")

  fmt.Println(univalue ,"longest univalue sum")
  printAllNode(invertTrees)

 levelOrderTraversal(tree)


}