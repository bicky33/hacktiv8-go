package controller

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	hello := "Hello Go developers"
	json.NewEncoder(w).Encode(hello)
}
