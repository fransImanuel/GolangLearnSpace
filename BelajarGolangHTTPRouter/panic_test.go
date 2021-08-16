package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T)  {
	router := httprouter.New()

	router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, error interface{} ){
		fmt.Fprint(rw, "Panic : ", error)
	}

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprint(rw, "Hello World")
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body , _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}