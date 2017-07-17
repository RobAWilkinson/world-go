package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:pw@/world")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {
			city := GetCity(db, id)
			b, err := json.Marshal(city)
			if err != nil {
				panic(err.Error())
			}
			w.Write([]byte(b))
		} else {
			w.Write([]byte("Gimme an id boy"))
		}
	})
	http.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		if page := r.URL.Query().Get("page"); page != "" {
			cities := GetCities(db, page)
			b, err := json.Marshal(cities)
			if err != nil {
				panic(err.Error())
			}
			w.Write([]byte(b))

		}
	})
	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {
			country := GetCountry(db, id)
			b, err := json.Marshal(country)
			if err != nil {
				panic(err.Error())
			}
			w.Write(b)
		}
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("listening on port 8080")
}
