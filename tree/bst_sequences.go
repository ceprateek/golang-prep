package tree

import (
	"fmt"
	"reflect"
)

type tree struct {
	root *treeNode
}

func PlayTree() {
	t := NewTree()
	t.Insert(10)
	t.Insert(5)
	t.Insert(20)
	t.Insert(2)
	t.Insert(6)
	t.Insert(30)
	t.Insert(11)
	t.Dump()
	result1 := t.GetSequences()
	fmt.Println(result1)
}

func (t *tree) GetSequences() [][]int {
	sq := make([]int, 0)
	result1 := getSq(t.root, &sq)
	return result1
}

func getSq(node *treeNode, sq *[]int) (result [][]int) {
	if node == nil {
		r := make([]int, len(*sq))
		copy(r, *sq)
		result = append(result, r)
		return result
	}
	if node.left == nil && node.right == nil{
		r := []int{node.val}
		return [][]int{r}
	}
	leftSubtee := getSq(node.left, sq)
	right := getSq(node.right, sq)
	for _, treeInSub := range leftSubtee {
		for _, rightSub := range right {
			weaved := getWeaved(treeInSub, rightSub)
			for _, w := range weaved {
				result = append(result, append([]int{node.val}, w...))
			}
		}
	}

	return result
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func NewTree() *tree {
	t := tree{root: nil}
	return &t
}

func (t *tree) Insert(val int) {
	if t.root == nil {
		root := treeNode{val: val}
		t.root = &root
		return
	}
	insert(t.root, val)
}

func insert(t *treeNode, val int) {
	nodeVal := t.val
	if val < nodeVal {
		//insert to left
		if t.left == nil {
			left := treeNode{
				val: val,
			}
			t.left = &left
		} else {
			insert(t.left, val)
		}
	} else {
		//insert to right
		if t.right == nil {
			right := treeNode{val: val}
			t.right = &right
		} else {
			insert(t.right, val)
		}
	}
}

func (t *tree) Dump() {
	queue := make([]*treeNode, 0)
	queue = append(queue, t.root)
	result := make([]int, 0)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		result = append(result, first.val)
		if first.left != nil {
			queue = append(queue, first.left)
		}
		if first.right != nil {
			queue = append(queue, first.right)
		}
	}
	fmt.Println(result)
}

func getWeaved(l1, l2 []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, len(l1)+len(l2))
	for i := 0; i < len(l2); i++ {
		temp[len(l1)+i] = 1
	}
	var temp1 []int
	permuteHelper(&temp, &result, &temp1)

	for p:=0;p<len(result);p++ {
		var f, s int
		for i:=0;i<len(result[p]);i++ {

			if result[p][i] == 0 {
				result[p][i] = l1[f]
				f++
			} else {
				result[p][i] = l2[s]
				s++
			}
		}
	}
	return result

}

func permuteHelper(list *[]int, result *[][]int, temp *[]int) {
	if len(*list) == 0 {
		for _, permute := range *result {
			if reflect.DeepEqual(permute, *temp) {
				return
			}
		}
		t := make([]int, len(*temp))
		copy(t, *temp)
		*result = append(*result, t)
		return
	}
	for i := 0; i < len(*list); i++ {
		t := (*list)[i]
		*list = append((*list)[:i], (*list)[i+1:]...)
		*temp = append(*temp, t)
		permuteHelper(list, result, temp)
		*list = append((*list)[:i+1], (*list)[i:]...)
		(*list)[i] = t
		*temp = (*temp)[:len(*temp)-1]
	}
}
