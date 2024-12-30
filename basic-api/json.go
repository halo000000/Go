package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func resError(w http.ResponseWriter,code int,msg string){

	if code>499{
		log.Printf("responding with %v Error\n",code)
		
	}

	type errorRespons struct{
		Error string `json:"error"`
	}

	resJson(w,code,errorRespons{Error: msg})


}

func resJson(w http.ResponseWriter,code int ,paylod interface{}){

data,err:=json.Marshal(paylod)

if err !=nil{
	fmt.Printf("there was a error with parsing json %v",paylod)
	w.WriteHeader(500)
	return
}
w.Header().Add("Content-Type","aplication/json")
w.WriteHeader(code)
w.Write(data)



}