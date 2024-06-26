package sliding_window

import "fmt"

/*
Day 4.4: Maximum Fruits into Baskets — Medium
*/

/*
You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.



Example 1:

Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.
Example 2:

Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].
Example 3:

Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].
Example 4:

Input: fruits = [3,3,3,1,2,1,1,2,3,3,4]
Output: 5
Explanation: We can pick from trees [1,2,1,1,2].
*/

func PlayMaxFruits(){
	fruits := []int{3,3,3,1,2,1,1,2,3,3,4}
	maxFruitsInBasket(fruits)
}

func maxFruitsInBasket(fruits []int) {
	var left, max_window_size int
	var pickedUpFruits []int
	pickedUpFruitsTypes := make(map[int]int)
	for right := 0; right < len(fruits); right++ {
		pickedUpFruits = append(pickedUpFruits, fruits[right])
		if _, ok := pickedUpFruitsTypes[fruits[right]]; !ok {
			pickedUpFruitsTypes[fruits[right]] = 0
		}
		if len(pickedUpFruitsTypes) > 2 {
			typeOfFruitsOnLeft := fruits[left]
			left = findLastIndexOfElementArr(fruits, typeOfFruitsOnLeft, left, right)+1
			delete(pickedUpFruitsTypes, typeOfFruitsOnLeft)
		}
		max_window_size = max_num(max_window_size, right-left+1)
	}
	fmt.Println(max_window_size)
}

func findLastIndexOfElementArr(in []int, element, i, j int) int{
	for k:= j; k>=i;k--{
		if in[k] == element{
			return k
		}
	}
	return -1
}
