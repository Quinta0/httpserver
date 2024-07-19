# Basic HTTP Server in Go
This document explains the functionalities of a basic HTTP server implemented in Go. The server responds to two different routes: `/` and `/hello`.

## Code Overview

### Package and Imports

```go
package main

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
)
```

The server starts by defining the main package and importing necessary libraries:

- errors: For handling errors.
- fmt: For formatted I/O operations.
- io: For basic input and output interfaces.
- net/http: For HTTP client and server implementations.
- os: For interacting with the operating system.

### Root Handler

```go
func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Request received: \n")
    io.WriteString(w, "This is my first web server. \n")
}
```
The getRoot function handles requests to the root URL (/). It prints a message to the console and responds to the client with a simple message.

### Hello Handler

```go
func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Hello request received: \n")
    io.WriteString(w, "Hello, World! \n")
}
```
The getHello function handles requests to the /hello URL. It prints a different message to the console and responds with a "Hello, World!" message.

### Main Function

```go
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
```
The main function sets up the HTTP server:

- It registers the getRoot handler for the root URL (/).
- It registers the getHello handler for the /hello URL.
- It starts the server on port 8080 using http.ListenAndServe.
- If the server is closed or an error occurs, appropriate messages are printed, and the program exits with a status code of 1.

### Error Handling
The server uses errors.Is to check if the error returned by http.ListenAndServe is http.ErrServerClosed. If a different error occurs, it prints the error and exits the program.


## Running the Server
To run the server, execute the following command in your terminal:
    
```bash
go run main.go
```

This will start the server on port 8080. You can access the root URL (`http://localhost:8080/`) and the hello URL (`http://localhost:8080/hello`) in your browser or using a tool like CURL.
```bash
curl http://localhost:8080/
curl http://localhost:8080/hello
```

The server will respond with the following messages for each route:
- http://localhost:8080/: Displays "This is my first web server."
- http://localhost:8080/hello: Displays "Hello, World!"

