package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

type Response struct {
	Status Status `json:"status"`
}

func GenerateRandomNumber() int {
	min := 1
	max := 100
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix())
	go func() {
		tick := time.Tick(time.Second * 15)
		for range tick {
			StatusWrite()
		}
	}()

	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}

func StatusWrite() {
	data := Status{Wind: GenerateRandomNumber(), Water: GenerateRandomNumber()}
	response := Response{Status: data}
	fileName := "./file/status.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	encode := json.NewEncoder(file)
	encode.Encode(response)
	defer file.Close()
}
