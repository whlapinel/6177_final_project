package main

import (
	"final_project/api_server"
	"final_project/web_server"
	"fmt"
)

func main() {
	fmt.Println("hello from main")
	web_server.Hello()
	api_server.Hello()
}
