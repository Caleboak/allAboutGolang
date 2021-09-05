package main

import "fmt"

func main() {
	wkOneSpend := [][]int{
		{150, 350, 657, 123},
		{984},
		{456, 781, 546},
		{451, 50, 750},
		{540, 548, 90},
	}

	wkTwoSpend := [][]int{
		{450, 350, 657},
		{984, 564, 779},
		{426, 741, 544},
		{451, 500, 790},
		{500, 578},
	}
	wkOneSum, wkTwoSum := calculateSpend(wkOneSpend, wkTwoSpend)
	fmt.Printf("The total spend for week one is %d\n", wkOneSum)
	fmt.Printf("The total spend for week two is %d\n", wkTwoSum)
}
func calculateSpend(wkOneSpend [][]int, wkTwoSpend [][]int) (int, int) {
	wkOneSum := 0
	wkTwoSum := 0
	for i := 0; i < len(wkOneSpend); i++ {
		for j := 0; j < len(wkOneSpend[i]); j++ {
			wkOneSum += wkOneSpend[i][j]

		}
	}

	for i := 0; i < len(wkTwoSpend); i++ {
		for j := 0; j < len(wkTwoSpend[i]); j++ {
			wkTwoSum += wkTwoSpend[i][j]

		}
	}

	return wkOneSum, wkTwoSum
}
