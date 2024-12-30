package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hi")
	godotenv.Load(".env")
	port:=os.Getenv("PORT")

	

	if port ==""{
		log.Fatal("there is no port pleas addd port on the .env file")
	}

	rout := chi.NewRouter()
	rout.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))


	  rout.Get("/redy",isRedy)
	  rout.HandleFunc("/error",handlErr)

	server:=&http.Server{
		Handler: rout,
		Addr: ":"+port,
	}
	fmt.Printf("http://localhost:%v\n", port)
	err:=server.ListenAndServe()
	if err !=nil{
		log.Fatal("there was a error ",err)
	}
	
}