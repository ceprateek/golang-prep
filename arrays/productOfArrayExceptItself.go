package arrays

import "fmt"

func PlayProductArratExceptItself()  {
	findProductArrayExceptItself([]int{1, 2, 3, 4, 5})
}

func findProductArrayExceptItself(in []int){
	product := 1
	for _, val := range in{
		product *= val
	}
	for _,val := range in{
		fmt.Println(product/val)
	}
}
