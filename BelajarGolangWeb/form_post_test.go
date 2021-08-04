package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err!= nil{
		panic(err)
	}

	firstname := req.PostForm.Get("firstname")
	lastname := req.PostForm.Get("lastname")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("firstname=Eko&lastname=khannedy")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}