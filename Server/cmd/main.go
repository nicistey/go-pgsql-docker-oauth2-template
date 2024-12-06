package main

import (
	"Server/pkg/api"
	"Server/pkg/repository"
	"log" 

	"github.com/gorilla/mux"
)


const connStrDB = "postgres://nicis:123@postgres:5432/kudago"
//const connStrDB = "postgres://nicis:123@localhost:5432/kudago"

func main() {
	log.Println("START SERVER")
	db,err := repository.New(connStrDB)
	if(err!=nil){
		log.Fatal(err.Error())
	}
	api:= api.New(mux.NewRouter(),db )
	api.Hadle()
	// log.Fatal(api.ListenAndServe("localhost:8090"))
	log.Fatal(api.ListenAndServe("0.0.0.0:8090"))
}
