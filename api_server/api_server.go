package api_server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello() {
	fmt.Println("hello from api_server")
}

func Run() {
	SPEECH_KEY := os.Getenv("SPEECH_KEY")
	if SPEECH_KEY == "" {
		panic("SPEECH_KEY is not set")
	}
	fmt.Println("SPEECH_KEY is set, value is", SPEECH_KEY)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get-voices", getVoices)
	r.Run(":8081") // listen and serve on
}

func getVoices(c *gin.Context) {
	// make call to endpoint instead of returning hardcoded values

	c.JSON(200, gin.H{
		"voices": []string{"en-US-AriaNeural", "en-US-GuyNeural", "en-US-JennyNeural", "en-US-ZiraNeural"},
	})
}
