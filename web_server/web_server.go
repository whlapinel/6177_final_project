package web_server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello() {
	fmt.Println("hello from web_server")
}

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", r)
}
