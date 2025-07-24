package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handlerPong(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Starting ping sequence!\n")

	requestURL := os.Getenv("HEALTH_URL")
	fmt.Println("Making request to " + requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)

	w.Write([]byte("ping"))
	fmt.Println("ping")
	fmt.Println(time.Now(), r.Method, r.RequestURI, "200")
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	fmt.Println("OK")
	fmt.Println(time.Now(), r.Method, r.RequestURI, "200")
}

func main() {
	http.HandleFunc("/pong", handlerPong)
	http.HandleFunc("/health2", handlerHealth)

	port := ":" + os.Getenv("PORT")
	fmt.Printf("pong listening on 0.0.0.0, port %s\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting pong server: ", err)
	}
}
