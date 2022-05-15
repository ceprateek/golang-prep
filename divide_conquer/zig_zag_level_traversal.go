package divide_conquer

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := FindHeight(root.Left)
	right := FindHeight(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	resultToReturn := make([][]int, 0)
	if root == nil {
		return resultToReturn
	}
	height := FindHeight(root)
	for i := 1; i <= height; i++ {
		result := make([]int, 0)
		printLevel(root, i, &result, i)
		resultToReturn = append(resultToReturn, result)
	}
	return resultToReturn
}

func printLevel(root *TreeNode, level int, result *[]int, n int) {
	if root == nil{
		return
	}
	if level == 1 {
		*result = append(*result, root.Val)
		return
	}
	if n%2 == 1 {
		printLevel(root.Left, level-1, result, n)
		printLevel(root.Right, level-1, result, n)
	} else {
		printLevel(root.Right, level-1, result, n)
		printLevel(root.Left, level-1, result, n)
	}
}

func PlayZigzag() {
	root := TreeNode{
		Val: 3,
	}
	l1 := TreeNode{Val: 9}
	r1 := TreeNode{Val: 20}
	l2 := TreeNode{Val: 15}
	r2 := TreeNode{Val: 7}
	root.Left = &l1
	root.Right = &r1
	r1.Left = &l2
	r1.Right = &r2
	r := zigzagLevelOrder(&root)
	fmt.Println(r)
}
