package main

import (
	"fmt"
	Logger "gotest/logger"
	"net/http"
)

func main() {
	Logger.Log("Program started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
