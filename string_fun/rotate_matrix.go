package string_fun

import "fmt"

func PlayRotateMatrix(){
	in := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	for _, row := range in{
		fmt.Println(row)
	}

	rotateMatrix(&in)
	for _, row := range in{
		fmt.Println(row)
	}
}

func rotateMatrix(in *[][]int){
	transposeMatrix(in)
	for _, row := range *in{
		fmt.Println(row)
	}
	input := *in
	for i:=0;i<len(input);i++{
		first := 0
		last := len(input[i])-1
		for first<last{
			input[i][first], input[i][last] = input[i][last], input[i][first]
			last--
			first++
		}
	}
}

func transposeMatrix(in *[][]int){
	input := *in
	for i:=0;i<len(input);i++{
		for j:=i;j<len(input);j++{
			input[i][j], input[j][i] = input[j][i], input[i][j]
		}
	}
}
