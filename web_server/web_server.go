package web_server

import (
	"encoding/json"
	"final_project/models"
	"final_project/views"
	"final_project/web_server/cache"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Hello() {
	fmt.Println("hello from web_server")
}

type Person struct {
	Name string
}

var dataCache = cache.NewCache()

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

func fetchVoices() (*[]models.Voice, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/get-voices", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	fmt.Println(res)
	var voices []models.Voice
	if err := json.NewDecoder(res.Body).Decode(&voices); err != nil {
		return nil, err
	}
	return &voices, nil
}

func getVoices() *[]models.Voice {
	// Check if the data is in the cache
	if data, found := dataCache.Get("voices"); found {
		fmt.Println("Data found in cache")
		return data.(*[]models.Voice)
	}
	// If not in cache, fetch the data
	fmt.Println("Data not found in cache")
	voices, err := fetchVoices()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dataCache.Set("voices", voices, 30*time.Second)
	return voices
}
