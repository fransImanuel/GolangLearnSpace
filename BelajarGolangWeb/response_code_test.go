package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, req *http.Request){
	name := req.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)//Bad Request
		fmt.Fprint(writer, "name is empty")
	}else{
		writer.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T){
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T){
	request := httptest.NewRequest("GET","http://localhost:8080?name=eko",nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}