package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received: \n")
	io.WriteString(w, "This is my first web server. \n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello request received: \n")
	io.WriteString(w, "Hello, World! \n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed! \n")
	} else if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Server started! \n")
}
