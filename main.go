package main

import (
	"final_project/api_server"
	"final_project/web_server"
	"final_project/web_server/cache"
	"fmt"
)

func main() {
	fmt.Println("hello from main")
	web_server.Hello()
	api_server.Hello()
	go cache.TestCache()
	go web_server.Run()
	api_server.Run()
}
