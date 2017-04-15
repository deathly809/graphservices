package graph

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/deathly809/gods/graph"
)

// BFSHandler will handle the breadth first search of a graph
func BFSHandler(w http.ResponseWriter, r *http.Request) {
	target := graph.New(graph.Properties{})

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleReadError(w, r, err)
		return
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		handleJSONUnMarshalError(w, r, err, data)
		return
	}

	results := graph.BFS(target)
	data, err = json.Marshal(results)
	if err != nil {
		w.Write([]byte("error" + err.Error()))
		panic(err)
	} else {
		w.Write(data)
	}
}
