package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurseStep(root *TreeNode, walkedPathPointer *[]int) {
	if root == nil {
		return
	} else {
		recurseStep(root.Left, walkedPathPointer)
		*walkedPathPointer = append(*walkedPathPointer, root.Val)
		recurseStep(root.Right, walkedPathPointer)
	}
}

func inorderTraversal(root *TreeNode) []int {
	walkedPath := []int{}
	recurseStep(root, &walkedPath)
	return walkedPath
}

func main() {
	fmt.Println(inorderTraversal(nil))
}
