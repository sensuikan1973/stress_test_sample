package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"
)

// Greeting : record of greetings table
type Greeting struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/greetings", func(w http.ResponseWriter, req *http.Request) {
		db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_for_qiita")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		if req.Method == "GET" {
			rows, err := db.Query("SELECT * FROM greetings ORDER BY RAND() LIMIT 1")
			if err != nil {
				panic(err.Error())
			}

			for rows.Next() {
				greeting := Greeting{}
				err := rows.Scan(&greeting.ID, &greeting.Text)
				if err != nil {
					panic(err.Error())
				}
				fmt.Fprintf(w, greeting.Text)
			}
		}

		if req.Method == "POST" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err.Error())
			}
			var greeting Greeting
			error := json.Unmarshal(body, &greeting)
			if error != nil {
				log.Fatal(error)
			}
			_, err = db.Exec("insert into greetings (text) values (?)", greeting.Text)
			if err != nil {
				panic(err.Error())
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
