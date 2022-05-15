package recursion

import "fmt"

func PlaySubsetSumFinder(){
	a := []int{3, 34, 4, 12, 5, 2}
	//s := 30
	fmt.Println(findSubsets(a))
}

func isSubsetWithSum(in []int, sum int) bool{
	temp := make([]int, 0)
	tempSum := 0
	return isSubsetHelper(&in, sum, &temp, &tempSum)
}

func isSubsetHelper(in *[]int, sum int, temp *[]int, tempsum *int) bool {
	if len(*in) == 0{
		if sum == *tempsum{
			fmt.Println(*temp)
			return true
		}
	}else {
		choose := (*in)[0]
		*in = (*in)[1:]
		result := isSubsetHelper(in, sum, temp, tempsum)
		if result{
			return result
		}
		*temp = append(*temp, choose)
		*tempsum += choose
		result = isSubsetHelper(in, sum, temp, tempsum)
		if result{
			return result
		}
		k := []int{choose}
		*in = append(k, *in...)
	}
	return false
}

func findSubsets(in []int) [][]int{
	var result [][]int
	t := make([]int, 0, len(in))
	findSubsetHelper(&in, &result, &t)
	return result
}

func findSubsetHelper(in *[]int, result *[][]int, temp *[]int){
	if len(*in) == 0 {
		a := make([]int, len(*temp))
		copy(a, *temp)
		*result = append(*result, a)
		t := make([]int, 0, len(*in))
		temp = &t
		return
	}
	for i := 0; i < len(*in); i++ {
		choose := (*in)[0]
		*in = (*in)[1:]
		fmt.Println(*in)
		findSubsetHelper(in, result, temp)
		*temp = append(*temp, choose)
		findSubsetHelper(in, result, temp)
		k := []int{choose}
		*in = append(k, *in...)
		*temp = (*temp)[:len(*temp)-1]
	}
}
