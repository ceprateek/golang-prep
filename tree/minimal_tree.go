package tree

import "fmt"

type node struct {
	data        int
	left, right *node
}

type binaryTree struct {
	root *node
}

func (b *binaryTree) InorderTraversal() {
	inorderTraversal(b.root)
}

func (b *binaryTree) PreorderTraversal() {
	preOrderTraversal(b.root)
}

func (b binaryTree) PostOrderTraversal() {
	postOrderTraversal(b.root)
}

func preOrderTraversal(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}

func inorderTraversal(root *node) {
	if root == nil {
		return
	}
	inorderTraversal(root.left)
	fmt.Println(root.data)
	inorderTraversal(root.right)
}

func postOrderTraversal(root *node) {
	if root == nil {
		return
	}
	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
	fmt.Println(root.data)
}

func sortedArrayToBST(nums []int, start, end int) *node {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	nodeToInsert := node{data: nums[mid]}
	nodeToInsert.left = sortedArrayToBST(nums, start, mid-1)
	nodeToInsert.right = sortedArrayToBST(nums, mid+1, end)
	return &nodeToInsert
}
func (b *binaryTree) Height() int {
	return height(b.root, 0)
}

func height(root *node, h int) int {
	if root == nil {
		return h
	}
	heightLeft := height(root.left, h+1)
	heightRight := height(root.right, h+1)
	return max(heightRight, heightLeft)

}

func (b *binaryTree) leverOrderTraversalBfs()[][]int{
	result = make([][]int, 0)
	leverOrderBfs(b.root, 0)
	return result
}

var result [][]int

func leverOrderBfs(root *node, level int){
	if root == nil{
		return
	}
	var currentLevel []int
	if len(result) <= level{
		currentLevel = make([]int, 0)
		result = append(result, currentLevel)
	}
	currentLevel = result[level]
	currentLevel = append(currentLevel, root.data)
	result[level] = currentLevel
	leverOrderBfs(root.left, level+1)
	leverOrderBfs(root.right, level+1)
}

func (b *binaryTree) Bfs() {
	q := make([]*node, 0)
	q = append(q, b.root)
	for len(q) > 0 {
		currentLength := len(q)
		for currentLength > 0 {
			firstElement := q[0]
			fmt.Printf("%d   ", firstElement.data)
			q = q[1:]
			if firstElement.left != nil {
				q = append(q, firstElement.left)
			}
			if firstElement.right != nil {
				q = append(q, firstElement.right)
			}
			currentLength--
		}
		fmt.Println()
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PlayMinimalTree() {
	nums := []int{1, 2, 3, 4, 5, 6}
	b := binaryTree{}
	b.root = sortedArrayToBST(nums, 0, len(nums)-1)
	result := b.leverOrderTraversalBfs()
	fmt.Println(result)
}
