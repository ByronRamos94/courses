package main

import (
	"fmt"
	"sync"
	"time"
)

func say(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(word)
}
func main() {
	/*The waitgropus are more efficient than using an sleep command*/
	var wg sync.WaitGroup
	fmt.Println("HOLA")
	wg.Add(1)
	go say("MUNDO", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("ADIOS")
	time.Sleep(time.Second * 1)
}
