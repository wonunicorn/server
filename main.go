package main

import(
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux@v1.7.4"
)

func main(){
	multiplexer := mux.NewRouter()

	srv := &http.Server{
		Add: "126.0.0.1:5050",
		Handler: multiplexer
	}

	multiplexer.HandleFunc("/user/", handler.HandleUserGet).Method("GET")
	multiplexer.HandleFunc("/user/",handler.HandleUserPost).Method("POST")
	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}