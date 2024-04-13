package models

type Voice struct {
	Name            string `json:"Name"`
	DisplayName     string `json:"DisplayName"`
	LocalName       string `json:"LocalName"`
	ShortName       string `json:"ShortName"`
	Gender          string `json:"Gender"`
	Locale          string `json:"Locale"`
	LocaleName      string `json:"LocaleName"`
	SampleRateHertz string `json:"SampleRateHertz"`
	VoiceType       string `json:"VoiceType"`
	Status          string `json:"Status"`
	WordsPerMinute  string `json:"WordsPerMinute"`
}
