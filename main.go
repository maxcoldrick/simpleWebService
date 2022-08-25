package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/default", handleDefault)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}

func handleDefault(w http.ResponseWriter, r *http.Request) {

	responseJson, _ := json.Marshal(DefaultResponse{Message: "hello world"})

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(responseJson)
	if err != nil {
		return
	}
}

type DefaultResponse struct {
	Message string `json:"message"`
}
