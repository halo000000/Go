package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"strconv"

	_ "github.com/lib/pq"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	db, err := sql.Open("postgres", "user=postgres password=123456789 dbname=go-basic-api sslmode=disable")
	if err != nil {
		fmt.Print("error connecting to the databse ", err)
	}

	homePage := func(w http.ResponseWriter, r *http.Request) {
		html := template.Must(template.ParseFiles("index.html"))
		html.Execute(w, html)
	}

	shorten := func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		url := r.PostFormValue("url")

		if url == "" {
			http.Error(w, "URL is empty", http.StatusBadRequest)
			return
		}
		db.Exec(`INSERT INTO urls (url) VALUES ($1)`, url)

		var id int
		err = db.QueryRow(`SELECT id FROM urls WHERE url = $1`, url).Scan(&id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving ID: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("http://localhost:8000/url/%d", id)))

	}

	redirect := func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid URL ID", http.StatusBadRequest)
			return
		}

		var url string
		fmt.Print("this is url", url)
		err = db.QueryRow(`SELECT url FROM urls WHERE id = $1`, id).Scan(&url)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving URL: %v", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/shorten/", shorten)
	http.HandleFunc("/url/{id}", redirect)

	fmt.Print("http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
