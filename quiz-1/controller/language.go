package controller

import (
	"net/http"
	"quiz-1/domain"
	"quiz-1/response"
	"strconv"

	"github.com/gorilla/mux"
)

var ListLanguage []domain.JsonObject

type LanguageController struct {
}

func (controller *LanguageController) Create(w http.ResponseWriter, r *http.Request) {
	var payload domain.JsonObject

	if err := response.WriteRequest(r, &payload); err != nil {
		response.WriteResponse(http.StatusInternalServerError, w, err)
		return
	}

	ListLanguage = append(ListLanguage, payload)
	res := response.ResponseLanguange{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Success create data",
		Data:    payload,
	}

	response.WriteResponse(http.StatusOK, w, res)
}

func (controller *LanguageController) GetAll(w http.ResponseWriter, r *http.Request) {
	if len(ListLanguage) == 0 {
		res := response.ResponseLanguages{
			Code:    http.StatusOK,
			Status:  "Ok",
			Message: "Success get data",
			Data:    make([]domain.JsonObject, 0),
		}
		response.WriteResponse(http.StatusOK, w, res)
		return
	}
	res := response.ResponseLanguages{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Success get data",
		Data:    ListLanguage,
	}
	response.WriteResponse(http.StatusOK, w, res)
}

func (controller *LanguageController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteResponse(http.StatusInternalServerError, w, err)
		return
	}
	var payload domain.JsonObject

	if err := response.WriteRequest(r, &payload); err != nil {
		response.WriteResponse(http.StatusInternalServerError, w, err)
		return
	}
	ListLanguage[id] = payload
	res := response.ResponseLanguange{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Success update data",
		Data:    ListLanguage[id],
	}
	response.WriteResponse(http.StatusOK, w, res)

}

func (controller *LanguageController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteResponse(http.StatusInternalServerError, w, err)
		return
	}
	ListLanguage = append(ListLanguage[:id], ListLanguage[id+1:]...)
	res := response.ResponseWeb{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Success delete data with id " + params["id"],
	}
	response.WriteResponse(http.StatusOK, w, res)

}
