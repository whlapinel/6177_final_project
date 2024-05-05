package api_server

import (
	"bytes"
	"encoding/json"
	"final_project/models"
	"final_project/secrets"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello() {
	fmt.Println("hello from api_server")
}

func Run() {
	r := gin.Default()
	r.Use(authorizationMiddleware())
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/get-voices", getVoices)
	r.GET("/api/tts", tts)
	r.Run("localhost:8081") // listen and serve on
}

func authorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("running authorization middleware")
		if c.GetHeader("Authorization") != secrets.GetApiToken() {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GetSpeechToken() (*string, error) {
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, issueTokenUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", secrets.GetSpeechKey())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-length", "0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("resp: ", resp)
	var token *string
	if resp.StatusCode != 200 {
		fmt.Println("status code not 200")
		return nil, fmt.Errorf("status code not 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	token = new(string)
	*token = string(body)

	// if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	return token, nil
}

func tts(c *gin.Context) {
	text := c.Query("text")
	voice := c.Query("voice")
	fmt.Println("text:", text)
	fmt.Println("voice:", voice)
	payload := fmt.Sprintf(
		`
	<speak version="1.0" xml:lang="en-US">
	<voice xml:lang='en-US' name='%s'>
	%s
	</voice>
	</speak>
	`,
		voice, text)

	fmt.Println("payload:", payload)
	token, err := GetSpeechToken()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "error getting token",
		})
		return
	}
	data := []byte(payload)

	client := &http.Client{}
	req, err := http.NewRequest("POST", ttsUrl, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", secrets.GetSpeechKey())
	req.Header.Set("Authorization", "Bearer "+*token)
	req.Header.Set("Content-Type", "application/ssml+xml")
	req.Header.Set("X-Microsoft-OutputFormat", "audio-16khz-32kbitrate-mono-mp3")
	req.Header.Set("User-Agent", "Will Lapinel's School Project")

	fmt.Println("req:", req)
	fmt.Println("req.Header:", req.Header)
	fmt.Println("req.Body:", req.Body)
	fmt.Println("So far so good.... !!")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("sending request to tts api failed")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"error": "error from tts api",
			"req":   req,
		})
		return
	}
	file, err := os.Create("output.mp3")
	if err != nil {
		fmt.Println("error creating file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error copying file")
		fmt.Println(err)
		return
	}
	c.File("output.mp3")
}

func getVoices(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", voicesUrl, nil)
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
