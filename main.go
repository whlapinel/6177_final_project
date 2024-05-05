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
	godotenv.Load()
	fmt.Println("HOST: ", os.Getenv("HOST"))
	fmt.Println("hello from main")
	web_server.Hello()
	api_server.Hello()
	go cache.TestCache()
	go web_server.Run()
	api_server.Run()
}
