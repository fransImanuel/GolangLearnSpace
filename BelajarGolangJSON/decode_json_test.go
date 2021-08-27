package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// type Customer struct {
// 	FirstName  string
// 	MiddleName string
// 	LastName   string
// 	Age        int
// 	Married    bool
// }

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":  "Eko","MiddleName": "Kurniawan","LastName":"Khannedy","Age":30,"Married": true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
