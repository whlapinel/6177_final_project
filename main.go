package main

import (
	"final_project/api_server"
	"final_project/web_server"
	"final_project/web_server/cache"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	fmt.Println("HOST: ", os.Getenv("HOST"))
	fmt.Println("SPEECH_KEY: ", os.Getenv("SPEECH_KEY"))
	fmt.Println("API_TOKEN: ", os.Getenv("API_TOKEN"))
	fmt.Println("SPEECH_REGION: ", os.Getenv("SPEECH_REGION"))
	fmt.Println("hello from main")
	web_server.Hello()
	api_server.Hello()
	go cache.TestCache()
	go web_server.Run()
	api_server.Run()
}
