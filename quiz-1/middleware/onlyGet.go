package middleware

import (
	"net/http"
	"quiz-1/response"
)

func OnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			res := response.ResponseLanguages{
				Code:   http.StatusMethodNotAllowed,
				Status: "Method Not Allowed",
				Error:  r.Method + " Not Allowd on this path " + r.Host + r.RequestURI,
			}
			response.WriteResponse(http.StatusMethodNotAllowed, w, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}
