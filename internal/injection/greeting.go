package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HandlerMyGreeting(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "word")
}

func main() {
	err := http.ListenAndServe(":5001", http.HandlerFunc(HandlerMyGreeting))

	if err != nil {
		fmt.Println(err)
	}
}
