package main

import (
	"Server/pkg/api"
	"Server/pkg/repository"
	"Server/config"
	"log" 
	"Server/pkg/cache"

	"github.com/gorilla/mux"
)


//const connStrDB = "postgres://nicis:123@postgres:5432/kudago"
//const connStrDB = "postgres://nicis:123@localhost:5432/kudago"

func main() {
	log.Println("START SERVER")
	cfg, err := config.LoadConfig()//подключение файла .env с конфигами
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		return
	}
	//подключение PgSQL
	db,err := repository.New(cfg.DBConnString)
	if(err!=nil){
		log.Fatal(err.Error())
	}

	
    redisClient := cache.NewClient(cfg.ReddisAddr, cfg.RedisPassword, 0)

	//обработчики
	api:= api.New(mux.NewRouter(),db,cfg,redisClient)
	api.Handle(cfg)
	//log.Fatal(api.ListenAndServe("localhost:8090"))
	log.Fatal(api.ListenAndServe("0.0.0.0:8090"))
}
