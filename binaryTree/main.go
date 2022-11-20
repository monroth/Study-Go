package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Walk(root *TreeNode, ch chan *TreeNode) {
	defer close(ch)
	var recurseStep func(*TreeNode)
	recurseStep = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			recurseStep(root.Left)
			ch <- root
			recurseStep(root.Right)
		}
	}
	recurseStep(root)
}

func inorderTraversal(root *TreeNode) []int {
	walkedPath := []int{}
	ch := make(chan *TreeNode)
	go Walk(root, ch)
	for node := range ch {
		walkedPath = append(walkedPath, node.Val)
	}
	return walkedPath
}

func equalNode(nodeFirst, nodeSecond *TreeNode) (bool, bool) {
	if nodeFirst == nil && nodeSecond != nil {
		return false, false
	}
	if nodeFirst != nil && nodeSecond == nil {
		return false, false
	}
	if nodeFirst == nil && nodeSecond == nil {
		return true, false
	}
	if nodeFirst.Val != nodeSecond.Val {
		return false, true
	}
	return true, true
}

func Same(rootFirst, rootSecond *TreeNode) bool {
	var (
		recurseStep   func(*TreeNode, *TreeNode)
		wg            sync.WaitGroup
		done          bool
		bufferedTrees = make(map[struct{ first, second *TreeNode }]struct{})
		exists        = struct{}{}
		mu            sync.Mutex
	)
	recurseStep = func(root1 *TreeNode, root2 *TreeNode) {
		defer wg.Done()
		if _, ok := bufferedTrees[struct{ first, second *TreeNode }{root1, root2}]; ok {
			return
		}
		if !done {
			if eql, notNil := equalNode(root1, root2); !eql {
				mu.Lock()
				done = true
				mu.Unlock()
			} else if notNil {
				wg.Add(2)
				go recurseStep(root1.Left, root2.Left)
				go recurseStep(root1.Right, root2.Right)
			}
		}
		mu.Lock()
		bufferedTrees[struct{ first, second *TreeNode }{root1, root2}] = exists
		mu.Unlock()
	}
	wg.Add(1)
	go recurseStep(rootFirst, rootSecond)

	wg.Wait()
	return !done
}

func DeepCopy(tree *TreeNode) *TreeNode {
	var (
		left  *TreeNode = nil
		right *TreeNode = nil
	)
	if tree.Left != nil {
		left = DeepCopy(tree.Left)
	}
	if tree.Right != nil {
		right = DeepCopy(tree.Right)
	}
	return &TreeNode{tree.Val, left, right}
}

func main() {
	goodTree := TreeNode{0, nil, nil}

	badTree := goodTree
	badTree.Left = &badTree
	badTree.Right = &badTree

	halfBadTree := badTree
	halfBadTree.Right = nil

	quaterBadTree := halfBadTree
	quaterBadTree.Left = &quaterBadTree

	firstTree := goodTree
	firstTree.Left = &goodTree

	secondTree := *DeepCopy(&firstTree)

	thirdTree := secondTree
	thirdTree.Right = &secondTree

	forthTree := *DeepCopy(&thirdTree)

	forthTree.Right.Left.Val = 2

	fmt.Println(Same(thirdTree.Left, forthTree.Left))
}
