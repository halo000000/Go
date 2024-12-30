package main

import (

	"net/http"
)


func isRedy(w http.ResponseWriter,r *http.Request){
	resJson(w,200,struct{}{})
}


func handlErr(w http.ResponseWriter,r *http.Request){


	resError(w,400,"there was a errore parsing the json")
}