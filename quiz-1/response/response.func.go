package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponse(code int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encode.Encode(err)
	}
}

func WriteRequest(r *http.Request, data interface{}) error {
	decode := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	err := decode.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
