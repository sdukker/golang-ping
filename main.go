package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Starting ping sequence!\n")

	requestURL := os.Getenv("HEALTH_URL")
	fmt.Println("Making request to " + requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	if res.StatusCode != 200 {
		fmt.Printf("Unexpected response: %s\n", err)
		os.Exit(1)
	}

	w.Write([]byte("pong"))
	fmt.Println("pong")
	fmt.Println(time.Now(), r.Method, r.RequestURI, "200")
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	fmt.Println("OK")
	fmt.Println(time.Now(), r.Method, r.RequestURI, "200")
}

func main() {
	http.HandleFunc("/ping", handlerPing)
	http.HandleFunc("/health", handlerHealth)

	port := ":" + os.Getenv("PORT")
	fmt.Printf("ping listening on 0.0.0.0, port %s\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting ping server: ", err)
	}
}
