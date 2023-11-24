package main

import "net/http"

func main() {
	http.HandleFunc("/", mapping)
	http.ListenAndServe(":8080", nil)
}
