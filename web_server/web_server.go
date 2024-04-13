package web_server

import (
	"encoding/json"
	"final_project/models"
	"fmt"
	"html/template"
	"net/http"

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
	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		people := map[string][]Person{
			"People": {
				{Name: "Alice"},
				{Name: "Bob"},
				{Name: "Charlie"},
			},
		}
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, people)
	})

	r.HandleFunc("/add-person", h2)
	r.HandleFunc("/get-voices", showVoices)
	http.ListenAndServe(":8080", r)
}

func showVoices(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/get-voices", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res)
	var voices []models.Voice
	if err := json.NewDecoder(res.Body).Decode(&voices); err != nil {
		return
	}
	var htmlStr string
	for _, voice := range voices {
		htmlStr += fmt.Sprintf("<li>%s</li>", voice.ShortName)
		fmt.Println(voice.Name)
	}
	tmpl, _ := template.New("foo").Parse(htmlStr)
	tmpl.Execute(w, nil)
}

func h2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from h2")
	fmt.Println(r.Header.Get("HX-Request"))
	name := r.FormValue("name")
	fmt.Println("name: ", name)
	htmlStr := fmt.Sprintf("<li>%s</li>", name)
	tmpl, _ := template.New("foo").Parse(htmlStr)
	tmpl.Execute(w, nil)
}
