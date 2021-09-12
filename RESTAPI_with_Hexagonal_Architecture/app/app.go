package app

import (
	"log"
	"net/http"
	"time"

	"github.com/fransimanuel/RestfulApiwithHexagonalArch/domain"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start(){
	router := mux.NewRouter()

	//Wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

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



func getDbClient() *sqlx.DB{
	client, err := sqlx.Open("mysql", "root@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}