package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// It returns the number of bytes written and any write error encountered.

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
	fmt.Fprintf(w, "Hello, %s", "world2")
	// Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
