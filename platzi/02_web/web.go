package main

import (
	"fmt"
	"io"
	"net/http"
)

type writerWeb struct {
}

func (writerWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	w := writerWeb{}
	io.Copy(w, response.Body)
}
