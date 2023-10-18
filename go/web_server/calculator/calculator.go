package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(input string, operator string) int {
	values := strings.Split(input, operator)
	firstOperator := parse(values[0])
	secondOperator := parse(values[1])
	var result int
	switch operator {
	case "+":
		result = firstOperator + secondOperator
		fmt.Println(result)
		return result
	case "-":
		result = firstOperator - secondOperator
		fmt.Println(result)
		return result
	case "*":
		result = firstOperator * secondOperator
		fmt.Println(result)
		return result
	case "/":
		result = firstOperator * secondOperator
		fmt.Println(result)
		return result
	default:
		fmt.Println("please type a valid operator")
		return 0

	}
}
func parse(input string) int {
	result, e := strconv.Atoi(input)
	if e != nil {
		fmt.Println(e)
	}
	return result
}

func writeInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}
func main() {
	operation := writeInput()
	operator := writeInput()
	c := calc{}
	c.operate(operation, operator)
}
