package stacks

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	outMap := make(map[int]int)
	out := make([]int, 0, len(nums1))

	stack := Stack{}
	for _, num := range nums2 {
		for !stack.IsEmpty() && num > stack.Peek().(int) {
			val := stack.Pop()
			outMap[val.(int)] = num
		}
		stack.Push(num)
	}

	for _, num := range nums1 {
		if val, ok := outMap[num]; ok {
			out = append(out, val)
		} else {
			out = append(out, -1)
		}
	}
	return out
}

func PlayNextGreaterElement() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
