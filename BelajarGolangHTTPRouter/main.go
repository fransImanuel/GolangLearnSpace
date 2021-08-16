package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main()  {
	router := httprouter.New()

	// router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, error interface{} ){
	// 	fmt.Fprint(rw, "Panic : ", error)
	// }

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello HttpRouter")
		// panic("Ups")
	})

	server := http.Server{
		Handler: router,
		Addr: "localhost:3000",
	}

	server.ListenAndServe()
}