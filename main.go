package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
	//"os"
)

type Response struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	data := Response{Title: "hello"}
	log.Println(data)
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(b)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
