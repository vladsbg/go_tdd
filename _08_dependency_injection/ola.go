package main

import (
	"fmt"
	"io"
	"net/http"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Olá, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "mundo")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
