package response

import (
	"encoding/json"
	"net/http"
)

func WriteResponseWeb(w http.ResponseWriter, response interface{}) {
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	if err != nil {
		encode.Encode(err.Error())
	}
}
