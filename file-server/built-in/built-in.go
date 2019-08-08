package main

import "net/http"

const PATH = "public"

func main() {
	http.ListenAndServer(":8080", http.FileServer(http.Dir(PATH)))
}
