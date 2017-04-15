package main

import (
	"log"
	"net/http"

	"services/graph"
)

type Handler func(http.ResponseWriter, *http.Request)

func CreateLogger(handle Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("connection from %v with request %v\n", r.RemoteAddr, r.RequestURI)
		handle(w, r)
	}
}

func main() {
	log.Println("starting server...")
	http.HandleFunc("/bfs", CreateLogger(graph.BFSHandler))
	http.HandleFunc("/dfs", CreateLogger(graph.DFSHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
