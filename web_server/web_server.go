package web_server

import (
	"encoding/json"
	"final_project/models"
	"final_project/views"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Hello() {
	fmt.Println("hello from web_server")
}

type Person struct {
	Name string
}

func Run() {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./static/")
	fmt.Println("staticFileDirectory: ", staticFileDirectory)
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	fmt.Println("staticFileHandler: ", staticFileHandler)
	r.PathPrefix("/static/").Handler(staticFileHandler)

	r.HandleFunc("/get-voices/{page}", showVoices)
	http.ListenAndServe(":8080", r)
}

func showVoices(w http.ResponseWriter, r *http.Request) {
	voices := getVoices()
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		page = 0
	}
	component := views.VoicesPage(voices, page)
	component.Render(r.Context(), w)
}

func getVoices() *[]models.Voice {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/get-voices", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	fmt.Println(res)
	var voices []models.Voice
	if err := json.NewDecoder(res.Body).Decode(&voices); err != nil {
		return nil
	}
	return &voices
}
