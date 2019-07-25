package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello Kubernetes!"))
	})

	server := http.Server{
		Handler: router,
		Addr: ":3000",
	}

	fmt.Println("Starting server at port 3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
