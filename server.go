package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/info", Info)
	router.HandleFunc("/infohost", Info)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Info(w http.ResponseWriter, r *http.Request) {
        hostname, _ := os.Hostname()
	fmt.Fprintln(w, hostname)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
