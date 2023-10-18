package main

import (
	"fmt"
	"workshop/golang/calculator"
)

func main() {
	params := calculator.Terms{
		FirstNum:  1,
		SecondNum: 2,
		CallBack: func() {
			fmt.Println("processing sum...")

		},
	}
	result := fmt.Sprintf("result: %v ", params.Sum())
	fmt.Println(result)

}
