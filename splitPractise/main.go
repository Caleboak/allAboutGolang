package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")
	name, _ := reader.ReadString('\n')

	newName := strings.Split(name, " ")

	first := strings.ToUpper(string(newName[0][0]))
	second := strings.ToUpper(string(newName[1][0]))

	fmt.Printf("My first Name is %s \n", first+newName[0][1:])
	fmt.Printf("My last Name is %s", second+newName[1][1:])

}
