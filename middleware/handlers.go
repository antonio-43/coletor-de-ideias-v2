package middleware

import (
	"cdi/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "35.246.126.180"
	port     = 5432
	user     = "gqzexqkm"
	password = "lX4v-GgpG4thqIIC34YVQo_xhj-I9baZ"
	dbname   = "gqzexqkm"
)

func openDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE] - CONNECTED -")

	return db
}

/*
func makeIdeia(a string, b string) *model.Idea {
    id := model.Idea{Title: a, Description: b}
    return &id
}
*/
func CreateIdea(w http.ResponseWriter, r *http.Request) {
	id := model.Idea{Title: "TESTE", Description: "WORK :)"}
	insertIdea(id)
}

func ShowData(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()
	defer db.Close()

	query := "SELECT * FROM ideia;"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var (
			head string
			body string
		)

		var resp model.Resp

		if err := rows.Scan(&head, &body); err != nil {
			panic(err)
		}

		resp.Head = head
		resp.Body = body

		b, err := json.Marshal(resp)

		if err != nil {
			panic(err)
			return
		}

		fmt.Fprintf(w, string(b))
	}

}

//----------------------------------------------------------------------------//

func insertIdea(idea model.Idea) {
	db := openDatabase()

	defer db.Close()

	query := "INSERT INTO ideia (title, description) VALUES ($1, $2);"

	err := db.QueryRow(query, idea.Title, idea.Description)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[DATABASE] - DATA INSERTED -")
}
