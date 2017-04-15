package graph

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type message struct {
	Name string
	Body string
	Time int64
}

func handleReadError(w http.ResponseWriter, r *http.Request, err error) {
	log.Print(r.RemoteAddr + ": " + err.Error())

	m := message{
		"Error",
		"There was an error reading:" + err.Error(),
		int64(time.Now().Nanosecond()),
	}
	res, _ := json.Marshal(m)
	w.Write(res)
}

func handleNILData(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RemoteAddr + " : data is nil")
}

func handleJSONUnMarshalError(w http.ResponseWriter, r *http.Request, err error, data []byte) {
	log.Print(r.RemoteAddr + ": " + err.Error())

	m := message{
		"Error",
		"There was an error (" + err.Error() + ") with your input:" + string(data),
		int64(time.Now().Nanosecond()),
	}
	res, _ := json.Marshal(m)
	w.Write(res)
}
