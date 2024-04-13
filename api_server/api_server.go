package api_server

import (
	"encoding/json"
	"final_project/models"
	"final_project/secrets"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello() {
	fmt.Println("hello from api_server")
}

func Run() {
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/get-voices", getVoices)
	r.Run(":8081") // listen and serve on
}

func getVoices(c *gin.Context) {

	url := "https://eastus.tts.speech.microsoft.com/cognitiveservices/voices/list"
	method := "GET"

	// payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", secrets.GetSpeechKey())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var voices []models.Voice
	if err := json.NewDecoder(resp.Body).Decode(&voices); err != nil {
		c.JSON(500, gin.H{
			"error": "error decoding response",
		})
		return
	}
	c.JSON(200, voices)
}
