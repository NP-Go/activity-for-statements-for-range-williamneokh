package main

import "fmt"

func main() {
	//Insert code for only even here
	var input1 int
	var input2 int

	var evenNum []int
	var oddNum []int

	fmt.Println("Please key in first integer number")
	_, _ = fmt.Scanln(&input1)
	fmt.Println("Please key in second integer number")
	_, _ = fmt.Scanln(&input2)

	for input1 <= input2 {
		if input1%2 == 0 {
			evenNum = append(evenNum, input1)
		} else if input1%2 == 1 {
			oddNum = append(oddNum, input1)

		}
		input1 = input1 + 1
	}
	numPrint(oddNum)
	numPrint(evenNum)

}

func numPrint(num []int) {
	for _, number := range num {
		fmt.Println(number)
	}
	fmt.Println("=============")
}
