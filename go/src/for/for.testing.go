package main

import "fmt"

func main() {
	num := 2
	for i := 1; i < 10; i++ {
		result := num * i
		output := fmt.Sprintf("%d x %d = %d", num, i, result)
		fmt.Println(output)
	}

	for j := 0; j <= 100; j++ {
		if j%2 != 0 || j == 0 {
			continue
		}
		fmt.Println(j)
	}
}
