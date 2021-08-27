package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer1 struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONArray(t *testing.T) {
	customer := Customer1{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        30,
		Married:    true,
		Hobbies: []string{
			"Gaming", "Coding", "Reading",
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true,"Hobbies":["Gaming","Coding","Reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer1{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer1{
		FirstName: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalan belum ada",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan lagi dibangun",
				Country:    "Indonesia",
				PostalCode: "888888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan belum ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan lagi dibangun","Country":"Indonesia","PostalCode":"888888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer1{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan belum ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan lagi dibangun",
			Country:    "Indonesia",
			PostalCode: "888888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
