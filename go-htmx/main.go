package main

import (
	"fmt"
	"html/template"
	"time"
	"database/sql"
	// "io"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

type Note struct{
	Id string
	Title string
	Text string
}

func main (){

	// notes := []note{

    //         {Title: "making diner", Text: "make diner for 4 pepole "},
    //         {Title: "buy milk", Text: "go to super market and buy sam milk "},
    //         {Title: "check the bed ", Text: "make your bed and do clening around my room"},
	// 	}



		db, err := sql.Open("postgres", "user=postgres password=123456789 dbname=go-basic-api sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

h1:=func (w http.ResponseWriter, r *http.Request)  {
	
row,err := db.Query("SELECT id, title, text FROM notes")
if err != nil {
	log.Printf("Error querying database: %v", err)
	http.Error(w, "Error querying database", http.StatusInternalServerError)
	return
}

defer row.Close()

var notes []Note
for row.Next() {
	var note Note
	row.Scan(&note.Id,&note.Title, &note.Text)
	notes = append(notes, note)
}

html, _ := template.ParseFiles("index.html")

err = html.Execute(w,notes)
if err != nil {
	fmt.Println(err)
	http.Error(w, "Error executing template", http.StatusInternalServerError)
	return
}
	
}



h2:=func (w http.ResponseWriter, r *http.Request)  {
	html:=template.Must(template.ParseFiles("price.html"))
	html.Execute(w,html)
}
mainPage:=func (w http.ResponseWriter, r *http.Request)  {
	html:=template.Must(template.ParseFiles("home.html"))
	html.Execute(w,html)
}


addNote:=func (w http.ResponseWriter, r *http.Request)  {
	time.Sleep(1 * time.Second)
// we use this to know if the requst is comming frome the HTMX
log.Print(r.Header.Get("HX-Request"))
title:=r.PostFormValue("title")
text:=r.PostFormValue("text")

_, err := db.Exec(`INSERT INTO notes (id,title, text) VALUES ($1, $2,$3)`, text,title,text )
if err != nil {
	log.Fatalf("Error inserting data: %v", err)
}

html:=fmt.Sprintf(`<li class='list-group-item d-flex justify-content-between align-items-start'>
          <div class='ms-2 me-auto'>
            <div class='fw-bold'>%v</div>
            %v
          </div>
          <span class='badge text-bg-primary rounded-pill' style="font-size: large;">✯</span>
        </li>`,title,text)

		temp,_:=template.New("a").Parse(html)
		temp.Execute(w,nil)

}

http.HandleFunc("/add-note/",addNote)
http.HandleFunc("/note",h1)
http.HandleFunc("/price",h2)
http.HandleFunc("/",mainPage)



fmt.Print("http://localhost:8000/")
log.Fatal(http.ListenAndServe(":8000",nil))


}