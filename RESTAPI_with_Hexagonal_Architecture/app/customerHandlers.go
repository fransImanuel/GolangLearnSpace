package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/fransimanuel/RestfulApiwithHexagonalArch/logger"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/service"
	"github.com/gorilla/mux"
)

// func greet(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(rw, "Hello Worlds")
// }

type CustomerHandlers struct{
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Aisish", "New Delhi", "110075"},
	// 	{"Robh", "New Delhi", "110076"},
	// }

	customers,err := ch.service.GetAllCustomer()
	if err!=nil {
		logger.Debug(strconv.Itoa(err.Code))
	}
	if rw.Header().Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
		writeResponse(rw, http.StatusOK, customers)
	}else{
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomers(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err!=nil {
		writeResponse(rw, err.Code, err.AsMessage())
	}else{
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}){
	
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil{
		panic(err)
	}
}