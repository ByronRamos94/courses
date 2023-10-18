package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
func main() {
	response, err := http.Get("http://google.com")

	e := webWriter{}
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(e, response.Body)

}
