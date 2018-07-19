package appcontext

import (
	"database/sql"
	"log"

	"github.com/go-squads/reuni-server/config"
	_ "github.com/lib/pq"
)

type appContext struct {
	db *sql.DB
}

var context *appContext

func initDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", config.GetConnectionString())

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db, nil
}

func InitContext() {
	db, _ := initDB()
	log.Print("Connection Established")
	context = &appContext{
		db: db,
	}
}

func GetDB() *sql.DB {
	return context.db
}
