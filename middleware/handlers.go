package middleware

import (
	"cdi/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
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

var id model.Idea

func MakeIdea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	titulo := params["t"]
	desc := params["d"]

	id.Title = titulo
	id.Description = desc
}

func CreateIdea(w http.ResponseWriter, r *http.Request) {
	insertIdea()
}

func DeleteIdea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	removeIdea(id)
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

func UpdateIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	current_title := vars["title"]
	new_title := vars["new_title"]
	new_desc := vars["new_description"]

	updateData(current_title, new_title, new_desc)
}

//----------------------------------------------------------------------------//

func insertIdea() {
	db := openDatabase()

	defer db.Close()

	query := "INSERT INTO ideia (title, description) VALUES ($1, $2);"

	err := db.QueryRow(query, id.Title, id.Description)

	if err != nil {
		fmt.Println("\n")
	}

	fmt.Println("[DATABASE] - DATA INSERTED -")
}

func removeIdea(id string) {
	db := openDatabase()

	query := "DELETE FROM ideia WHERE title=($1);"

	fmt.Println(id)

	err := db.QueryRow(query, id)

	if err != nil {
		fmt.Println("[DATABASE] - DATA REMOVED -")
	}
}

func updateData(old string, title string, desc string) {
	db := openDatabase()

	defer db.Close()

	query := "UPDATE ideia SET title=($1), description=($2) WHERE title=($3)"

	err := db.QueryRow(query, title, desc, old)

	if err != nil {
		fmt.Println("DATA UPDATED")
	}
}
