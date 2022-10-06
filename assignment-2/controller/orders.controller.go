package controller

import (
	"assignment-2/domain"
	"assignment-2/repository"
	"assignment-2/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderController struct {
	OrderRepository repository.OrderRepository
}

func (controller *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	var payload response.Order
	var data domain.Order
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	data.CustomerName = payload.CustomerName
	data.OrderedAt = payload.OrderedAt
	for _, v := range payload.Items {
		data.Items = append(data.Items, domain.Item{ItemCode: v.ItemCode, Description: v.Description, Quantity: v.Quantity})
	}
	err = controller.OrderRepository.Create(data)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Success create data",
	}
	response.WriteResponseWeb(w, webResponse)

}

func (controller *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramsId := params["id"]
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	err = controller.OrderRepository.Delete(orderId)
	fmt.Println(err)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Success delete data",
	}
	response.WriteResponseWeb(w, webResponse)
}

func (controller *OrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res []response.Order
	result, err := controller.OrderRepository.GetAll()
	for _, v := range result {
		var itemsRes []response.Item
		for _, z := range v.Items {
			var itemRes response.Item
			itemRes.Description = z.Description
			itemRes.ItemCode = z.ItemCode
			itemRes.LineItemId = z.ItemId
			itemRes.Quantity = z.Quantity
			itemsRes = append(itemsRes, itemRes)
		}
		res = append(res, response.Order{OrderId: v.OrderId, CustomerName: v.CustomerName, OrderedAt: v.OrderedAt, Items: itemsRes})
	}

	if err != nil {
		response.WriteResponseWeb(w, err.Error())
	}
	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "success get data",
		Data:    res,
	}
	response.WriteResponseWeb(w, webResponse)
}

func (controller *OrderController) Update(w http.ResponseWriter, r *http.Request) {
	var payload response.Order
	var data domain.Order
	params := mux.Vars(r)
	orderId, err := strconv.Atoi(params["id"])
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	data.CustomerName = payload.CustomerName
	data.OrderedAt = payload.OrderedAt

	for _, v := range payload.Items {
		data.Items = append(data.Items, domain.Item{ItemCode: v.ItemCode, Description: v.Description, Quantity: v.Quantity, ItemId: v.LineItemId})
	}
	err = controller.OrderRepository.Update(orderId, data)
	if err != nil {
		response.WriteResponseWeb(w, err)
		return
	}
	webResponse := response.WebResponse{
		Code:    http.StatusNoContent,
		Status:  "UPDATED",
		Message: "success update data",
	}
	response.WriteResponseWeb(w, webResponse)
}
