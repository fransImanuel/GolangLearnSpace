package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/fransimanuel/RestfulApiwithHexagonalArch/service"
)


type Customer struct{
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

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

	customers, _ := ch.service.GetAllCustomer()
	if rw.Header().Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	}else{
		rw.Header().Add("Content-Type", "application/xml")
		json.NewEncoder(rw).Encode(customers)
	}


}