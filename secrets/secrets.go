package secrets

import (
	"os"

	"github.com/joho/godotenv"
)

func GetSpeechKey() string {
	godotenv.Load()
	SPEECH_KEY := os.Getenv("SPEECH_KEY")
	if SPEECH_KEY == "" {
		panic("SPEECH_KEY is not set")
	}
	return SPEECH_KEY
}

func GetApiToken() string {
	API_TOKEN := os.Getenv("API_TOKEN")
	if API_TOKEN == "" {
		panic("API_TOKEN is not set")
	}
	return API_TOKEN
}
