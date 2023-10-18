package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go validateServer(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	timeTranscurrent := time.Since(start)
	fmt.Printf("time elapsed: %s\n", timeTranscurrent)
}

func validateServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		channel <- server + " is available"
	} else {
		channel <- server + " not available"
	}
}
