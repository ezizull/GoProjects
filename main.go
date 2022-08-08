package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	route "goproject.com/pkg/routes"
)

func main() {
	/* HRM */
	route.HRMroutes()
	
	/* CRM */
	route.CRMroutes()

	/* web service and api */
	r := mux.NewRouter()

	http.Handle("/", r)
	
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
