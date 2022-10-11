package controller

import (
	"net/http"
	"quiz-1/response"
)

func PalindromController(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("text")
	var res response.ResponsePalindrom
	// palindrom check
	for i := 0; i < len(param); i++ {
		if param[i] != param[len(param)-1-i] {
			res.Code = http.StatusBadRequest
			res.Status = "Bad Request"
			res.Error = "Not Palindrom"
			response.WriteResponse(http.StatusBadRequest, w, res)
			return
		}
	}
	res.Code = http.StatusOK
	res.Status = "Ok"
	res.Message = "Palindrom"
	response.WriteResponse(http.StatusBadRequest, w, res)
}
