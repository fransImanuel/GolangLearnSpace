package app

import (
	"log"
	"net/http"

	"github.com/fransimanuel/RestfulApiwithHexagonalArch/domain"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/service"
	"github.com/gorilla/mux"
)

func Start(){
	router := mux.NewRouter()

	//Wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id}", getCustomers).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// func createCustomers(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Post request received")
// }

// func getCustomers(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

