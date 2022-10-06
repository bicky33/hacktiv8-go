package main

import "api-go/routers"

func main() {
	var PORT = ":8080"
	server := routers.StartServer()
	server.Run(PORT)
}
