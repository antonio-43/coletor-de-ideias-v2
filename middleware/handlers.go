package middleware

import (
	model "cdi/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver
)

func openDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loadin .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+"password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println()
		fmt.Println(err)
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println()
		fmt.Println(err)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func UpdateIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	currentTitle := vars["title"]
	newTitle := vars["new_title"]
	newDesc := vars["new_description"]

	updateData(currentTitle, newTitle, newDesc)
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
