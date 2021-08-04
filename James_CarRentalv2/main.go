package main

import (
	booking "first_tutorial/booking"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", booking.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/booking", booking.AddBooking).Methods("POST")
	myRouter.HandleFunc("/all/booking", booking.GetAllBooking).Methods("GET")
	myRouter.HandleFunc("/delete/{id}", booking.DeleteBooking).Methods("DELETE")
	myRouter.HandleFunc("/booking/{id}", booking.GetBooking).Methods("GET")
	myRouter.HandleFunc("/booking/{id}", booking.UpdateBooking).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	handleRequest()
}
