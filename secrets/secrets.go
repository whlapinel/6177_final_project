package secrets

import (
	"os"
)

func GetSpeechKey() string {
	SPEECH_KEY := os.Getenv("SPEECH_KEY")
	if SPEECH_KEY == "" {
		panic("SPEECH_KEY is not set")
	}
	return SPEECH_KEY
}
