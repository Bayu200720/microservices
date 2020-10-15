package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bayu200720/microservices/menu-service/handler"
	"github.com/gorilla/mux"
	"github.com/wskurniawan/intro-microservice/auth/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
