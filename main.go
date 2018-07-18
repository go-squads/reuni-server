package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"database/sql"
	"github.com/go-squads/reuni-server/server"
    _ "github.com/lib/pq"
)


func multiply(x,y int) int {
	return x*y
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(os.Getenv("POSTGRES_DB")))
}



func main() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("DB_USER"),
	os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection established")
	router := server.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}