package belajargolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(writer, req, "./resources/ok.html")
	} else {
		http.ServeFile(writer, req, "./resources/notfound.html")

	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOK)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
