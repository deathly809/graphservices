package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/deathly809/gods/graph"
)

func run(g graph.Graph, url string) {
	jsonStr, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Custom-Header", "json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
	fmt.Println("response header:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	buffer := bytes.NewBuffer(nil)
	err = json.Indent(buffer, body, "=", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println("response body:", buffer.String())
}

func bfs(g graph.Graph) {
	url := "http://127.0.0.1:8080/bfs"
	run(g, url)
}

func dfs(g graph.Graph) {
	url := "http://127.0.0.1:8080/dfs"
	run(g, url)
}

func main() {

	g := graph.New(graph.Properties{Directed: true})
	zero := g.AddVertex()  // 0
	one := g.AddVertex()   // 1
	two := g.AddVertex()   // 2
	three := g.AddVertex() // 3
	four := g.AddVertex()  // 4

	g.AddEdge(zero, one)
	g.AddEdge(zero, two)

	g.AddEdge(one, three)
	g.AddEdge(two, three)

	g.AddEdge(three, four)

	fmt.Println()
	fmt.Println("DFS")
	fmt.Println("--------------------------------------------------------------------")
	dfs(g)
	fmt.Println()

	fmt.Println()
	fmt.Println("BFS")
	fmt.Println("--------------------------------------------------------------------")
	bfs(g)
	fmt.Println()

}
