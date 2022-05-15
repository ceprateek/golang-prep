package divide_conquer

import "fmt"

func PlayRoman(){
	s := intToRoman(3)
	fmt.Println(s)
}

func intToRoman(num int) string {
	numerals := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	mapOfSymbols := getMapOfRomanPossibilities()
	currentRoman := ""
	k := 0
	for num>0 {
		if num >= numerals[k]{
			numOfSymbols := num/numerals[k]
			num = num%numerals[k]
			currentRoman = generateRoman(mapOfSymbols[numerals[k]],currentRoman,numOfSymbols)
		}
		k++
	}
	return currentRoman
}

func generateRoman(symbol,currentRoman string, no int) string{
	for i:=0;i<no;i++{
		currentRoman += symbol
	}
	return currentRoman
}

func getMapOfRomanPossibilities() map[int]string{
	mapOfSymbols := make(map[int]string, 13)
	mapOfSymbols[1000] = "M"
	mapOfSymbols[900]  = "CM"
	mapOfSymbols[500]  = "D"
	mapOfSymbols[400]  = "CD"
	mapOfSymbols[100]  = "C"
	mapOfSymbols[90]  = "XC"
	mapOfSymbols[50]   ="L"
	mapOfSymbols[40]   ="XL"
	mapOfSymbols[10]   ="X"
	mapOfSymbols[9]    = "IX"
	mapOfSymbols[5]    = "V"
	mapOfSymbols[4]    = "IV"
	mapOfSymbols[1]    = "I"
	return mapOfSymbols

}
