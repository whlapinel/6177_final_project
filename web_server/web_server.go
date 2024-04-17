package web_server

import (
	"encoding/json"
	"final_project/models"
	"final_project/views"
	"final_project/views/components"
	"final_project/web_server/cache"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	r.HandleFunc("/test-call", testCall)
	r.HandleFunc("/get-audio", renderAudioElement)

	r.HandleFunc("/get-voices/{page}", renderVoices)
	r.HandleFunc("/home", renderHome)
	r.HandleFunc("/", renderHome)
	r.HandleFunc("/about", renderAbout)
	http.ListenAndServe(":8080", r)
}

func renderAudioElement(w http.ResponseWriter, r *http.Request) {
	voice := r.FormValue("voice")
	fmt.Println("Voice: ", voice)
	text := r.FormValue("text")
	fmt.Println("Text: ", text)
	components.AudioElement(text, voice).Render(r.Context(), w)

}

func testCall(w http.ResponseWriter, r *http.Request) {
	voice := r.FormValue("voice")
	fmt.Println("Voice: ", voice)
	text := r.FormValue("text")
	fmt.Println("Text: ", text)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/tts?text="+text+"&voice="+voice, nil)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("setting authorization header")
	req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: ", res.Status)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()
	fmt.Println("setting content type header")
	w.Header().Set("Content-Type", "audio/mp3")
	io.Copy(w, res.Body)
}

func renderAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rendering about page")
	views.AboutPage().Render(r.Context(), w)
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rendering home page")
	voices := getVoices()
	views.HomePage(voices).Render(r.Context(), w)
}

func renderVoices(w http.ResponseWriter, r *http.Request) {
	voices := getVoices()
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		page = 0
	}
	if strings.Contains(r.Header.Get("HX-Current-URL"), "get-voices") {
		fmt.Println("HX header indicates user is already on the voices page, returning partial content")
		component := views.VoicesPage(voices, page, true)
		component.Render(r.Context(), w)
	} else {
		fmt.Println("Returning full page content")
		component := views.VoicesPage(voices, page, false)
		component.Render(r.Context(), w)
	}
}

func fetchVoices() (*[]models.Voice, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/get-voices", nil)
	req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
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
	dataCache.Set("voices", voices, 5*time.Minute)
	return voices
}
