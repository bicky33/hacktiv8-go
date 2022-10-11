package response

import "quiz-1/domain"

type ResponseLanguages struct {
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
	Message string              `json:"message,omitempty"`
	Error   string              `json:"error,omitempty"`
	Data    []domain.JsonObject `json:"data"`
}

type ResponseLanguange struct {
	Code    int               `json:"code"`
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
	Data    domain.JsonObject `json:"data,omitempty"`
}

type ResponseWeb struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type ResponsePalindrom struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error"`
}
