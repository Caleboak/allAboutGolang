package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	minReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Pls, input the minimum number: \n") //recieve user input
	minInput, _ := minReader.ReadString('\n')
	minimum := strings.Replace(minInput, "\n", "", -1)
	min, err := strconv.Atoi(minimum) //convert to int from str
	if err != nil {                   //if error then it is not a number
		fmt.Println("Not a number")
		return //exit
	}
	maxReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Pls, input the maximum number: \n")
	maxInput, _ := maxReader.ReadString('\n')
	maximum := strings.Replace(maxInput, "\n", "", -1)
	max, err1 := strconv.Atoi(maximum)
	if err1 != nil {
		fmt.Println("Not a number")
		return
	}
	if max < min { //if max is lesser than min
		fmt.Println("max can't be lesser than min")
		return //exit
	}
	sum := 0
	//THis for loop prints all the summation and the result
	for i := min; i <= max; i++ {
		if i%2 == 1 { //if it's odd, continue
			continue
		} else if i == max { //if its the max, don't add the '+' sign
			fmt.Printf("%d", i)
			sum += i
		} else if i%2 == 0 && i+1 == max { //if the cur is even and last is odd, don't add the '+' sign
			fmt.Printf("%d", i)
			sum += i
		} else {
			fmt.Printf("%d+", i)
			sum += i
		}
	}
	fmt.Printf("=%d\n", sum)
}
