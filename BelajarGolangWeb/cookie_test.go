package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(wr http.ResponseWriter, req *http.Request )  {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(wr, cookie)
	fmt.Fprintf(wr, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("X-PZN-Name")
	if err!=nil {
		fmt.Fprint(writer, "No Cookie")
	}else{
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err!=nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=eko", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Eko"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder,request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

