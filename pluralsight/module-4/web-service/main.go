package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
  http.HandleFunc("/", Handler)
  http.ListenAndServe("127.0.0.1:8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}