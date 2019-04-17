package main

import (
	"fmt"
	"net/http"

	"github.com/ayubmalik/nhsfinder"
	goji "goji.io"
	"goji.io/pat"
)

var postcodeDB *nhsfinder.PostcodeDB
var pharmacies []nhsfinder.Pharmacy

// GreeterService is a service
type GreeterService interface {
	GetMessage() string
}

// FinderService is ace
type FinderService struct {
	greeter GreeterService
}

func (f *FinderService) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", f.greeter.GetMessage())
}

type stubHello struct{}

func (s stubHello) GetMessage() string {
	return "boom"
}

func main() {

	var finder = FinderService{
		greeter: stubHello{},
	}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello"), finder.serveHTTP)
	fmt.Println("Server started at http://localhost:8000/hello")
	http.ListenAndServe("localhost:8000", mux)
}
