package main

import (
	"context"
	"fmt"
	"log"
	"to_do_app/api"
	"to_do_app/config"
	"to_do_app/service"
	"to_do_app/storage/postgres"
)

func main(){
	cfg := config.Load()

	postgreStore, err := postgres.New(context.Background(),cfg)
	if err != nil{
		log.Fatalln("Error while connecting to postgres database!", err.Error())
		return 
	}

	defer postgreStore.Close()

	services := service.New(postgreStore)
	server := api.New(services)

	err = server.Run("localhost:8080")
	if err != nil{
		log.Fatalln("Error, router is not listening and serving in http server!", err.Error())
		return
	}

	fmt.Println("SERVER LISTENING ON LOCALHOST:8080")
}