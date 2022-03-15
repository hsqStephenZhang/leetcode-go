package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	curPath := []int{}
	backtrace(root, targetSum, 0, &curPath, &res)
	return res
}

func backtrace(root *TreeNode, targetSum int, curSum int, curPath *[]int, res *[][]int) {
	if root == nil {
		return
	}
	curSum += root.Val
	*curPath = append(*curPath, root.Val)
	if root.Left == nil && root.Right == nil && curSum == targetSum {
		newPath := make([]int, len(*curPath))
		copy(newPath, *curPath)
		*res = append(*res, newPath)
	}
	if root.Left != nil {
		backtrace(root.Left, targetSum, curSum, curPath, res)
	}
	if root.Right != nil {
		backtrace(root.Right, targetSum, curSum, curPath, res)
	}
	*curPath = (*curPath)[:len(*curPath)-1]
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2}},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1}},
		},
	}
	fmt.Println(pathSum(root, 22))
}
