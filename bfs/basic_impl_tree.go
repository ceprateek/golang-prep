package bfs

import "fmt"

// A Tree is a binary tree with integer values.
type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &TreeNode{
			Left:  nil,
			Value: val,
			Right: nil,
		}
	} else {
		t.root.Insert(val)
	}
}

func (n *TreeNode) Insert(v int) {
	if n == nil {
		return
	} else {
		if v < n.Value {
			if n.Left == nil {
				n.Left = &TreeNode{
					Left:  nil,
					Value: v,
					Right: nil,
				}
			} else {
				n.Left.Insert(v)
			}
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{
					Left:  nil,
					Value: v,
					Right: nil,
				}
			} else {
				n.Right.Insert(v)
			}
		}
	}
}

func (t *Tree) PrintLevelOrder(result chan int) {
	heightOfTree := t.root.Height()
	fmt.Println(fmt.Sprintf("height: %d", heightOfTree))
	for i := 1; i <= heightOfTree; i++ {
		t.root.PrintLevel(i, result)
	}
	close(result)
}

func (root *TreeNode) PrintLevel(level int, result chan int) {
	if root == nil {
		return
	}
	if level == 1 {
		result <- root.Value
	} else if level > 1 {
		root.Left.PrintLevel(level-1, result)
		root.Right.PrintLevel(level-1, result)
	}

}

func (root *TreeNode) Height() int {
	if root == nil {
		return 0
	}
	leftHeight := root.Left.Height()
	rightHeight := root.Right.Height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
