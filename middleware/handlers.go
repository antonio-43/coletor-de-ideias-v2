package middleware

import (
	"cdi/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	//	"log"
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

/*
   1. criar o form - ok
   2. receber na api os dados
   3. criar um model com os dados recebidos
   4. inserir na base de dados
*/

var id model.Idea

func MakeIdea(w http.ResponseWriter, r *http.Request) {
	html := "<html><form><p1>Titulo</p1><input type='text' name='titulo'></input><p1>Descrição</p1><input type='text' name='descricao'><input type='submit' value='REGISTAR'></form></html>"

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, html)
	}

	id.Title = r.FormValue("titulo")
	id.Description = r.FormValue("descricao")
}

func CreateIdea(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()

	defer db.Close()

	query := "INSERT INTO ideia (title, description) VALUES ($1, $2);"

	err := db.QueryRow(query, id.Title, id.Description)

	if err != nil {
		fmt.Println("data ok")
	}

	fmt.Println("[DATABASE] - DATA INSERTED -")
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
